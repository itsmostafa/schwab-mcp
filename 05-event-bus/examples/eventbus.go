package main

import (
	"fmt"
	"sync"
	"time"
)

// --- Event Envelope ---
// The envelope is a standard pattern in event-driven systems.
// It separates routing metadata from domain-specific payload,
// so infrastructure code (bus, relay, middleware) can inspect
// metadata without knowing anything about the payload schema.

// Event represents the canonical event envelope.
// Interview point: a well-designed envelope includes an ID for
// deduplication, a type for routing, a timestamp for ordering
// insights, a source for provenance, and a correlation ID for
// distributed tracing across services.
type Event struct {
	ID            string    // Unique event ID (UUID in production) — used as the idempotency key by consumers.
	Type          string    // Dot-namespaced past-tense fact, e.g. "order.placed" — NOT a command.
	OccurredAt    time.Time // When the event actually happened (UTC). Useful for debugging lag.
	Source        string    // Originating service or component. Helps trace where events come from.
	CorrelationID string    // Ties together all events in a single business workflow / request chain.
	Payload       any       // The versioned domain data. Kept as `any` for flexibility; in production, use typed schemas.
}

// Handler is the subscriber callback signature.
// Returning an error lets the bus (or middleware) decide on retry/DLQ strategy.
type Handler func(Event) error

// --- EventBus ---
// Interview point: an in-process bus is the simplest event-driven
// architecture. It's great for modular monoliths where you want
// loose coupling between packages without network overhead. The
// tradeoff: events don't survive process crashes — if the process
// dies mid-publish, pending handlers never run and the event is lost.

// EventBus provides synchronous, in-process pub/sub.
type EventBus struct {
	mu   sync.RWMutex          // Protects concurrent access to the subscribers map.
	subs map[string][]Handler  // Maps event type -> slice of registered handlers.
}

// NewEventBus creates a ready-to-use bus.
func NewEventBus() *EventBus {
	return &EventBus{
		subs: make(map[string][]Handler),
	}
}

// Subscribe registers a handler for a given event type.
// Multiple handlers can subscribe to the same event type —
// this is the fan-out pattern within a single process.
func (b *EventBus) Subscribe(eventType string, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subs[eventType] = append(b.subs[eventType], handler)
}

// Publish synchronously calls every handler registered for the event's type.
//
// Interview point: synchronous delivery means the publisher BLOCKS
// until ALL handlers complete. This gives strong ordering guarantees
// (handlers run sequentially) but a slow handler delays the publisher
// and all subsequent handlers. In production you'd consider async
// delivery with bounded worker pools for isolation.
func (b *EventBus) Publish(event Event) error {
	// Take a read lock to copy the handler slice, then release it
	// before calling handlers. This avoids holding the lock during
	// potentially slow handler execution and prevents deadlock if
	// a handler tries to subscribe.
	b.mu.RLock()
	handlers := make([]Handler, len(b.subs[event.Type]))
	copy(handlers, b.subs[event.Type])
	b.mu.RUnlock()

	for i, h := range handlers {
		if err := h(event); err != nil {
			// Interview point: failing fast on the first error is one
			// strategy. Alternatives include collecting all errors, or
			// skipping failed handlers and continuing. The right choice
			// depends on whether handlers are independent or ordered.
			return fmt.Errorf("handler %d for %q failed: %w", i, event.Type, err)
		}
	}
	return nil
}

// --- Demo ---

func main() {
	fmt.Println("=== In-Process Event Bus Demo ===")
	fmt.Println()

	bus := NewEventBus()

	// --- Subscribe multiple handlers to "user.created" ---
	// Interview point: fan-out within a process lets separate modules
	// react independently. Adding a new subscriber doesn't change the
	// publisher at all — that's the core decoupling benefit.

	bus.Subscribe("user.created", func(e Event) error {
		fmt.Printf("  [WelcomeEmail handler] Sending welcome email for event %s (correlation: %s)\n",
			e.ID, e.CorrelationID)
		fmt.Printf("    Payload: %v\n", e.Payload)
		return nil
	})

	bus.Subscribe("user.created", func(e Event) error {
		fmt.Printf("  [AuditLog handler] Recording audit entry for event %s from %s\n",
			e.ID, e.Source)
		return nil
	})

	// --- Subscribe to "order.placed" ---

	bus.Subscribe("order.placed", func(e Event) error {
		fmt.Printf("  [Inventory handler] Reserving stock for event %s (occurred at: %s)\n",
			e.ID, e.OccurredAt.Format(time.RFC3339))
		fmt.Printf("    Payload: %v\n", e.Payload)
		return nil
	})

	// --- Publish "user.created" ---
	// Both handlers will be called synchronously in subscription order.

	fmt.Println("Publishing user.created ...")
	err := bus.Publish(Event{
		ID:            "evt-001",
		Type:          "user.created",
		OccurredAt:    time.Now().UTC(),
		Source:        "user-service",
		CorrelationID: "req-abc-123", // Traces this event back to the originating API request.
		Payload: map[string]string{
			"user_id": "u-42",
			"email":   "alice@example.com",
		},
	})
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	fmt.Println()

	// --- Publish "order.placed" ---

	fmt.Println("Publishing order.placed ...")
	err = bus.Publish(Event{
		ID:            "evt-002",
		Type:          "order.placed",
		OccurredAt:    time.Now().UTC(),
		Source:        "order-service",
		CorrelationID: "req-def-456",
		Payload: map[string]string{
			"order_id": "o-99",
			"item":     "widget",
			"qty":      "3",
		},
	})
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}

	fmt.Println()

	// --- Publish an event with no subscribers ---
	// Interview point: publishing to an event type with no subscribers
	// is a no-op. This is by design — the producer doesn't need to know
	// or care whether anyone is listening.

	fmt.Println("Publishing payment.received (no subscribers) ...")
	err = bus.Publish(Event{
		ID:            "evt-003",
		Type:          "payment.received",
		OccurredAt:    time.Now().UTC(),
		Source:        "payment-service",
		CorrelationID: "req-ghi-789",
		Payload:       nil,
	})
	if err != nil {
		fmt.Printf("  ERROR: %v\n", err)
	}
	fmt.Println("  (no handlers called — this is expected)")

	fmt.Println()
	fmt.Println("=== Summary ===")
	fmt.Println("- Sync delivery: publisher blocked until all handlers finished.")
	fmt.Println("- Fan-out: user.created triggered 2 handlers from one publish.")
	fmt.Println("- Correlation ID travels with the event for end-to-end tracing.")
	fmt.Println("- In-process bus is simple but events are lost if the process crashes.")
}
