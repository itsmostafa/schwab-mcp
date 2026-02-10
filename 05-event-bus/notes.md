# Event Bus Pattern

## Key Concepts

### What the event bus pattern solves

- An event bus decouples producers from consumers:
  - Producer emits `OrderCreated` and does not know who handles it.
  - Consumers (billing, notifications, sync jobs) subscribe independently.
- This improves extensibility and team autonomy, but introduces eventual consistency and operational complexity.
- Interview framing:
  - choose an event bus when cross-service reactions are needed and tight request/response coupling becomes fragile.

### Event contract and envelope design

- Treat events as public contracts.
- Use a stable envelope and versioned payload:
  - `event_id` (UUID for dedupe)
  - `event_type` (for example `customer.updated.v1`)
  - `occurred_at` (UTC timestamp)
  - `source` (service name)
  - `correlation_id` / `trace_id`
  - `payload` (schema-versioned data)
- Prefer additive schema changes; avoid breaking deletes/renames.
- Keep event names as facts in past tense (`invoice.generated`), not commands.

### In-process event bus (single binary)

- Fast and simple for modular monoliths.
- Typical shape: `map[string][]Handler` protected by `sync.RWMutex`.
- Be explicit about semantics:
  - synchronous handlers: simple ordering, publisher blocked by slow handlers
  - async handlers: better isolation, must handle worker lifecycle and errors
- Useful starter implementation:

```go
type Event struct {
    Type string
    Data any
}

type Handler func(context.Context, Event) error

type Bus struct {
    mu   sync.RWMutex
    subs map[string][]Handler
}

func (b *Bus) Subscribe(eventType string, h Handler) {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.subs[eventType] = append(b.subs[eventType], h)
}

func (b *Bus) Publish(ctx context.Context, e Event) error {
    b.mu.RLock()
    handlers := append([]Handler(nil), b.subs[e.Type]...)
    b.mu.RUnlock()

    for _, h := range handlers {
        if err := h(ctx, e); err != nil {
            return err
        }
    }
    return nil
}
```

### Distributed bus options and tradeoffs

- Core NATS:
  - very low latency pub/sub
  - at-most-once delivery (no persistence by default)
  - queue subscriptions for load-balanced consumers
- NATS JetStream:
  - persistence + acknowledgments + redelivery
  - durable consumers, pull/push modes, replay options
  - commonly used for at-least-once workflows
- Kafka (for example `kafka-go`, Sarama clients):
  - partitioned log, high throughput, retention/replay
  - order guaranteed per partition, not globally
  - consumer groups scale reads and manage offsets
- RabbitMQ/AMQP-style queues:
  - flexible routing and queue semantics
  - strong fit for task distribution and work queues
  - operational model differs from log-based systems like Kafka

### Delivery semantics you should explain clearly

- At-most-once: no redelivery; can lose messages on failure.
- At-least-once: redelivery possible; duplicates must be handled.
- Exactly-once: usually scoped and conditional; in practice, design for idempotent consumers and call it effectively-once at business level.

### Ordering, partitioning, and keys

- Ordering is usually constrained:
  - NATS subjects: subscriber sees publish order per subject path, but distributed processing can change observed order.
  - Kafka: strict order per partition only.
- Choose partition/subject keys by aggregate identity (`customer_id`, `order_id`) to keep related events ordered where needed.
- If you need cross-aggregate total ordering, expect throughput tradeoffs.

### Consumer groups, flow control, and backpressure

- Consumer groups scale a subscription horizontally.
- Backpressure controls:
  - bounded internal channels
  - max in-flight messages
  - pull consumers + batch limits
  - visibility via lag metrics
- Prevent slow-consumer cascades by isolating heavy handlers into separate consumers.

### Error handling, retries, and dead-letter strategy

- Handle transient vs permanent failures differently.
- Standard policy:
  - retry with exponential backoff for transient failures
  - move poison messages to DLQ/poison topic after max attempts
  - include failure metadata (`attempt`, `error`, `last_failed_at`)
- Watermill pattern: middleware-driven retries + poison queue routing.

### Idempotency and deduplication

- Assume duplicates in at-least-once systems.
- Common techniques:
  - idempotency key table keyed by `event_id`
  - upsert-based writes
  - compare-and-swap with expected version
- Side effects (emails, external API calls) must also be idempotent or guarded.

### Transactional outbox (critical interview topic)

- Problem: DB commit succeeds but event publish fails (or vice versa), causing inconsistency.
- Solution:
  - write business data + outbox row in one DB transaction
  - separate relay publishes outbox rows to broker
  - mark row published only after broker ack
- Minimal relay flow:

```go
for {
    rows := repo.FetchPendingOutbox(ctx, 100)
    for _, row := range rows {
        if err := broker.Publish(ctx, row.Topic, row.Payload); err != nil {
            continue // retry later with backoff
        }
        _ = repo.MarkPublished(ctx, row.ID)
    }
    time.Sleep(200 * time.Millisecond)
}
```

- DeepWiki watermill docs highlight this via SQL pub/sub + `Forwarder` component.

### Testing event-driven systems

- Unit tests:
  - handler logic, idempotency behavior, retry classification
- Integration tests:
  - real broker in containers, verify ack/retry/DLQ flows
- Contract tests:
  - validate event schema compatibility between producer and consumer
- Failure-path tests:
  - broker unavailable, consumer crash mid-processing, duplicate delivery
- In Go, prefer table-driven tests for handler cases and deterministic fake clocks for retry logic.

## Interview Questions

1. **What problem does an event bus solve compared to direct HTTP calls between services?**

   Direct HTTP calls create tight coupling: the caller must know the address, API shape, and availability of every downstream service. If billing, notifications, and sync all need to react to `OrderCreated`, the order service ends up with explicit calls to each, making it fragile and hard to extend. An event bus inverts this: the producer emits a single event and is unaware of consumers. New subscribers can be added without changing the producer. This improves team autonomy and extensibility. The tradeoff is eventual consistency — consumers process events asynchronously, so downstream state lags behind the source.

2. **What is the difference between pub/sub and queue semantics?**

   In pub/sub, every subscriber receives a copy of each message. If three services subscribe to `order.created`, all three get every event independently. In queue semantics, messages are distributed across competing consumers so that each message is processed by exactly one worker. NATS implements this via queue groups, Kafka via consumer groups, and RabbitMQ via queues with multiple consumers. Pub/sub is for fan-out (multiple independent reactions), while queue semantics are for load-balanced processing of a single workload.

3. **Explain at-most-once vs at-least-once delivery with examples.**

   At-most-once means a message is delivered zero or one times. If the broker or consumer crashes before processing completes, the message is lost. Core NATS works this way — it fires and forgets with no persistence. At-least-once means the broker redelivers until it gets an acknowledgment. If a consumer processes the message but crashes before acking, the message is redelivered, producing a duplicate. NATS JetStream and Kafka both provide at-least-once semantics. The practical consequence is that at-least-once systems require idempotent consumers.

4. **Why is idempotency mandatory in event-driven systems?**

   At-least-once delivery guarantees that duplicates will eventually occur — a consumer can crash after processing but before acknowledging, causing redelivery. Without idempotency, duplicates cause incorrect state: double charges, duplicate emails, or inconsistent data. Common techniques include keeping an idempotency key table keyed by `event_id`, using upserts instead of inserts, and compare-and-swap with expected versions. Side effects like external API calls must also be guarded, either by making them naturally idempotent or by checking a "processed" flag before executing.

5. **How do you design an event envelope for long-term compatibility?**

   Use a stable outer envelope with metadata fields that never change shape: `event_id` (UUID for deduplication), `event_type` (e.g., `customer.updated.v1`), `occurred_at` (UTC timestamp), `source` (originating service), `correlation_id`/`trace_id` (for distributed tracing), and a `payload` field containing the versioned domain data. The envelope schema stays fixed; only the payload evolves. Name events as past-tense facts (`invoice.generated`), not commands, since they describe something that already happened.

6. **How do you version events without breaking existing consumers?**

   Prefer additive-only changes: add new fields with defaults, never rename or remove existing fields. Consumers should ignore unknown fields. When a breaking change is unavoidable, introduce a new event type with an incremented version suffix (e.g., `order.created.v2`) and run both versions in parallel during a migration window. Producers emit both v1 and v2, consumers migrate at their own pace, and v1 is retired once all consumers have moved. Schema registries (like Confluent's for Kafka) can enforce compatibility checks at publish time.

7. **What are queue groups/consumer groups and why do they matter?**

   They allow multiple instances of a consumer to share the processing load for a subscription. In NATS, subscribers joining the same queue group receive messages in a load-balanced fashion — each message goes to exactly one member. In Kafka, a consumer group assigns partitions across members, so each partition is read by one consumer in the group. This enables horizontal scaling: you add more instances to increase throughput. It also provides fault tolerance — when one member dies, its work is redistributed to the remaining members.

8. **How does ordering differ between Kafka partitions and general pub/sub systems?**

   Kafka guarantees strict ordering within a single partition. Messages with the same partition key (e.g., `customer_id`) always land in the same partition and are read in order. Across partitions, there is no ordering guarantee. In general pub/sub systems like NATS, a single subscriber sees messages in publish order on a given subject, but once you add queue groups or multiple subscribers, observed order can vary because different workers process messages at different speeds. If you need related events ordered (all events for one customer), use a partition/subject key tied to the aggregate identity.

9. **What is backpressure, and how do you protect a consumer from overload?**

   Backpressure is when a consumer cannot keep up with the rate of incoming messages. Without protection, this leads to unbounded memory growth, cascading timeouts, or dropped messages. Protection strategies include: bounded internal channels (so goroutines block when the buffer is full), configuring max in-flight messages on the subscription, using pull-based consumers with batch size limits (JetStream pull consumers), and isolating slow handlers into separate consumer groups so they don't block faster ones. Monitor consumer lag metrics to detect backpressure before it becomes critical.

10. **How would you design retries and dead-letter handling?**

    Classify errors as transient (network timeout, temporary unavailability) or permanent (malformed payload, business rule violation). For transient errors, retry with exponential backoff and jitter to avoid thundering herds. Set a max retry count (e.g., 3–5 attempts). After exhausting retries, route the message to a dead-letter queue (DLQ) with failure metadata attached: `attempt` count, `error` message, and `last_failed_at` timestamp. Permanent errors should go to the DLQ immediately without retrying. The DLQ should be monitored and have tooling for inspection, replay after fixes, and manual resolution. Watermill implements this via middleware-driven retries and a poison queue router.

11. **What is a poison message, and how should it be handled operationally?**

    A poison message is one that consistently causes consumer failures — for example, due to a malformed payload, an unexpected schema version, or a bug in the handler. Without intervention, it blocks the consumer in an infinite retry loop. The standard approach is to track retry attempts and move the message to a DLQ/poison topic after a configured max. Operationally, you need: alerting when messages land in the DLQ, tooling to inspect the message and its failure metadata, the ability to replay messages after deploying a fix, and a process for manual resolution of truly unprocessable messages.

12. **Describe the transactional outbox pattern and which failure it prevents.**

    The outbox pattern prevents the dual-write problem: when a service needs to both commit a database change and publish an event, one can succeed while the other fails, causing inconsistency. The solution is to write both the business data and an outbox row in a single database transaction. A separate relay process polls the outbox table, publishes each pending row to the message broker, and marks the row as published only after receiving a broker acknowledgment. This guarantees that if the business data is committed, the event will eventually be published. The relay must be idempotent — if it crashes after publishing but before marking, it will re-publish on restart, so consumers must handle duplicates.

13. **Why is exactly-once often a misleading claim in distributed systems?**

    True exactly-once delivery across independent systems is impossible in the presence of network partitions and crashes — this is a consequence of the Two Generals problem. What systems like Kafka transactions provide is exactly-once *within a scoped boundary* (e.g., consume-transform-produce within Kafka). Once you cross system boundaries (write to a database, call an external API), you're back to at-least-once with the possibility of duplicates. The practical approach is to design for at-least-once delivery with idempotent consumers and call the result "effectively-once" at the business level.

14. **When would you choose NATS Core vs JetStream?**

    Use NATS Core when you need very low latency pub/sub and can tolerate message loss — fire-and-forget telemetry, real-time UI updates, or cache invalidation signals where missing one message is acceptable. Use JetStream when you need persistence, acknowledgments, and redelivery — order processing, data synchronization, or any workflow where losing a message means data inconsistency. JetStream adds durable consumers, replay from a point in time, pull/push delivery modes, and at-least-once guarantees. The tradeoff is added operational complexity and slightly higher latency.

15. **When would Kafka be a better fit than NATS or RabbitMQ-style queues?**

    Kafka excels when you need: high-throughput event streaming (millions of messages/sec), long-term retention and replay (consumers can re-read the log from any offset), strict per-partition ordering, and multiple independent consumer groups reading the same topic at different speeds. It's ideal for event sourcing, change data capture, and analytics pipelines. NATS is simpler operationally and better for low-latency request-reply or lightweight pub/sub. RabbitMQ fits when you need flexible routing (topic/fanout/header exchanges), per-message TTLs, priority queues, or traditional task distribution. Kafka's partitioned log model is fundamentally different from RabbitMQ's queue model.

16. **How do you monitor an event pipeline in production?**

    Key areas to monitor: publish rate and error rate per topic, consumer lag (how far behind consumers are from the latest message), consumer processing latency (time from event creation to processing completion), DLQ depth and growth rate, broker health (node availability, disk usage, replication status), and end-to-end latency using correlation IDs. Use structured logging with `event_id`, `correlation_id`, and `trace_id` for traceability. Set up alerts on consumer lag thresholds, DLQ growth, and publish failures. OpenTelemetry spans across publish and consume give visibility into the full event flow.

17. **What metrics indicate consumer health and lag risk?**

    Consumer lag (the offset difference between the latest published message and the consumer's current position) is the primary indicator. Rising lag means the consumer isn't keeping up. Other signals: processing latency per message (increasing latency suggests resource pressure), error rate and retry rate (high retries consume capacity), DLQ ingestion rate (rising poison messages suggest a systemic issue), and consumer rebalancing frequency (frequent rebalances in Kafka indicate instability). Memory and CPU usage of consumer processes complete the picture. Alert on sustained lag growth, not momentary spikes.

18. **How do you test event handlers for duplicates and reordering?**

    For duplicates: write table-driven unit tests that publish the same `event_id` twice and assert the handler produces the correct state (e.g., only one database row, no duplicate side effects). Test the idempotency key table or upsert logic directly. For reordering: publish events with a known sequence (e.g., `created` then `updated` then `deleted`) and deliver them out of order to verify the handler either rejects stale events (using version numbers) or converges to the correct state. Integration tests with a real broker in containers verify ack/retry/DLQ flows end-to-end. Use deterministic fake clocks for retry timing tests.

19. **Where should correlation IDs and trace IDs live in event flows?**

    They belong in the event envelope metadata, not in the payload. The producing service generates or propagates a `correlation_id` (ties together all events in a business workflow) and a `trace_id` (from OpenTelemetry or similar, ties together distributed traces). When a consumer processes an event, it extracts these IDs and sets them on its own context so that any downstream operations (database calls, HTTP requests, further events) carry the same identifiers. This enables end-to-end tracing across services. In NATS, these are typically set as message headers; in Kafka, as record headers.

20. **How would you roll out a new event field safely across many consumers?**

    Use an additive approach: add the new field to the payload with a sensible default or as optional (nullable). Consumers that haven't been updated will simply ignore the unknown field. Roll out in phases: (1) deploy updated consumers that can handle the new field but don't require it, (2) deploy the updated producer that starts populating the field, (3) once all consumers are updated and confirmed working, optionally make the field required in documentation. If the field replaces an old one, keep both during a deprecation window and emit both values until all consumers have migrated. Never remove the old field until you've confirmed zero consumers depend on it.

Practice prompts:
- Implement a simple in-process bus with `Subscribe`, `Publish`, and graceful shutdown.
- Build an outbox table + relay worker and prove no data/event mismatch under crash tests.
- Compare one use case implemented with NATS JetStream vs Kafka and explain tradeoffs.

## Resources

- NATS docs (concepts): https://docs.nats.io/nats-concepts
- NATS JetStream docs: https://docs.nats.io/nats-concepts/jetstream
- `nats.go` repository: https://github.com/nats-io/nats.go
- Kafka concepts (official): https://kafka.apache.org/documentation/
- `segmentio/kafka-go` repository: https://github.com/segmentio/kafka-go
- `IBM/sarama` repository: https://github.com/IBM/sarama
- RabbitMQ docs (concepts): https://www.rabbitmq.com/docs
- Watermill docs: https://watermill.io/
- Transactional Outbox pattern: https://microservices.io/patterns/data/transactional-outbox.html

DeepWiki sources used for this note:
- `nats-io/nats-server` wiki: https://deepwiki.com/nats-io/nats-server/wiki
- `nats-io/nats.go` wiki: https://deepwiki.com/nats-io/nats.go/wiki
- `segmentio/kafka-go` wiki: https://deepwiki.com/segmentio/kafka-go/wiki
- `IBM/sarama` wiki: https://deepwiki.com/IBM/sarama/wiki
- `ThreeDotsLabs/watermill` wiki: https://deepwiki.com/ThreeDotsLabs/watermill/wiki
