// middleware.go — Middleware chain: RequestID, Logging, Recovery, Auth, Timeout
//
// Run:
//   go run middleware.go
//
// Test with curl:
//   curl -v localhost:8080/                         # 401 — no auth header
//   curl -v -H "Authorization: Bearer secret" localhost:8080/        # 200
//   curl -v -H "Authorization: Bearer secret" localhost:8080/panic   # 500 — recovered
//   curl -v -H "Authorization: Bearer secret" localhost:8080/slow    # 503 — timeout
//
// Observe:
//   - X-Request-Id header in every response
//   - Log lines showing method, path, status, and duration
//   - Middleware execution order in log output

package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// --- Middleware pattern ---
//
// The canonical Go middleware signature is:
//   func(http.Handler) http.Handler
//
// This creates an "onion model": each middleware wraps the next, forming
// concentric layers. A request passes inward through each layer's pre-logic,
// hits the core handler, then passes outward through each layer's post-logic.
//
// Composition: Recovery(Logging(Auth(handler)))
// Execution order on request:  Recovery.pre → Logging.pre → Auth.pre → handler
// Execution order on response: handler → Auth.post → Logging.post → Recovery.post

// --- RequestID middleware ---

// requestIDKey is an unexported type to avoid context key collisions.
// Using a named type instead of a bare string prevents any other package
// from accidentally overwriting this value in the context.
type requestIDKey struct{}

// RequestID generates a unique ID for each request, adds it to the context
// (so downstream handlers and services can include it in logs/traces), and
// sets it as a response header (so clients can reference it in bug reports).
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := generateID()

		// Propagate through context so any code with access to r.Context()
		// can retrieve the request ID — this is how you correlate logs across
		// layers without passing the ID as a function parameter.
		ctx := context.WithValue(r.Context(), requestIDKey{}, id)

		// Set the header before calling next, because once next writes the
		// status code, headers are frozen (you can't add more).
		w.Header().Set("X-Request-Id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// generateID creates a short random hex string. In production you'd use
// a UUID library or pull a trace ID from OpenTelemetry.
func generateID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// GetRequestID extracts the request ID from context. Exported so handlers
// and service-layer code can use it for structured logging.
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey{}).(string); ok {
		return id
	}
	return "unknown"
}

// --- Logging middleware ---

// statusRecorder wraps http.ResponseWriter to capture the status code.
// ResponseWriter doesn't expose the status after WriteHeader is called,
// so we need this wrapper to log it after the handler completes.
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

// Logging records the method, path, status code, and duration of each request.
// It runs both before and after the inner handler — the "pre" part captures
// the start time, the "post" part (after next.ServeHTTP) computes duration.
func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	// Returning a closure lets us inject dependencies (the logger) into
	// middleware while keeping the func(Handler)Handler signature.
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the writer so we can capture the status code.
			rec := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

			// Call the next handler in the chain.
			next.ServeHTTP(rec, r)

			// Post-handler: log the completed request.
			reqID := GetRequestID(r.Context())
			logger.Printf("[%s] %s %s → %d (%s)",
				reqID, r.Method, r.URL.Path, rec.statusCode, time.Since(start))
		})
	}
}

// --- Recovery middleware ---

// Recovery catches panics from downstream handlers and converts them into
// 500 responses instead of crashing the entire server process.
//
// Interview note: this MUST be the outermost middleware (first in the chain)
// so it wraps everything. If a panic occurs in Logging or Auth middleware,
// Recovery still catches it. Go's HTTP server also has a built-in recovery
// per connection goroutine, but it closes the connection — this middleware
// lets you return a proper JSON error and keep the connection alive.
func Recovery(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					reqID := GetRequestID(r.Context())
					logger.Printf("[%s] PANIC recovered: %v", reqID, err)

					// Only write the error response if headers haven't been
					// sent yet. If the handler already called WriteHeader,
					// writing again is a no-op (and logs a warning), so we
					// still attempt it as a best effort.
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, `{"error":"internal server error"}`)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// --- Auth middleware ---

// Auth is a simplified authentication middleware that checks for a specific
// Authorization header value. In production, you'd validate a JWT or session
// token, extract claims, and add the authenticated user to the context.
//
// Interview note: auth middleware should run AFTER logging (so rejected
// requests are still logged) but BEFORE the business handler (so
// unauthenticated requests never reach business logic).
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error":"missing authorization header"}`)
			return // Short-circuit: don't call next — request stops here.
		}

		// In a real system you'd validate the token, extract user info,
		// and store it in the context:
		//   ctx := context.WithValue(r.Context(), userKey{}, user)
		//   next.ServeHTTP(w, r.WithContext(ctx))

		next.ServeHTTP(w, r)
	})
}

// --- Timeout middleware ---

// Timeout wraps each request with a context deadline. If the handler doesn't
// finish within the duration, the context is canceled. Note: this cancels the
// context but does NOT automatically write a response — the handler must check
// ctx.Done() or ctx.Err() to detect the cancellation and return early.
//
// For automatic 503 responses on timeout, use http.TimeoutHandler from stdlib.
// We show both approaches here for comparison.
func Timeout(dt time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// http.TimeoutHandler is the stdlib's built-in timeout middleware.
		// It runs the handler in a separate goroutine, and if it exceeds dt,
		// it writes a 503 Service Unavailable response automatically.
		// The message parameter is the response body on timeout.
		return http.TimeoutHandler(next, dt, `{"error":"request timed out"}`)
	}
}

// --- Chain helper ---

// Chain applies middleware in the order given. The first middleware in the
// slice becomes the outermost layer (runs first on request, last on response).
// This avoids deeply nested calls like Recovery(Logging(Auth(handler))).
func Chain(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	// Apply in reverse so the first middleware in the list is outermost.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// --- Handlers ---

func helloHandler(w http.ResponseWriter, r *http.Request) {
	reqID := GetRequestID(r.Context())
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"hello","request_id":"%s"}`, reqID)
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	// This panic will be caught by the Recovery middleware.
	panic("something went terribly wrong")
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate slow work. The Timeout middleware will cancel this request
	// if it exceeds the deadline.
	select {
	case <-time.After(5 * time.Second):
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"finished slow work"}`)
	case <-r.Context().Done():
		// Context was canceled (by timeout or client disconnect).
		// In a real handler, you'd stop any in-progress work here
		// (cancel DB queries, stop HTTP calls, etc.).
		return
	}
}

// --- Main ---

func main() {
	logger := log.New(os.Stdout, "HTTP ", log.LstdFlags|log.Lmicroseconds)

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/panic", panicHandler)
	mux.HandleFunc("/slow", slowHandler)

	// Build the middleware chain. Order matters — read bottom to top to
	// understand the onion layers from outermost to innermost:
	//
	//   Request →  Recovery → RequestID → Logging → Timeout → Auth → Handler
	//   Response ← Recovery ← RequestID ← Logging ← Timeout ← Auth ← Handler
	//
	// - Recovery is outermost: catches panics from ALL other layers
	// - RequestID is early: every subsequent layer can read the ID from context
	// - Logging wraps Timeout+Auth+Handler: logs the final status and duration
	// - Timeout bounds the total handler execution time
	// - Auth is innermost middleware: rejects unauthenticated requests before handler
	handler := Chain(
		mux,
		Recovery(logger),  // outermost
		RequestID,         //
		Logging(logger),   //
		Timeout(2*time.Second),
		Auth,              // innermost middleware
	)

	fmt.Println("middleware demo listening on :8080")
	fmt.Println("try: curl -v localhost:8080/")
	fmt.Println("try: curl -v -H 'Authorization: Bearer secret' localhost:8080/")
	fmt.Println("try: curl -v -H 'Authorization: Bearer secret' localhost:8080/panic")
	fmt.Println("try: curl -v -H 'Authorization: Bearer secret' localhost:8080/slow")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
