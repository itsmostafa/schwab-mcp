# Simulated Phishing Platform (Security Awareness Training)

## Key Concepts

### System Architecture Overview

A simulated phishing platform lets organizations voluntarily test employee susceptibility to phishing. Employees receive realistic (but harmless) phishing emails; the platform tracks opens and clicks to measure awareness. This is the model behind tools like KnowBe4 and Gophish.

```
┌─────────────────────────────────────────────────────────┐
│                     API Server (Gin)                     │
│  Org CRUD │ Employee Mgmt │ Campaign CRUD │ Analytics   │
└────────┬──────────┬──────────────┬──────────────────────┘
         │          │              │
         ▼          ▼              ▼
┌────────────┐ ┌──────────┐ ┌───────────────┐
│  MongoDB   │ │  Redis   │ │ Tracking Svc  │
│ (primary   │ │ (cache,  │ │ (pixel + link │
│  store)    │ │  jobs,   │ │  endpoints)   │
└────────────┘ │  rate    │ └───────┬───────┘
               │  limit,  │         │
               │  stats)  │         ▼
               └────┬─────┘   ┌───────────┐
                    │         │  MongoDB   │
                    ▼         │ (events)   │
              ┌──────────┐   └───────────┘
              │  Asynq   │
              │  Worker   │
              │  Pool     │
              └─────┬────┘
                    ▼
              ┌──────────┐
              │   SMTP   │
              │ (send    │
              │  emails) │
              └──────────┘
```

| Component | Responsibility | Key Libraries |
|-----------|---------------|---------------|
| API Server | REST API for orgs, employees, campaigns, analytics | `gin-gonic/gin` |
| Campaign Scheduler | Randomized send-time calculation, task enqueueing | `crypto/rand`, `hibiken/asynq` |
| Email Worker Pool | Dequeue tasks, render templates, send via SMTP | `hibiken/asynq`, `net/smtp`, `html/template` |
| Tracking Service | Record email opens (pixel) and link clicks (redirect) | `gin-gonic/gin`, `go.mongodb.org/mongo-driver` |
| MongoDB | Primary data store for orgs, employees, campaigns, events | `go.mongodb.org/mongo-driver` |
| Redis | Job queue (asynq), rate limiting, HyperLogLog stats, caching, distributed locks | `redis/go-redis/v9`, `redis_rate`, `redislock` |
| Circuit Breaker | Wrap SMTP calls to prevent cascading failures | `sony/gobreaker` |

---

### Data Models — MongoDB Collections

The platform uses four primary collections. MongoDB is chosen over a relational database because email event tracking produces high write volumes with a naturally document-shaped schema (variable metadata per event type), and the schema evolves as new event types are added.

**`organizations`** — Enrolled companies:

```go
type Organization struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
    Domain    string             `bson:"domain"`
    PlanTier  string             `bson:"plan_tier"` // "free", "pro", "enterprise"
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
}
```

**`employees`** — Per-organization employee roster:

```go
type Employee struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    OrgID        primitive.ObjectID `bson:"org_id"`
    Email        string             `bson:"email"`
    FirstName    string             `bson:"first_name"`
    LastName     string             `bson:"last_name"`
    Department   string             `bson:"department"`
    ImportedAt   time.Time          `bson:"imported_at"`
}
```

**`campaigns`** — Phishing simulation campaigns with embedded schedule and denormalized stats:

```go
type Campaign struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    OrgID       primitive.ObjectID `bson:"org_id"`
    Name        string             `bson:"name"`
    Status      string             `bson:"status"` // draft, scheduled, sending, completed, cancelled
    Template    EmailTemplate      `bson:"template"`
    Schedule    Schedule           `bson:"schedule"`
    TargetIDs   []primitive.ObjectID `bson:"target_ids"`
    Stats       CampaignStats      `bson:"stats"`
    CreatedAt   time.Time          `bson:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at"`
}

type CampaignStats struct {
    TotalTargets int `bson:"total_targets"`
    Sent         int `bson:"sent"`
    Delivered    int `bson:"delivered"`
    Opened       int `bson:"opened"`
    Clicked      int `bson:"clicked"`
}
```

**`email_events`** — Event-sourcing style tracking log:

```go
type EmailEvent struct {
    ID         primitive.ObjectID `bson:"_id,omitempty"`
    CampaignID primitive.ObjectID `bson:"campaign_id"`
    EmployeeID primitive.ObjectID `bson:"employee_id"`
    Token      string             `bson:"token"` // unique per recipient
    EventType  string             `bson:"event_type"` // sent, delivered, opened, clicked
    Metadata   EventMetadata      `bson:"metadata,omitempty"`
    OccurredAt time.Time          `bson:"occurred_at"`
}

type EventMetadata struct {
    IP        string `bson:"ip,omitempty"`
    UserAgent string `bson:"user_agent,omitempty"`
    LinkURL   string `bson:"link_url,omitempty"`
}
```

**Index strategy:**

| Collection | Index | Type | Purpose |
|-----------|-------|------|---------|
| `organizations` | `{domain: 1}` | Unique | Prevent duplicate org domains |
| `employees` | `{org_id: 1, email: 1}` | Unique compound | Prevent duplicate employees per org |
| `campaigns` | `{org_id: 1, status: 1}` | Compound | List campaigns by org and status |
| `email_events` | `{token: 1}` | Unique | Token-based event lookup for tracking |
| `email_events` | `{campaign_id: 1, event_type: 1}` | Compound | Aggregation queries for campaign stats |
| `email_events` | `{occurred_at: 1}` | TTL (90 days) | Auto-delete old events for data retention/GDPR |

---

### Campaign Scheduling & Execution

Campaigns follow a state machine:

```
  draft ──→ scheduled ──→ sending ──→ completed
    │           │
    └───────────┴──→ cancelled
```

Valid transitions:
- `draft → scheduled`: when the admin sets a send window and launches
- `scheduled → sending`: when the first email task is dequeued
- `sending → completed`: when all emails have been sent
- `draft → cancelled` / `scheduled → cancelled`: admin cancellation
- `sending → cancelled` is allowed but already-sent emails cannot be recalled

**Randomized send times:** To simulate realistic phishing, emails are not sent all at once. A random time within a configurable window is computed for each recipient using `crypto/rand` (not `math/rand` — the randomness should be unpredictable to avoid patterns):

```go
func randomTimeInWindow(start, end time.Time) (time.Time, error) {
    window := end.Sub(start)
    nBig, err := rand.Int(rand.Reader, big.NewInt(int64(window)))
    if err != nil {
        return time.Time{}, err
    }
    return start.Add(time.Duration(nBig.Int64())), nil
}
```

Each send becomes an asynq task scheduled with `ProcessAt()`:

```go
task := asynq.NewTask("email:send", payload)
_, err := client.Enqueue(task, asynq.ProcessAt(sendTime))
```

**Worker pool concurrency** is configured via `asynq.Config.Concurrency` — typically 10–50 workers depending on SMTP throughput and rate limits.

---

### Email Sending

The email worker handles the full send pipeline:

1. **Idempotency check** — Query `email_events` for an existing `sent` event with this token. If found, skip (asynq may redeliver on worker crash).
2. **Rate limit check** — Sliding window rate limiter per organization and per domain to avoid SMTP throttling.
3. **Template rendering** — `html/template` renders the phishing email with per-recipient personalization (name, company). The template embeds:
   - A **tracking pixel**: `<img src="https://track.example.com/t/open/{token}" width="1" height="1" />`
   - **Tracked links**: All URLs are rewritten to `https://track.example.com/t/click/{token}?url={encoded_original_url}`
4. **SMTP send** — `net/smtp.SendMail` wrapped in a circuit breaker (`sony/gobreaker`). The breaker opens after N consecutive failures, preventing the worker from hammering a dead SMTP server.
5. **Record event** — Insert a `sent` event into `email_events`.
6. **Update Redis counters** — Increment campaign send count.

**Token generation** uses `crypto/rand` for 32-byte URL-safe tokens:

```go
func generateToken() (string, error) {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}
```

**Circuit breaker** configuration:

```go
cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
    Name:        "smtp",
    MaxRequests: 3,              // half-open: allow 3 test requests
    Interval:    60 * time.Second, // closed-state reset window
    Timeout:     30 * time.Second, // open → half-open after 30s
    ReadyToTrip: func(counts gobreaker.Counts) bool {
        return counts.ConsecutiveFailures > 5
    },
})
```

---

### Tracking & Analytics

**Tracking pixel endpoint** — `GET /t/open/:token`:
- Looks up the token in `email_events` to find the campaign/employee.
- Inserts an `opened` event with IP and user-agent metadata.
- Returns a 1x1 transparent GIF with `Content-Type: image/gif` and `Cache-Control: no-store`.
- The response is a hardcoded 43-byte GIF (smallest valid GIF).

**Link click endpoint** — `GET /t/click/:token`:
- Looks up the token, validates the `url` query parameter against an allowlist of original campaign URLs.
- Inserts a `clicked` event.
- Returns a `302 Found` redirect to the educational landing page (not the original phishing URL — the user sees a training message).

**Real-time counters in Redis:**

- **HyperLogLog** for unique opens/clicks per campaign. `PFADD campaign:{id}:unique_opens {employee_id}` — memory-efficient probabilistic counting (12 KB per key regardless of cardinality, 0.81% standard error).
- **INCR** for total event counts: `INCR campaign:{id}:total_opens`.
- **Pipelining** — both commands are sent in a single Redis round-trip.

**MongoDB aggregation for campaign reports:**

```go
pipeline := mongo.Pipeline{
    {{Key: "$match", Value: bson.M{"campaign_id": campaignID}}},
    {{Key: "$group", Value: bson.M{
        "_id":   "$event_type",
        "count": bson.M{"$sum": 1},
    }}},
}
```

Time-series bucketing for charts (events per hour):

```go
{{Key: "$group", Value: bson.M{
    "_id": bson.M{
        "event_type": "$event_type",
        "hour": bson.M{"$dateTrunc": bson.M{
            "date": "$occurred_at",
            "unit": "hour",
        }},
    },
    "count": bson.M{"$sum": 1},
}}}
```

---

### API Design

RESTful endpoints with cursor-based pagination:

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/v1/orgs` | Create organization |
| `GET` | `/api/v1/orgs/:id` | Get organization |
| `PUT` | `/api/v1/orgs/:id` | Update organization |
| `DELETE` | `/api/v1/orgs/:id` | Delete organization |
| `GET` | `/api/v1/orgs/:id/employees` | List employees (cursor paginated) |
| `POST` | `/api/v1/orgs/:id/employees` | Add employee |
| `POST` | `/api/v1/orgs/:id/employees/import` | Bulk CSV import |
| `DELETE` | `/api/v1/orgs/:id/employees/:eid` | Remove employee |
| `POST` | `/api/v1/orgs/:id/campaigns` | Create campaign |
| `GET` | `/api/v1/orgs/:id/campaigns` | List campaigns (cursor paginated) |
| `GET` | `/api/v1/orgs/:id/campaigns/:cid` | Get campaign details |
| `PUT` | `/api/v1/orgs/:id/campaigns/:cid` | Update campaign |
| `POST` | `/api/v1/orgs/:id/campaigns/:cid/launch` | Launch campaign |
| `POST` | `/api/v1/orgs/:id/campaigns/:cid/cancel` | Cancel campaign |
| `GET` | `/api/v1/orgs/:id/campaigns/:cid/analytics` | Campaign analytics |
| `GET` | `/t/open/:token` | Tracking pixel |
| `GET` | `/t/click/:token` | Link click tracker |

**Cursor-based pagination** using MongoDB `_id` (which is monotonically increasing):

```go
filter := bson.M{"org_id": orgID}
if cursor != "" {
    cursorID, _ := primitive.ObjectIDFromHex(cursor)
    filter["_id"] = bson.M{"$gt": cursorID}
}
opts := options.Find().SetLimit(int64(pageSize)).SetSort(bson.D{{Key: "_id", Value: 1}})
```

This avoids the performance problems of offset-based pagination (`skip` + `limit`) on large collections.

---

### Redis Usage Patterns

Seven patterns used throughout the platform:

**1. Job Queue (asynq)**

Asynq uses Redis streams under the hood. Tasks are enqueued with type and payload, processed by workers, and retried on failure:

```go
client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
task := asynq.NewTask("email:send", payload)
client.Enqueue(task, asynq.MaxRetry(3), asynq.ProcessAt(sendTime))
```

**2. Sliding Window Rate Limiting**

Using `go-redis/redis_rate` to enforce per-org send limits:

```go
limiter := redis_rate.NewLimiter(rdb)
res, err := limiter.Allow(ctx, fmt.Sprintf("rate:org:%s", orgID), redis_rate.PerHour(1000))
if res.Remaining == 0 {
    // back off until res.RetryAfter
}
```

**3. HyperLogLog for Unique Counts**

Probabilistic counting of unique opens/clicks with 12 KB per key:

```go
pipe := rdb.Pipeline()
pipe.PFAdd(ctx, fmt.Sprintf("campaign:%s:unique_opens", cid), employeeID)
pipe.Incr(ctx, fmt.Sprintf("campaign:%s:total_opens", cid))
pipe.Exec(ctx)
```

**4. Template Caching**

Cache rendered email templates to avoid repeated DB lookups:

```go
cached, err := rdb.Get(ctx, fmt.Sprintf("tmpl:%s", templateID)).Result()
if err == redis.Nil {
    tmpl := loadFromMongo(ctx, templateID)
    rdb.Set(ctx, fmt.Sprintf("tmpl:%s", templateID), tmpl, 15*time.Minute)
}
```

**5. Distributed Locks**

Using `bsm/redislock` to ensure only one scheduler instance processes a campaign at a time:

```go
lock, err := locker.Obtain(ctx, fmt.Sprintf("lock:campaign:%s", cid), 30*time.Second, nil)
if err == redislock.ErrNotObtained {
    return // another instance is handling this campaign
}
defer lock.Release(ctx)
```

**6. Real-Time Dashboard Counters**

Using Redis hashes for campaign progress:

```go
rdb.HIncrBy(ctx, fmt.Sprintf("progress:%s", cid), "sent", 1)
rdb.HIncrBy(ctx, fmt.Sprintf("progress:%s", cid), "delivered", 1)
// Dashboard reads: rdb.HGetAll(ctx, fmt.Sprintf("progress:%s", cid))
```

**7. Campaign Progress Tracking (HSET)**

Store campaign state transitions and worker progress:

```go
rdb.HSet(ctx, fmt.Sprintf("campaign:%s:state", cid),
    "status", "sending",
    "sent_count", 42,
    "total_count", 100,
    "last_sent_at", time.Now().Unix(),
)
```

---

### MongoDB Usage Patterns

**Embedding vs. separate collections:**

- **Embed** when the data is read together and bounded in size (template inside campaign, stats inside campaign).
- **Separate collection** when the data grows unboundedly (email events — one per recipient per event type) or needs independent querying.

**Compound index strategy:**

Design indexes around query patterns, not just fields. `{org_id: 1, status: 1}` on campaigns supports both "all campaigns for org X" and "all active campaigns for org X" efficiently. The field with higher selectivity usually comes first, but matching the query predicates matters more than selectivity ordering in MongoDB.

**TTL indexes for data retention/GDPR:**

```go
indexModel := mongo.IndexModel{
    Keys:    bson.D{{Key: "occurred_at", Value: 1}},
    Options: options.Index().SetExpireAfterSeconds(90 * 24 * 60 * 60), // 90 days
}
```

MongoDB automatically deletes documents when the indexed field passes the TTL threshold. This handles GDPR retention requirements without batch cleanup jobs.

**Aggregation pipelines:**

Used for campaign reports — group events by type, bucket by time period, and compute rates. Pipelines run server-side, avoiding transferring raw events to the application.

**Bulk writes with `ordered: false`:**

For CSV employee imports, use unordered bulk inserts so that a duplicate in row 50 doesn't block rows 51–1000:

```go
opts := options.BulkWrite().SetOrdered(false)
result, err := coll.BulkWrite(ctx, operations, opts)
// result.InsertedCount tells you how many succeeded
// Check err for individual write errors (duplicates)
```

**Change streams:**

MongoDB change streams enable real-time reactions to data changes (e.g., updating a dashboard when new events arrive):

```go
stream, _ := coll.Watch(ctx, mongo.Pipeline{})
for stream.Next(ctx) {
    var event bson.M
    stream.Decode(&event)
    // push update to websocket/SSE
}
```

---

### Production Concerns

**Idempotency:**
Every email send is guarded by a token-based deduplication check. Before sending, the worker queries for an existing `sent` event with the same token. If found, the task is acknowledged without re-sending. This handles asynq retries after worker crashes.

**Graceful worker shutdown:**
The asynq server supports graceful shutdown via `srv.Shutdown()`. On SIGTERM/SIGINT, workers finish in-progress tasks (up to a deadline) before exiting. No emails are half-sent.

**Circuit breaker for SMTP:**
The `sony/gobreaker` circuit breaker wraps SMTP calls. After N consecutive failures, the breaker opens and immediately rejects further sends for a cooldown period. This prevents a dead SMTP server from consuming all worker capacity with timeouts.

**OpenTelemetry instrumentation:**
- **Traces**: Span per email send (covering template render → SMTP → event record), span per tracking request.
- **Metrics**: `emails_sent_total`, `emails_failed_total`, `tracking_requests_total` (by event type), `smtp_circuit_breaker_state`, `campaign_send_duration_seconds`.

**Security considerations:**
- Tokens are generated with `crypto/rand` (not `math/rand`) — unpredictable, preventing token guessing.
- The platform never captures real credentials. The educational landing page explains that the email was a simulation.
- Tracking endpoints are rate-limited to prevent abuse (someone hammering the pixel endpoint to inflate stats).
- Redirect URLs in the click tracker are validated against the original campaign's URL list — no open redirect vulnerability.
- The tracking pixel uses `Cache-Control: no-store` to prevent cached images from suppressing open tracking.

---

## Interview Questions

1. **Why MongoDB over PostgreSQL for email tracking events?**

   Email events are append-heavy, high-volume, and document-shaped — each event type carries different metadata (an "opened" event has IP and user-agent; a "clicked" event has the link URL). MongoDB handles schema flexibility naturally with `bson` tags and embedded structs. The event data doesn't need relational joins — it's queried by campaign_id and aggregated. MongoDB's TTL indexes handle automatic data expiration for GDPR without custom batch jobs. Write throughput on a single collection with proper indexes is higher than inserting into a normalized relational schema. The tradeoff: you lose ACID transactions across collections (though MongoDB supports multi-document transactions when needed) and SQL's expressive query language for ad-hoc analysis.

2. **How would you design randomized but controllable send-time scheduling?**

   Define a send window (start time + end time) for the campaign. For each recipient, generate a random time within that window using `crypto/rand` (not `math/rand` — the randomness should be cryptographically unpredictable so employees can't predict timing patterns). Create an asynq task per recipient with `ProcessAt(randomTime)`. The scheduler produces all tasks upfront and asynq handles delayed execution via Redis sorted sets. Controllability comes from the window parameters: a 1-hour window produces a burst, a 5-day window spreads emails naturally. Cancellation works by removing pending tasks from the asynq queue.

3. **Which Redis data structure would you use for counting unique email opens, and why?**

   HyperLogLog (`PFADD` / `PFCOUNT`). It counts unique elements probabilistically using only 12 KB of memory per key, regardless of the cardinality. For unique opens, you `PFADD campaign:{id}:unique_opens {employee_id}` — adding the same employee_id multiple times doesn't increase the count. The tradeoff is a 0.81% standard error, which is acceptable for analytics dashboards. The alternative — a Redis Set with `SADD` — uses O(n) memory proportional to the number of unique elements, which becomes expensive for large campaigns. HyperLogLog also supports `PFMERGE` to combine counts across time windows.

4. **How would you prevent duplicate email sends when asynq retries a failed task?**

   Token-based idempotency. Each recipient in a campaign gets a unique token (generated with `crypto/rand`). Before sending, the worker queries MongoDB for an existing `sent` event with that token. If found, the task is acknowledged without re-sending. The token is the idempotency key — it's deterministic per campaign-recipient pair, so retries always check the same key. The `email_events` collection has a unique index on `token`, so even concurrent workers can't create duplicate `sent` events (the second insert would fail with a duplicate key error, which the worker handles gracefully by acknowledging the task).

5. **How would you design a campaign state machine and enforce valid transitions?**

   Define states as constants (`draft`, `scheduled`, `sending`, `completed`, `cancelled`) and valid transitions as an adjacency list. Use a method that validates the transition before updating:

   ```go
   var validTransitions = map[string][]string{
       "draft":     {"scheduled", "cancelled"},
       "scheduled": {"sending", "cancelled"},
       "sending":   {"completed", "cancelled"},
   }

   func (c *Campaign) TransitionTo(next string) error {
       allowed := validTransitions[c.Status]
       for _, s := range allowed {
           if s == next { c.Status = next; return nil }
       }
       return fmt.Errorf("invalid transition: %s → %s", c.Status, next)
   }
   ```

   In MongoDB, use a conditional update with a filter on the current status to prevent race conditions: `filter: {_id: id, status: currentStatus}`. If the update matches zero documents, the transition was invalid or another process already changed the state.

6. **How do tracking pixels work, and what are the privacy/security considerations?**

   A tracking pixel is a 1x1 transparent GIF embedded as an `<img>` tag in the email HTML. When the email client renders the image, it makes an HTTP request to the tracking server, which records the "opened" event along with IP and user-agent metadata. The server returns the tiny GIF with `Cache-Control: no-store` to prevent caching from suppressing future open events. Privacy considerations: many email clients now block remote images by default (Apple Mail Privacy Protection, Gmail image proxy), so open rates are inherently inaccurate — they undercount (blocked images) and sometimes overcount (pre-fetching proxies). Security: the tracking endpoint must be rate-limited to prevent abuse, and the token must be cryptographically random to prevent enumeration.

7. **How would you rate-limit outbound email sends per organization?**

   Use a sliding window rate limiter backed by Redis. The `go-redis/redis_rate` library implements this with sorted sets: each send records a timestamp, and the limiter counts entries within the window to determine if the limit is exceeded. Configure per-org limits (e.g., 1000 emails/hour for "pro" tier, 5000 for "enterprise"). The email worker checks the rate limiter before every send. If the limit is hit, the task is retried later (asynq supports retry with backoff). Also apply per-domain rate limiting (e.g., max 100 emails/hour to `@gmail.com`) to avoid triggering SMTP server throttling or blacklisting.

8. **How would you handle SMTP failures at scale, and what role does a circuit breaker play?**

   SMTP failures fall into two categories: transient (network timeouts, temporary server errors like 421) and permanent (invalid recipient, 550 errors). Transient failures are retried by asynq with exponential backoff. Permanent failures are recorded as `bounced` events and not retried. The circuit breaker (`sony/gobreaker`) wraps the SMTP call. After N consecutive failures (e.g., 5), the breaker opens and immediately returns an error for subsequent sends for a cooldown period (e.g., 30 seconds). This prevents all workers from blocking on a dead SMTP server with full TCP timeout waits. After the cooldown, the breaker enters half-open state and allows a few test requests. If they succeed, it closes and normal operation resumes. Without a circuit breaker, an SMTP outage would exhaust all worker goroutines waiting on TCP timeouts, halting the entire pipeline.

9. **How would you build a MongoDB aggregation pipeline for campaign analytics?**

   A two-stage pipeline: first `$match` by `campaign_id`, then `$group` by `event_type` with `$sum` to count each type. For time-series charts, add a nested `$group` using `$dateTrunc` to bucket events by hour or day:

   ```go
   pipeline := mongo.Pipeline{
       {{Key: "$match", Value: bson.M{"campaign_id": campaignID}}},
       {{Key: "$group", Value: bson.M{
           "_id": bson.M{
               "type": "$event_type",
               "hour": bson.M{"$dateTrunc": bson.M{"date": "$occurred_at", "unit": "hour"}},
           },
           "count": bson.M{"$sum": 1},
       }}},
       {{Key: "$sort", Value: bson.D{{Key: "_id.hour", Value: 1}}}},
   }
   ```

   This runs server-side in MongoDB, so only the aggregated results are transferred to the application. For real-time dashboard updates, combine this with Redis HyperLogLog counters for unique counts and `HGETALL` for progress tracking, avoiding repeated aggregation queries.

10. **What are the failure modes in a distributed email pipeline, and how do you mitigate each?**

    | Failure | Impact | Mitigation |
    |---------|--------|------------|
    | Worker crash mid-send | Email may or may not have been sent; task is redelivered | Token-based idempotency prevents duplicate sends |
    | SMTP server down | All sends fail, workers exhaust goroutines on timeouts | Circuit breaker opens after N failures, prevents cascading |
    | Redis down | No new tasks enqueued, rate limiting unavailable | Asynq has built-in Redis reconnection; fail-open or fail-closed on rate limits depends on policy |
    | MongoDB down | Events not recorded, tracking endpoints fail | Return 200 from tracking endpoints anyway (best-effort), buffer events in Redis for later flush |
    | Duplicate task delivery | Same email sent twice | Idempotency check (query for existing `sent` event by token) |
    | Scheduler crash | Campaign tasks not created | Recover by re-scanning campaigns in `scheduled` state on startup |
    | Network partition between worker and Redis | Task timeout, redelivery | Asynq's at-least-once semantics + idempotency |

---

## Resources

- mongo-go-driver: https://github.com/mongodb/mongo-go-driver
- go-redis/redis (v9): https://github.com/redis/go-redis
- hibiken/asynq: https://github.com/hibiken/asynq
- go-redis/redis_rate: https://github.com/go-redis/redis_rate
- bsm/redislock: https://github.com/bsm/redislock
- sony/gobreaker: https://github.com/sony/gobreaker
- go-gomail/gomail: https://github.com/go-gomail/gomail
- gophish (architectural reference): https://github.com/gophish/gophish
- MongoDB Aggregation Pipeline docs: https://www.mongodb.com/docs/manual/core/aggregation-pipeline/
- MongoDB TTL Indexes: https://www.mongodb.com/docs/manual/core/index-ttl/
- Redis HyperLogLog: https://redis.io/docs/latest/develop/data-types/probabilistic/hyperloglogs/
