package main

import (
	"context"
	"fmt"
	"time"
)

// Context: WithCancel, WithTimeout, WithValue, and cancellation propagation.
//
// Key interview points:
//   - context.Context forms a tree: cancelling a parent cancels all children.
//   - ctx.Err() returns nil while active, context.Canceled after cancel(),
//     and context.DeadlineExceeded after a timeout/deadline expires.
//   - Always call the cancel func (even if the context times out) to release
//     resources. Failing to call cancel() leaks the context's internal timer.
//   - Convention: ctx is always the first parameter of a function.
//   - WithValue is for request-scoped data (trace IDs, auth tokens), NOT for
//     passing optional parameters — that's a code smell.

func main() {
	withCancelDemo()
	withTimeoutDemo()
	withValueDemo()
	cancellationPropagationDemo()
}

// withCancelDemo shows explicit cancellation: the parent calls cancel()
// and the child goroutine observes it via ctx.Done().
func withCancelDemo() {
	fmt.Println("=== WithCancel ===")

	// context.Background() is the root of any context tree.
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// ctx.Err() tells us *why* the context ended.
				fmt.Println("  worker stopped:", ctx.Err())
				return
			case <-time.After(100 * time.Millisecond):
				fmt.Println("  worker: doing work...")
			}
		}
	}(ctx)

	// Let the worker run briefly, then cancel.
	time.Sleep(350 * time.Millisecond)
	cancel() // Signals all goroutines watching this context.
	time.Sleep(50 * time.Millisecond) // Give goroutine time to print.
	fmt.Println()
}

// withTimeoutDemo shows automatic cancellation after a deadline.
// The cancel func should still be called (typically via defer) to free
// resources early if work finishes before the timeout.
func withTimeoutDemo() {
	fmt.Println("=== WithTimeout ===")

	// This context will auto-cancel after 200ms.
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // Always defer cancel — even with a timeout.

	// Simulate a slow operation that takes longer than the timeout.
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  operation completed (won't happen)")
	case <-ctx.Done():
		// context.DeadlineExceeded is returned when the timeout fires.
		fmt.Println("  operation timed out:", ctx.Err())
	}
	fmt.Println()
}

// withValueDemo shows passing request-scoped data through the context.
// This is common in HTTP middleware for things like request IDs and auth info.
func withValueDemo() {
	fmt.Println("=== WithValue ===")

	// Use a custom type for context keys to avoid collisions with other
	// packages that might use the same string key. This is a best practice.
	type contextKey string
	const requestIDKey contextKey = "requestID"

	ctx := context.WithValue(context.Background(), requestIDKey, "req-abc-123")

	// A downstream function retrieves the value.
	processRequest(ctx, requestIDKey)
	fmt.Println()
}

// processRequest reads a request-scoped value from the context.
func processRequest(ctx context.Context, key any) {
	// Value returns nil if the key is not found — always check.
	if reqID, ok := ctx.Value(key).(string); ok {
		fmt.Println("  processing request:", reqID)
	} else {
		fmt.Println("  no request ID in context")
	}
}

// cancellationPropagationDemo shows that cancelling a parent context
// automatically cancels all derived child contexts. This is how a
// single cancellation (e.g., HTTP request abort) cleans up an entire
// tree of goroutines.
func cancellationPropagationDemo() {
	fmt.Println("=== Cancellation Propagation (Parent -> Child -> Grandchild) ===")

	parent, cancelParent := context.WithCancel(context.Background())

	// Child and grandchild are derived from parent.
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild() // Good practice, though parent cancel covers it.

	grandchild, cancelGrandchild := context.WithCancel(child)
	defer cancelGrandchild()

	// Launch goroutines at each level.
	for _, label := range []struct {
		name string
		ctx  context.Context
	}{
		{"parent", parent},
		{"child", child},
		{"grandchild", grandchild},
	} {
		go func(name string, ctx context.Context) {
			<-ctx.Done()
			fmt.Printf("  %s cancelled: %v\n", name, ctx.Err())
		}(label.name, label.ctx)
	}

	// Cancelling the parent cascades to child and grandchild.
	fmt.Println("  cancelling parent...")
	cancelParent()

	time.Sleep(50 * time.Millisecond) // Let goroutines print.
	fmt.Println()
}
