package main

import (
	"fmt"
	"time"
)

// Channels: buffered vs unbuffered, nil channel in select, close/range.
//
// Channel axioms (must-know for interviews):
//   1. Send to a nil channel    -> blocks forever.
//   2. Receive from nil channel -> blocks forever.
//   3. Send to a closed channel -> panic.
//   4. Receive from closed channel -> returns zero value immediately (ok == false).
//   5. Close a nil channel      -> panic.
//   6. Close an already-closed channel -> panic.
//
// Happens-before guarantees:
//   - A send on a channel happens-before the corresponding receive completes.
//   - The close of a channel happens-before a receive of the zero value.
//   - For unbuffered: the receive happens-before the send completes
//     (both goroutines rendezvous at the channel operation).

func main() {
	unbufferedDemo()
	bufferedDemo()
	nilChannelSelectDemo()
	closeRangeDemo()
}

// unbufferedDemo shows that an unbuffered channel is a rendezvous point:
// the sender blocks until a receiver is ready, and vice versa.
// This gives a strong synchronization guarantee.
func unbufferedDemo() {
	fmt.Println("=== Unbuffered Channel (Rendezvous) ===")

	ch := make(chan string) // capacity 0 — unbuffered

	go func() {
		// This send will block until main's receive is ready.
		ch <- "hello from goroutine"
	}()

	// This receive blocks until the goroutine sends.
	// The two sides "meet" — neither can proceed alone.
	msg := <-ch
	fmt.Println("received:", msg)
	fmt.Println()
}

// bufferedDemo shows that a buffered channel allows sends to proceed
// without a receiver, up to the buffer capacity.
func bufferedDemo() {
	fmt.Println("=== Buffered Channel ===")

	ch := make(chan int, 3) // capacity 3

	// We can send up to 3 values without blocking, even with no receiver yet.
	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40 would block here because the buffer is full.

	fmt.Printf("buffered channel: len=%d, cap=%d\n", len(ch), cap(ch))

	// Drain the buffer.
	fmt.Println("received:", <-ch)
	fmt.Println("received:", <-ch)
	fmt.Println("received:", <-ch)
	fmt.Println()
}

// nilChannelSelectDemo shows how setting a channel to nil in a select
// effectively disables that case. This is a common pattern for merging
// multiple channels and stopping reads from exhausted sources.
func nilChannelSelectDemo() {
	fmt.Println("=== Nil Channel in Select (Dynamic Case Disabling) ===")

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	ch1 <- "a1"
	ch1 <- "a2"
	close(ch1)

	ch2 <- "b1"
	close(ch2)

	// We want to drain both channels. When one is exhausted, we set it
	// to nil so its select case is never chosen again (receive on nil
	// blocks forever, so select skips it).
	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				fmt.Println("  ch1 closed, disabling")
				ch1 = nil // Disable this case in future iterations.
				continue
			}
			fmt.Println("  ch1:", v)
		case v, ok := <-ch2:
			if !ok {
				fmt.Println("  ch2 closed, disabling")
				ch2 = nil
				continue
			}
			fmt.Println("  ch2:", v)
		}
	}
	fmt.Println()
}

// closeRangeDemo shows the idiomatic pattern: a producer closes the channel
// when done, and consumers use range to read until closure.
func closeRangeDemo() {
	fmt.Println("=== Close + Range Pattern ===")

	ch := make(chan int)

	// Producer: sends values then closes the channel.
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * i
			time.Sleep(20 * time.Millisecond) // Simulate work.
		}
		close(ch) // Signals that no more values will be sent.
	}()

	// Consumer: range loops until ch is closed.
	// This is equivalent to: for { v, ok := <-ch; if !ok { break }; ... }
	for v := range ch {
		fmt.Println("  received:", v)
	}

	// After close, any receive returns the zero value immediately.
	val, ok := <-ch
	fmt.Printf("  after close: val=%d, ok=%v (zero value, channel drained)\n", val, ok)
	fmt.Println()
}
