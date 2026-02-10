// client.go — Reusable http.Client, Transport tuning, retries with exponential backoff
//
// Run:
//   go run client.go
//
// This example starts a local test server and makes requests to it, so it is
// fully self-contained. Watch the output to see connection reuse, retries, and
// backoff in action.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

// --- Reusable client setup ---

// NewHTTPClient creates a production-ready http.Client with tuned Transport
// settings. The key interview point: NEVER use http.DefaultClient in production
// because it has no timeout (requests can hang forever) and the default transport
// settings may not match your workload.
func NewHTTPClient() *http.Client {
	// Transport controls the low-level connection pooling and TLS behavior.
	// These settings are the most commonly tuned knobs.
	transport := &http.Transport{
		// MaxIdleConns is the total number of idle (keep-alive) connections
		// across ALL hosts. The default is 100, which is fine for most services.
		// Increase if you talk to many distinct hosts concurrently.
		MaxIdleConns: 100,

		// MaxIdleConnsPerHost is the max idle connections kept per host.
		// The default is only 2, which is too low if you send many concurrent
		// requests to the same host — excess connections get closed and
		// reopened, wasting TCP/TLS handshakes. Set this to match your
		// expected concurrency per host.
		MaxIdleConnsPerHost: 10,

		// IdleConnTimeout is how long an idle connection sits in the pool
		// before being closed. 90 seconds is the default. Lower it if you
		// want to free resources faster; raise it if you have bursty traffic
		// with long pauses between bursts.
		IdleConnTimeout: 90 * time.Second,
	}

	return &http.Client{
		// Client.Timeout is an end-to-end deadline covering DNS lookup,
		// TCP connect, TLS handshake, request send, and response read.
		// This is your safety net — without it, a slow server can block
		// your goroutine indefinitely.
		Timeout: 30 * time.Second,

		Transport: transport,
	}
}

// --- Retry with exponential backoff ---

var (
	// ErrMaxRetries is returned when all retry attempts are exhausted.
	ErrMaxRetries = errors.New("max retries exceeded")
)

// RetryConfig holds parameters for the retry strategy.
type RetryConfig struct {
	MaxAttempts int           // total attempts including the first one
	BaseDelay   time.Duration // initial wait before first retry
	MaxDelay    time.Duration // cap on the backoff duration
}

// retryable checks if a status code should trigger a retry.
// Only retry on server errors and rate limiting — client errors (4xx) are
// the caller's fault and retrying won't help (except 429).
func retryable(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests,     // 429 — rate limited, try again later
		http.StatusBadGateway,           // 502 — upstream error, often transient
		http.StatusServiceUnavailable,   // 503 — server overloaded/deploying
		http.StatusGatewayTimeout:       // 504 — upstream timeout
		return true
	default:
		return false
	}
}

// DoWithRetry executes an HTTP request with exponential backoff and jitter.
//
// Interview notes:
// - Only retry idempotent methods (GET, PUT, DELETE, HEAD) or operations
//   protected by an idempotency key. Retrying a POST blindly can create
//   duplicate resources.
// - Jitter prevents the "thundering herd" problem where many clients retry
//   at the exact same time after a shared failure.
// - The context deadline acts as an overall timeout across all attempts.
func DoWithRetry(ctx context.Context, client *http.Client, req *http.Request, cfg RetryConfig) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if attempt > 0 {
			// Exponential backoff: delay doubles each attempt.
			// 1 * base, 2 * base, 4 * base, ...
			delay := cfg.BaseDelay * (1 << (attempt - 1))
			if delay > cfg.MaxDelay {
				delay = cfg.MaxDelay
			}

			// Add jitter: random value between 0 and delay.
			// This spreads retry storms across time.
			jitter := time.Duration(rand.Int63n(int64(delay)))
			wait := delay/2 + jitter/2 // center the jitter around the delay

			log.Printf("  retry %d/%d — waiting %s", attempt, cfg.MaxAttempts-1, wait)

			select {
			case <-time.After(wait):
				// Continue with retry.
			case <-ctx.Done():
				// Caller's context expired — stop retrying.
				return nil, ctx.Err()
			}
		}

		// Clone the request for each attempt. The original request's body
		// may have been consumed by a previous attempt. For GET requests
		// (nil body) this is fine as-is. For POST/PUT, the caller should
		// provide a GetBody function or pass a body that can be re-read.
		reqCopy := req.Clone(ctx)
		resp, err = client.Do(reqCopy)
		if err != nil {
			// Network error (DNS failure, connection refused, timeout).
			// These are often transient, so we retry.
			log.Printf("  attempt %d failed: %v", attempt+1, err)
			continue
		}

		// Got a response — check if we should retry.
		if !retryable(resp.StatusCode) {
			return resp, nil // Success or non-retryable error.
		}

		// Response is retryable — close this body before retrying.
		// Failing to close the body leaks the TCP connection.
		log.Printf("  attempt %d got status %d (retryable)", attempt+1, resp.StatusCode)
		resp.Body.Close()
	}

	if err != nil {
		return nil, fmt.Errorf("%w: last error: %v", ErrMaxRetries, err)
	}
	return resp, fmt.Errorf("%w: last status %d", ErrMaxRetries, resp.StatusCode)
}

// --- Demo helpers ---

// readBody reads and closes the response body. This two-step pattern is
// critical: you must ALWAYS close resp.Body, even if you don't read it.
// Forgetting to close it leaks the underlying TCP connection back to the
// pool, eventually exhausting available connections.
func readBody(resp *http.Response) ([]byte, error) {
	// defer resp.Body.Close() ensures the body is closed even if ReadAll
	// returns an error. Always defer the close immediately after checking
	// that resp is not nil.
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func main() {
	// --- Start a test server that simulates flaky behavior ---
	callCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		switch r.URL.Path {
		case "/ok":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success",
				"method":  r.Method,
			})

		case "/flaky":
			// Fail the first 2 requests, succeed on the 3rd.
			// This simulates a temporarily overloaded service.
			if callCount <= 2 {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, "service unavailable (attempt %d)", callCount)
				return
			}
			callCount = 0 // reset for next demo
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "recovered"})

		case "/echo":
			// Echo back the POST body.
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{
				"echo": string(body),
			})

		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	// Create our reusable client.
	client := NewHTTPClient()

	ctx := context.Background()

	// --- Demo 1: Simple GET request ---
	fmt.Println("=== Demo 1: Simple GET ===")
	{
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL+"/ok", nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		// Always close the body. The defer pattern here guarantees cleanup
		// even if the code below panics or returns early.
		body, err := readBody(resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  status: %d\n", resp.StatusCode)
		fmt.Printf("  body:   %s\n", body)
	}

	// --- Demo 2: POST with JSON body ---
	fmt.Println("\n=== Demo 2: POST with JSON body ===")
	{
		payload := map[string]string{"name": "gopher", "role": "backend"}
		jsonBytes, _ := json.Marshal(payload)

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, ts.URL+"/echo", bytes.NewReader(jsonBytes))
		if err != nil {
			log.Fatal(err)
		}
		// Always set Content-Type when sending a body. The server may reject
		// requests without it or misparse the body.
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, err := readBody(resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  status: %d\n", resp.StatusCode)
		fmt.Printf("  body:   %s\n", body)
	}

	// --- Demo 3: Retry with backoff on flaky endpoint ---
	fmt.Println("\n=== Demo 3: Retry with exponential backoff ===")
	{
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL+"/flaky", nil)
		if err != nil {
			log.Fatal(err)
		}

		retryCfg := RetryConfig{
			MaxAttempts: 5,
			BaseDelay:   200 * time.Millisecond,
			MaxDelay:    2 * time.Second,
		}

		resp, err := DoWithRetry(ctx, client, req, retryCfg)
		if err != nil {
			log.Fatalf("all retries failed: %v", err)
		}
		body, err := readBody(resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  status: %d\n", resp.StatusCode)
		fmt.Printf("  body:   %s\n", body)
	}

	// --- Demo 4: Context timeout cancels retries ---
	fmt.Println("\n=== Demo 4: Context timeout cancels retries ===")
	{
		// Set a very short deadline so retries are cut off.
		shortCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
		defer cancel()

		req, err := http.NewRequestWithContext(shortCtx, http.MethodGet, ts.URL+"/flaky", nil)
		if err != nil {
			log.Fatal(err)
		}

		retryCfg := RetryConfig{
			MaxAttempts: 10,
			BaseDelay:   500 * time.Millisecond, // longer than context deadline
			MaxDelay:    5 * time.Second,
		}

		_, err = DoWithRetry(shortCtx, client, req, retryCfg)
		if err != nil {
			// Expected: context deadline exceeded stops retries.
			fmt.Printf("  retries stopped: %v\n", err)
		}
	}

	fmt.Println("\ndone")
}
