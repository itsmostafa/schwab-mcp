package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Fan-out/fan-in with WaitGroup and channels.
//
// GMP model (interview essentials):
//   G = goroutine  — a lightweight user-space thread (~2-4 KB initial stack, grows as needed).
//   M = OS thread  — the actual kernel thread that executes code.
//   P = processor  — a logical processor that holds a run queue of Gs.
//   The runtime multiplexes many Gs onto a smaller set of Ms via Ps.
//
// Goroutine cost:
//   Each goroutine starts with a ~2-4 KB stack (vs ~1 MB for an OS thread).
//   This makes spawning thousands of goroutines practical and cheap.
//
// Asynchronous preemption (Go 1.14+):
//   Before 1.14, a goroutine had to hit a function call to be preempted (cooperative).
//   Since 1.14, the runtime can preempt goroutines at (almost) any safe point via signals,
//   preventing a tight loop from starving other goroutines.

func main() {
	fmt.Println("=== Fan-Out / Fan-In with Worker Pool ===\n")

	const numWorkers = 3
	const numJobs = 9

	// jobs channel: workers read from this (fan-out).
	// buffered so the producer doesn't block on every send.
	jobs := make(chan int, numJobs)

	// results channel: workers write to this (fan-in).
	results := make(chan string, numJobs)

	// WaitGroup tracks when all workers are done so we can safely close
	// the results channel. Without this, we'd risk reading from a channel
	// that's still being written to, or closing it prematurely (panic).
	var wg sync.WaitGroup

	// --- Fan-out: launch N workers ---
	// Each worker is a goroutine that pulls jobs from the shared channel.
	// Because channels are concurrency-safe, no extra locking is needed.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// --- Send jobs ---
	// The producer feeds work into the jobs channel.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Signal workers that no more jobs are coming.

	// --- Wait and close results ---
	// We launch this in a goroutine so the main goroutine can start reading
	// results immediately (below). If we called wg.Wait() on main, we'd
	// deadlock because the results channel would fill up with no reader.
	go func() {
		wg.Wait()
		close(results) // Safe to close: all writers are done.
	}()

	// --- Fan-in: collect results ---
	// range over a channel reads until it is closed.
	fmt.Println("Results:")
	for r := range results {
		fmt.Println("  ", r)
	}

	fmt.Println("\nAll jobs processed.")
}

// worker simulates a unit of work. Each worker pulls jobs until the jobs
// channel is closed, processes them, and sends results to the results channel.
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	// defer ensures Done is called even if the function panics.
	defer wg.Done()

	for job := range jobs {
		// Simulate variable processing time to show interleaved execution.
		duration := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(duration)

		result := fmt.Sprintf("worker %d processed job %d (took %v)", id, job, duration)
		results <- result
	}
}
