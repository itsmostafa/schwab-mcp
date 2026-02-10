package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/sony/gobreaker/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	redis_rate "github.com/go-redis/redis_rate/v10"
)

// Asynq email worker: dequeues send tasks, renders templates, sends via SMTP.
//
// Key interview concepts:
//   - Idempotency: check for existing "sent" event before sending to handle retries.
//   - Rate limiting: per-org sliding window limiter prevents SMTP throttling.
//   - Circuit breaker: wraps SMTP calls to fail fast when the server is down.
//   - Graceful shutdown: asynq server drains in-progress tasks on SIGTERM.
//   - Worker pool concurrency: configurable via asynq.Config.Concurrency.

// WorkerDeps holds dependencies injected into the email send handler.
type WorkerDeps struct {
	Events   *mongo.Collection
	RDB      *redis.Client
	Limiter  *redis_rate.Limiter
	Breaker  *gobreaker.CircuitBreaker[[]byte]
	SMTPAddr string
	SMTPFrom string
}

// HandleSendEmail processes an email:send task.
//
// Pipeline:
//  1. Deserialize payload (campaign ID, employee ID, token).
//  2. Idempotency check — skip if a "sent" event already exists for this token.
//  3. Rate limit check — back off if the org has exceeded its send limit.
//  4. Render the email template with per-recipient data.
//  5. Send via SMTP (wrapped in circuit breaker).
//  6. Record "sent" event in MongoDB.
//  7. Update Redis counters.
func (d *WorkerDeps) HandleSendEmail(ctx context.Context, t *asynq.Task) error {
	var p SendEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		// Permanent error — no point retrying with malformed payload.
		return fmt.Errorf("unmarshal payload: %w", err)
	}

	campaignID, _ := primitive.ObjectIDFromHex(p.CampaignID)
	employeeID, _ := primitive.ObjectIDFromHex(p.EmployeeID)

	// --- 1. Idempotency check ---
	// If a "sent" event with this token already exists, another worker (or a
	// previous attempt) already sent this email. Acknowledge the task.
	//
	// This prevents duplicate sends when asynq retries after a worker crash.
	// The unique index on token + event_type makes this check race-safe:
	// even if two workers check simultaneously and both find nothing, only
	// one InsertOne will succeed (the other gets a duplicate key error).
	count, err := d.Events.CountDocuments(ctx, bson.M{
		"token":      p.Token,
		"event_type": "sent",
	})
	if err != nil {
		return fmt.Errorf("idempotency check: %w", err)
	}
	if count > 0 {
		log.Printf("[skip] email already sent for token %s", p.Token)
		return nil // acknowledge — don't retry
	}

	// --- 2. Rate limit check ---
	// Sliding window rate limiter per organization. Prevents overwhelming
	// the SMTP server and getting blacklisted.
	//
	// redis_rate uses sorted sets: each send records a timestamp, and the
	// limiter counts entries within the window.
	res, err := d.Limiter.Allow(ctx,
		fmt.Sprintf("rate:org:%s", p.CampaignID), // key per org (simplified: using campaign ID)
		redis_rate.PerHour(1000),                  // 1000 emails per hour per org
	)
	if err != nil {
		return fmt.Errorf("rate limit check: %w", err)
	}
	if res.Remaining == 0 {
		// Rate limited — return error so asynq retries later with backoff.
		return fmt.Errorf("rate limited, retry after %v", res.RetryAfter)
	}

	// --- 3. Render email template ---
	// In production, the template would be loaded from the campaign document
	// (or Redis cache). Here we use a simplified inline template.
	emailBody, err := renderEmail(p.Token, "John Doe", "Acme Corp")
	if err != nil {
		return fmt.Errorf("render email: %w", err)
	}

	// --- 4. Send via SMTP (circuit breaker) ---
	// The circuit breaker wraps the SMTP call. If the SMTP server has failed
	// consecutively, the breaker opens and returns ErrOpenState immediately,
	// preventing workers from blocking on TCP timeouts.
	//
	// State transitions:
	//   Closed (normal) → Open (after 5 consecutive failures)
	//   Open → Half-Open (after 30s timeout)
	//   Half-Open → Closed (if 3 test requests succeed)
	//   Half-Open → Open (if any test request fails)
	recipient := fmt.Sprintf("employee_%s@example.com", employeeID.Hex()[:8])
	_, err = d.Breaker.Execute(func() ([]byte, error) {
		return nil, sendSMTP(d.SMTPAddr, d.SMTPFrom, recipient, "Security Update Required", emailBody)
	})
	if err != nil {
		// If the circuit is open, this returns immediately with ErrOpenState.
		// asynq will retry later when the breaker may have recovered.
		return fmt.Errorf("smtp send: %w", err)
	}

	// --- 5. Record "sent" event ---
	_, err = d.Events.InsertOne(ctx, bson.M{
		"campaign_id": campaignID,
		"employee_id": employeeID,
		"token":       p.Token,
		"event_type":  "sent",
		"occurred_at": time.Now(),
	})
	if err != nil {
		// If this fails with a duplicate key error, another worker beat us.
		// The email was still sent, but we won't double-count.
		if !mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("recording sent event: %w", err)
		}
	}

	// --- 6. Update Redis counters ---
	pipe := d.RDB.Pipeline()
	pipe.Incr(ctx, fmt.Sprintf("campaign:%s:total_sent", campaignID.Hex()))
	pipe.HIncrBy(ctx, fmt.Sprintf("progress:%s", campaignID.Hex()), "sent", 1)
	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("[warn] redis counter update failed: %v", err)
		// Non-fatal — counters are eventually consistent.
	}

	log.Printf("[sent] campaign=%s employee=%s token=%s",
		p.CampaignID, p.EmployeeID, p.Token[:16]+"...")

	return nil
}

// SendEmailPayload matches the scheduler's payload format.
type SendEmailPayload struct {
	CampaignID string `json:"campaign_id"`
	EmployeeID string `json:"employee_id"`
	Token      string `json:"token"`
}

// renderEmail renders a phishing simulation email with tracking pixel and links.
// Uses html/template for safe HTML rendering (auto-escapes user data).
func renderEmail(token, recipientName, orgName string) (string, error) {
	const tmpl = `<!DOCTYPE html>
<html>
<body>
<p>Dear {{.Name}},</p>
<p>Your {{.OrgName}} account requires immediate verification.</p>
<p><a href="https://track.example.com/t/click/{{.Token}}?url=https%3A%2F%2Fexample.com%2Fverify">
   Click here to verify your account
</a></p>
<p>Best regards,<br>IT Security Team</p>
<img src="https://track.example.com/t/open/{{.Token}}" width="1" height="1" alt="" />
</body>
</html>`

	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = t.Execute(&buf, map[string]string{
		"Name":    recipientName,
		"OrgName": orgName,
		"Token":   token,
	})
	return buf.String(), err
}

// sendSMTP sends an email via SMTP. In production, you'd use a library like
// gomail for MIME support, attachments, and connection pooling.
func sendSMTP(addr, from, to, subject, body string) error {
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s",
		from, to, subject, body)
	return smtp.SendMail(addr, nil, from, []string{to}, []byte(msg))
}

func main() {
	fmt.Println("=== Email Send Worker ===\n")

	// --- MongoDB connection ---
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("mongo connect:", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// --- Redis connection ---
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()

	// --- Circuit breaker ---
	// Settings explanation:
	//   MaxRequests: 3     — in half-open state, allow 3 test requests before deciding
	//   Interval:    60s   — in closed state, reset failure counts every 60s
	//   Timeout:     30s   — stay open for 30s before transitioning to half-open
	//   ReadyToTrip:       — open the circuit after 5 consecutive failures
	cb := gobreaker.NewCircuitBreaker[[]byte](gobreaker.Settings{
		Name:        "smtp",
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Printf("[circuit-breaker] %s: %s → %s", name, from, to)
		},
	})

	// --- Worker dependencies ---
	deps := &WorkerDeps{
		Events:   mongoClient.Database("phishing_platform").Collection("email_events"),
		RDB:      rdb,
		Limiter:  redis_rate.NewLimiter(rdb),
		Breaker:  cb,
		SMTPAddr: "localhost:1025", // MailHog or similar for local dev
		SMTPFrom: "security@acme.com",
	}

	// --- Asynq server ---
	// Concurrency: 20 workers process email tasks in parallel.
	// StrictPriority: the "email" queue is processed before "default".
	// ShutdownTimeout: workers get 30s to finish in-progress tasks on SIGTERM.
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 20,
			Queues: map[string]int{
				"email":   6, // priority weight
				"default": 1,
			},
			ShutdownTimeout: 30 * time.Second,
		},
	)

	// Register the handler for email:send tasks.
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", deps.HandleSendEmail)

	// --- Graceful shutdown ---
	// asynq.Server.Run blocks until it receives SIGTERM or SIGINT.
	// On signal:
	//   1. Stop accepting new tasks from Redis.
	//   2. Wait for in-progress tasks to complete (up to ShutdownTimeout).
	//   3. Return from Run.
	//
	// This ensures no emails are half-sent — the SMTP call either completes
	// or the task stays in Redis for retry.
	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatal("asynq server:", err)
		}
	}()

	fmt.Println("Worker started. Concurrency: 20")
	fmt.Println("Listening on queue: email")
	fmt.Println("Press Ctrl+C for graceful shutdown.\n")

	// Wait for shutdown signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("\nReceived %s, shutting down gracefully...\n", sig)

	srv.Shutdown()
	fmt.Println("Worker stopped.")
}
