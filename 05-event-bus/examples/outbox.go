package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// --- Transactional Outbox Pattern ---
//
// Interview point: the outbox pattern solves the "dual-write" problem.
// When a service needs to commit a DB change AND publish an event,
// doing them as two separate operations is unsafe — one can succeed
// while the other fails, leaving the system in an inconsistent state.
//
// Solution: write the business data AND an outbox row in a SINGLE
// database transaction. A separate relay process polls the outbox,
// publishes events to the broker, and marks them as published.
// This guarantees: if the business data is committed, the event
// WILL eventually be published (at-least-once).

// OutboxEntry represents a row in the outbox table.
// In production this would be a database table with columns matching
// these fields, typically with an index on (published, created_at).
type OutboxEntry struct {
	ID        string
	EventType string
	Payload   string    // JSON-serialized event payload.
	CreatedAt time.Time
	Published bool      // Relay sets this to true after successful broker publish.
}

// Order is a simple domain entity for demonstration.
type Order struct {
	ID     string
	Item   string
	Amount float64
}

// --- Simulated Database ---
// In production, orders and outbox entries live in the SAME database
// so they can be written in a single transaction.

type Database struct {
	mu      sync.Mutex
	orders  []Order
	outbox  []OutboxEntry
	nextID  int
}

func NewDatabase() *Database {
	return &Database{}
}

// CreateOrderWithOutbox simulates the critical operation: writing
// business data and an outbox entry atomically in one transaction.
//
// Interview point: the key insight is that both writes happen in the
// SAME database transaction. If the transaction fails, neither the
// order nor the outbox entry is persisted. If it succeeds, the relay
// is guaranteed to eventually find and publish the event. This
// eliminates the dual-write failure mode.
func (db *Database) CreateOrderWithOutbox(order Order) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// --- BEGIN TRANSACTION (simulated) ---

	// 1. Write the business entity.
	db.orders = append(db.orders, order)

	// 2. Write the outbox entry in the same transaction.
	payload, err := json.Marshal(map[string]any{
		"order_id": order.ID,
		"item":     order.Item,
		"amount":   order.Amount,
	})
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	db.nextID++
	entry := OutboxEntry{
		ID:        fmt.Sprintf("outbox-%d", db.nextID),
		EventType: "order.placed",
		Payload:   string(payload),
		CreatedAt: time.Now().UTC(),
		Published: false,
	}
	db.outbox = append(db.outbox, entry)

	// --- COMMIT TRANSACTION (simulated) ---

	fmt.Printf("  [DB] Committed order %s and outbox entry %s in one transaction\n",
		order.ID, entry.ID)
	return nil
}

// FetchUnpublished returns outbox entries that haven't been relayed yet.
// In production: SELECT * FROM outbox WHERE published = false ORDER BY created_at LIMIT N
func (db *Database) FetchUnpublished() []OutboxEntry {
	db.mu.Lock()
	defer db.mu.Unlock()

	var pending []OutboxEntry
	for _, e := range db.outbox {
		if !e.Published {
			pending = append(pending, e)
		}
	}
	return pending
}

// MarkPublished sets the published flag on an outbox entry.
// In production: UPDATE outbox SET published = true WHERE id = ?
//
// Interview point: the relay marks the entry ONLY after the broker
// confirms receipt. If the relay crashes between publishing and
// marking, the entry stays unpublished and will be re-published
// on the next poll. This is why consumers must be idempotent —
// the outbox guarantees at-least-once, not exactly-once.
func (db *Database) MarkPublished(id string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i := range db.outbox {
		if db.outbox[i].ID == id {
			db.outbox[i].Published = true
			return
		}
	}
}

// --- Simple Event Bus (for the relay to publish into) ---

type Handler func(eventType, payload string)

type EventBus struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[string][]Handler)}
}

func (b *EventBus) Subscribe(eventType string, h Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventType] = append(b.handlers[eventType], h)
}

func (b *EventBus) Publish(eventType, payload string) {
	b.mu.RLock()
	hs := make([]Handler, len(b.handlers[eventType]))
	copy(hs, b.handlers[eventType])
	b.mu.RUnlock()

	for _, h := range hs {
		h(eventType, payload)
	}
}

// --- Outbox Relay ---
//
// The relay is a background goroutine that polls the outbox table,
// publishes pending entries to the event bus (broker), and marks
// them as published.
//
// Interview point: the relay runs asynchronously and introduces a
// small delay between the DB commit and event delivery. This is
// the tradeoff for atomicity — you get eventual consistency instead
// of immediate delivery, but you eliminate the dual-write risk.

func startRelay(db *Database, bus *EventBus, pollInterval time.Duration, done chan struct{}) {
	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}

			pending := db.FetchUnpublished()
			for _, entry := range pending {
				fmt.Printf("  [Relay] Publishing outbox entry %s (type: %s)\n",
					entry.ID, entry.EventType)

				// Publish to the event bus (simulating broker publish).
				bus.Publish(entry.EventType, entry.Payload)

				// Mark as published ONLY after successful publish.
				// Interview point: if we crash here, we'll re-publish
				// on the next poll cycle. This is at-least-once delivery.
				db.MarkPublished(entry.ID)
				fmt.Printf("  [Relay] Marked %s as published\n", entry.ID)
			}

			time.Sleep(pollInterval)
		}
	}()
}

// --- Demo ---

func main() {
	fmt.Println("=== Transactional Outbox Pattern Demo ===")
	fmt.Println()

	db := NewDatabase()
	bus := NewEventBus()

	// Subscribe a handler that simulates downstream processing.
	bus.Subscribe("order.placed", func(eventType, payload string) {
		fmt.Printf("  [Handler] Received %s event — payload: %s\n", eventType, payload)
	})

	// Step 1: Create orders (business operation + outbox write, atomic).
	fmt.Println("Step 1: Creating orders (atomic DB write + outbox entry)")
	fmt.Println()

	_ = db.CreateOrderWithOutbox(Order{ID: "order-1", Item: "Widget", Amount: 29.99})
	_ = db.CreateOrderWithOutbox(Order{ID: "order-2", Item: "Gadget", Amount: 49.99})

	fmt.Println()

	// At this point, orders exist in the DB and outbox entries are pending.
	// No events have been published yet — that's the relay's job.
	fmt.Println("Step 2: Outbox entries are pending, no events published yet")
	pending := db.FetchUnpublished()
	fmt.Printf("  Unpublished outbox entries: %d\n", len(pending))

	fmt.Println()

	// Step 3: Start the relay — it will find and publish the pending entries.
	fmt.Println("Step 3: Starting relay (polls every 100ms)")
	fmt.Println()

	done := make(chan struct{})
	startRelay(db, bus, 100*time.Millisecond, done)

	// Give the relay time to pick up and process the outbox entries.
	time.Sleep(300 * time.Millisecond)

	fmt.Println()

	// Step 4: Verify everything was published.
	fmt.Println("Step 4: Checking outbox after relay processed entries")
	remaining := db.FetchUnpublished()
	fmt.Printf("  Unpublished outbox entries remaining: %d\n", len(remaining))

	fmt.Println()

	// Step 5: Create another order while relay is running — it picks it up automatically.
	fmt.Println("Step 5: Creating another order while relay is running")
	fmt.Println()

	_ = db.CreateOrderWithOutbox(Order{ID: "order-3", Item: "Doohickey", Amount: 9.99})

	// Give the relay time to pick up the new entry.
	time.Sleep(200 * time.Millisecond)

	// Shut down the relay.
	close(done)

	fmt.Println()
	fmt.Println("=== Summary ===")
	fmt.Println("- Business data and outbox entry are written in one atomic transaction.")
	fmt.Println("- Relay asynchronously polls and publishes, then marks entries as done.")
	fmt.Println("- If relay crashes after publish but before marking, it re-publishes (at-least-once).")
	fmt.Println("- Consumers MUST be idempotent to handle potential duplicate delivery.")
	fmt.Println("- This eliminates the dual-write problem: no event without data, no data without event.")
}
