package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Campaign scheduler: randomizes send times within a window and enqueues
// each email as an asynq task for delayed processing.
//
// Key interview concepts:
//   - crypto/rand for unpredictable randomness (employees can't predict timing)
//   - asynq.ProcessAt for delayed task execution via Redis sorted sets
//   - One task per recipient: enables individual retry and idempotency
//   - Distributed lock prevents duplicate scheduling if multiple scheduler
//     instances run concurrently (see redis_analytics.go for lock example)

// TaskTypeSendEmail is the asynq task type identifier.
const TaskTypeSendEmail = "email:send"

// SendEmailPayload is serialized as JSON in the asynq task body.
type SendEmailPayload struct {
	CampaignID string `json:"campaign_id"`
	EmployeeID string `json:"employee_id"`
	Token      string `json:"token"`
}

// randomTimeInWindow generates a uniformly random time between start and end
// using crypto/rand.
//
// Why crypto/rand instead of math/rand:
//   - math/rand is deterministic given the seed — if employees know the seed
//     (or the algorithm), they could predict when their email will arrive.
//   - crypto/rand reads from the OS entropy source (/dev/urandom on Linux),
//     producing unpredictable values.
//   - The performance difference is negligible for scheduling (we generate one
//     random number per recipient, not millions per second).
func randomTimeInWindow(start, end time.Time) (time.Time, error) {
	window := end.Sub(start)
	if window <= 0 {
		return time.Time{}, fmt.Errorf("end must be after start")
	}

	// rand.Int returns a uniform random value in [0, max).
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(window)))
	if err != nil {
		return time.Time{}, fmt.Errorf("generating random offset: %w", err)
	}

	return start.Add(time.Duration(nBig.Int64())), nil
}

// scheduleCampaign creates one asynq task per target employee, each scheduled
// at a random time within the campaign's send window.
//
// Flow:
//  1. For each target, generate a unique token (for tracking + idempotency).
//  2. Pick a random send time within the window.
//  3. Enqueue the task with ProcessAt so asynq delivers it at that time.
//
// asynq stores delayed tasks in a Redis sorted set (score = unix timestamp).
// A scheduler goroutine inside the asynq server moves tasks to the active
// queue when their time arrives.
func scheduleCampaign(
	ctx context.Context,
	client *asynq.Client,
	campaignID string,
	targetEmployeeIDs []string,
	windowStart, windowEnd time.Time,
) error {
	for _, empID := range targetEmployeeIDs {
		// Generate a unique token for this campaign-recipient pair.
		// This token is used for:
		//   1. Tracking pixel and link click URLs
		//   2. Idempotency key (prevent duplicate sends on retry)
		token, err := generateSchedulerToken()
		if err != nil {
			return fmt.Errorf("generating token for employee %s: %w", empID, err)
		}

		// Pick a random send time within the window.
		sendTime, err := randomTimeInWindow(windowStart, windowEnd)
		if err != nil {
			return fmt.Errorf("generating send time: %w", err)
		}

		// Build the task payload.
		payload, err := json.Marshal(SendEmailPayload{
			CampaignID: campaignID,
			EmployeeID: empID,
			Token:      token,
		})
		if err != nil {
			return fmt.Errorf("marshaling payload: %w", err)
		}

		// Enqueue with delayed processing.
		//   MaxRetry(3) — asynq retries up to 3 times on worker failure.
		//   ProcessAt(sendTime) — task won't be dequeued until this time.
		//   Queue("email") — route to a dedicated queue for email workers.
		task := asynq.NewTask(TaskTypeSendEmail, payload)
		info, err := client.EnqueueContext(ctx, task,
			asynq.MaxRetry(3),
			asynq.ProcessAt(sendTime),
			asynq.Queue("email"),
		)
		if err != nil {
			return fmt.Errorf("enqueueing task for employee %s: %w", empID, err)
		}

		fmt.Printf("  Enqueued task %s for employee %s at %s\n",
			info.ID, empID, sendTime.Format(time.RFC3339))
	}

	return nil
}

// generateSchedulerToken creates a 32-byte URL-safe random token.
// Used as the tracking/idempotency key for each recipient.
func generateSchedulerToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// Encode as hex (64 chars) for URL safety without padding issues.
	return fmt.Sprintf("%x", b), nil
}

func main() {
	fmt.Println("=== Campaign Scheduler ===\n")

	// Connect to Redis via asynq client.
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close()

	// Simulate a campaign with 5 target employees and a 2-hour send window.
	campaignID := primitive.NewObjectID().Hex()
	employees := []string{
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
	}

	windowStart := time.Now().Add(1 * time.Hour)
	windowEnd := windowStart.Add(2 * time.Hour)

	fmt.Printf("Campaign: %s\n", campaignID)
	fmt.Printf("Window:   %s to %s\n",
		windowStart.Format(time.RFC3339),
		windowEnd.Format(time.RFC3339))
	fmt.Printf("Targets:  %d employees\n\n", len(employees))

	ctx := context.Background()
	if err := scheduleCampaign(ctx, client, campaignID, employees, windowStart, windowEnd); err != nil {
		log.Fatal("scheduling campaign:", err)
	}

	fmt.Println("\nAll tasks enqueued successfully.")

	// Demonstrate the randomness distribution.
	fmt.Println("\n--- Random Time Distribution Demo ---")
	fmt.Println("Generating 10 random times in a 1-hour window:\n")

	base := time.Now()
	end := base.Add(1 * time.Hour)
	for i := 0; i < 10; i++ {
		t, _ := randomTimeInWindow(base, end)
		offset := t.Sub(base)
		fmt.Printf("  %2d. +%v\n", i+1, offset.Round(time.Second))
	}
}
