// graceful_shutdown.go — srv.Shutdown with signal handling and in-flight request draining
//
// Run:
//   go run graceful_shutdown.go
//
// Test:
//   1. In one terminal: go run graceful_shutdown.go
//   2. In another terminal: curl localhost:8080/slow  (takes 3 seconds)
//   3. While the curl is running, press Ctrl+C in the server terminal
//   4. Observe: the slow request completes before the server exits
//   5. Try: curl localhost:8080/slow after Ctrl+C — connection refused (no new requests)

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// A fast handler for quick testing.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello at %s\n", time.Now().Format(time.RFC3339))
	})

	// A slow handler that simulates real work (DB query, external API call).
	// During shutdown, this handler should be allowed to finish if it started
	// before the shutdown signal arrived.
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		log.Println("slow handler: started (will take 3 seconds)")

		// Simulate work while respecting context cancellation.
		// In production, pass r.Context() to downstream calls (DB, HTTP client)
		// so they stop if the shutdown deadline expires.
		select {
		case <-time.After(3 * time.Second):
			log.Println("slow handler: work completed")
			fmt.Fprintf(w, "slow response completed at %s\n", time.Now().Format(time.RFC3339))
		case <-r.Context().Done():
			// This fires if the shutdown context deadline expires before
			// the handler finishes. The server gave up waiting.
			log.Println("slow handler: context canceled, aborting")
			http.Error(w, "request canceled", http.StatusServiceUnavailable)
		}
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// --- Start the server in a goroutine ---
	//
	// ListenAndServe blocks, so we run it in a goroutine to let main()
	// continue to the signal-handling logic below.
	go func() {
		log.Println("server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// ErrServerClosed is the expected error when Shutdown() is called.
			// Any other error means the server failed to start or crashed.
			log.Fatalf("listen error: %v", err)
		}
		// Interview note: after Shutdown() is called, ListenAndServe returns
		// http.ErrServerClosed. This is normal — it means the server stopped
		// gracefully, not that something went wrong.
		log.Println("server stopped accepting new connections")
	}()

	// --- Signal handling ---
	//
	// We listen for SIGINT (Ctrl+C) and SIGTERM (sent by Kubernetes, systemd,
	// Docker, etc. during pod/service shutdown).
	//
	// The channel is buffered with size 1. This is important: if the signal
	// arrives before we're ready to receive, it won't be lost. signal.Notify
	// does not block on send — it drops the signal if the channel is full.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a shutdown signal.
	sig := <-quit
	log.Printf("received signal: %v — starting graceful shutdown", sig)

	// --- Graceful shutdown ---
	//
	// srv.Shutdown does three things:
	// 1. Closes the listener — no new connections are accepted.
	// 2. Closes idle connections immediately.
	// 3. Waits for active (in-flight) requests to complete.
	//
	// The context deadline is a safety net: if in-flight requests take longer
	// than the deadline, Shutdown returns an error and the server stops anyway.
	//
	// Interview note: Shutdown vs Close:
	// - Shutdown: graceful — stops new connections, drains in-flight requests.
	// - Close: immediate — forcibly closes all connections, in-flight requests
	//   are interrupted. Use Close only as a last resort (e.g., if Shutdown
	//   times out and you must exit).
	//
	// In Kubernetes, the kubelet sends SIGTERM and waits for the grace period
	// (default 30s) before SIGKILL. Set your shutdown timeout shorter than
	// the grace period (e.g., 25s) to ensure a clean exit.

	shutdownTimeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	log.Printf("waiting up to %s for in-flight requests to complete...", shutdownTimeout)

	if err := srv.Shutdown(ctx); err != nil {
		// Shutdown returned an error — the context deadline expired before
		// all in-flight requests finished. In production, you might call
		// srv.Close() here to force-stop, then exit.
		log.Printf("shutdown error (some requests may have been interrupted): %v", err)

		// Force close as a fallback.
		if closeErr := srv.Close(); closeErr != nil {
			log.Printf("force close error: %v", closeErr)
		}
	} else {
		log.Println("all in-flight requests completed")
	}

	// --- Post-shutdown cleanup ---
	//
	// Close other resources AFTER the HTTP server shuts down, not before.
	// In-flight handlers may still need the DB pool, message queues, etc.
	// Order: stop accepting requests → drain requests → close dependencies.
	log.Println("cleaning up resources (DB pools, message queues, etc.)...")
	time.Sleep(100 * time.Millisecond) // simulate cleanup

	log.Println("server exited cleanly")

	// Summary of the shutdown sequence:
	// 1. Signal received (SIGINT or SIGTERM)
	// 2. srv.Shutdown called — listener closed, no new connections
	// 3. Wait for in-flight requests to finish (bounded by context timeout)
	// 4. Close application dependencies (DB, caches, queues)
	// 5. Process exits
}
