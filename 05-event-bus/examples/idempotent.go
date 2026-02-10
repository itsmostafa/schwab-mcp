package main

import (
	"fmt"
	"sync"
)

// --- Idempotent Event Handler ---
//
// Interview point: in any at-least-once delivery system (NATS JetStream,
// Kafka, outbox relay), duplicate events WILL arrive. A consumer that
// processes the same event twice without guarding against it will produce
// incorrect state: double charges, duplicate records, repeated emails.
// Idempotency is not optional — it's a hard requirement.

// Event is a minimal event with an ID used as the idempotency key.
type Event struct {
	ID      string // The unique event identifier — this is the deduplication key.
	Payload string
}

// --- IdempotencyStore ---
//
// Tracks which event IDs have already been processed. In production,
// this would be a database table with a UNIQUE constraint on event_id:
//
//   CREATE TABLE processed_events (
//       event_id  VARCHAR PRIMARY KEY,
//       processed_at TIMESTAMPTZ DEFAULT NOW()
//   );
//
// Interview point: consider adding a TTL or retention policy to clean up
// old entries. Without cleanup, the store grows unboundedly. A 7-day TTL
// is common — events older than the broker's retention period will never
// be redelivered, so their IDs can be safely purged.

type IdempotencyStore struct {
	mu        sync.Mutex
	processed map[string]bool
}

func NewIdempotencyStore() *IdempotencyStore {
	return &IdempotencyStore{
		processed: make(map[string]bool),
	}
}

// TryClaimAndMark attempts to atomically claim an event ID for processing.
// Returns true if the caller is the first to claim it (should process),
// false if it was already claimed (duplicate — skip).
//
// Interview point: the check-and-mark MUST be atomic. If IsProcessed and
// MarkProcessed are separate calls, a TOCTOU race exists: two goroutines
// can both see "not processed" and both proceed. In production, a DB
// unique constraint (INSERT ... ON CONFLICT DO NOTHING) provides this
// atomicity — the INSERT either succeeds (you own it) or is a no-op.
//
// Note: we mark BEFORE processing here to prevent concurrent duplicates.
// The tradeoff is that if the handler crashes after claiming but before
// finishing, the event is "swallowed." To avoid this in production,
// use a two-phase approach: claim with a "pending" status, do work,
// then update to "done" — or do work + mark in one DB transaction.
func (s *IdempotencyStore) TryClaim(eventID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.processed[eventID] {
		return false
	}
	s.processed[eventID] = true
	return true
}

// --- Idempotent Handler Wrapper ---
//
// This is a decorator/middleware pattern: it wraps any handler function
// with deduplication logic. The inner handler only runs if the event
// hasn't been seen before.
//
// Interview point: this separation of concerns is important. Business
// logic handlers shouldn't need to know about deduplication — that's
// infrastructure. Middleware-based idempotency keeps handlers clean
// and testable.

type HandlerFunc func(Event) error

func WithIdempotency(store *IdempotencyStore, inner HandlerFunc) HandlerFunc {
	return func(e Event) error {
		// Atomically check-and-claim. If another goroutine already
		// claimed this event, we skip immediately.
		if !store.TryClaim(e.ID) {
			fmt.Printf("  [Idempotency] Event %s already processed — skipping (duplicate)\n", e.ID)
			return nil // Not an error — duplicates are expected, not exceptional.
		}

		// We own this event — run the actual business logic.
		if err := inner(e); err != nil {
			// Interview point: on failure, you may want to un-claim so
			// redelivery can retry. Omitted here for simplicity.
			return err
		}

		return nil
	}
}

// --- Demo ---

func main() {
	fmt.Println("=== Idempotent Handler Demo ===")
	fmt.Println()

	store := NewIdempotencyStore()

	// The actual business handler — this should only run once per event.
	orderHandler := func(e Event) error {
		fmt.Printf("  [OrderHandler] Processing event %s: %s\n", e.ID, e.Payload)
		return nil
	}

	// Wrap it with idempotency.
	safeHandler := WithIdempotency(store, orderHandler)

	// --- Scenario 1: First delivery — should process normally ---

	fmt.Println("Scenario 1: First delivery of event evt-001")
	err := safeHandler(Event{ID: "evt-001", Payload: "order 42 placed"})
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	fmt.Println()

	// --- Scenario 2: Duplicate delivery — should be skipped ---
	// Interview point: this is the common case in at-least-once systems.
	// The broker redelivers because the consumer crashed after processing
	// but before acknowledging. Without idempotency, the order would be
	// processed twice.

	fmt.Println("Scenario 2: Duplicate delivery of event evt-001")
	err = safeHandler(Event{ID: "evt-001", Payload: "order 42 placed"})
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	fmt.Println()

	// --- Scenario 3: Different event — should process normally ---

	fmt.Println("Scenario 3: New event evt-002")
	err = safeHandler(Event{ID: "evt-002", Payload: "order 43 placed"})
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	fmt.Println()

	// --- Scenario 4: Concurrent duplicate delivery ---
	// Simulates multiple goroutines trying to process the same event
	// simultaneously, as might happen with competing consumers or
	// a race between the original delivery and a redelivery.
	//
	// Interview point: the mutex in IdempotencyStore serializes access,
	// ensuring only one goroutine processes the event. In production,
	// you'd rely on a DB unique constraint — INSERT ... ON CONFLICT DO
	// NOTHING — which provides the same serialization at the database level.

	fmt.Println("Scenario 4: Concurrent duplicate delivery of event evt-003")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			evt := Event{ID: "evt-003", Payload: "order 44 placed"}
			handlerErr := safeHandler(evt)
			if handlerErr != nil {
				fmt.Printf("  [Worker %d] ERROR: %v\n", workerID, handlerErr)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println()
	fmt.Println("=== Summary ===")
	fmt.Println("- At-least-once delivery means duplicates WILL happen — handlers must be idempotent.")
	fmt.Println("- The idempotency key is the event ID (UUID); checked before processing.")
	fmt.Println("- Wrapper/decorator pattern keeps deduplication separate from business logic.")
	fmt.Println("- In production: use a DB table with UNIQUE(event_id) instead of in-memory map.")
	fmt.Println("- Consider a TTL to clean old entries (events past broker retention won't redeliver).")
	fmt.Println("- evt-003 was delivered 5 times concurrently but processed only once.")
}
