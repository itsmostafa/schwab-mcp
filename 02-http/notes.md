# HTTP in Go

## Key Concepts

- HTTP basics that matter in backend interviews:
  - Methods and semantics: `GET` (safe, idempotent), `PUT`/`DELETE` (idempotent), `POST` (not idempotent by default).
  - Status code classes: `2xx` success, `4xx` client errors, `5xx` server errors.
  - Headers and body responsibilities: `Content-Type`, `Accept`, `Authorization`, caching headers, correlation/request IDs.
  - Keep API behavior consistent: predictable error format, stable field naming, versioning strategy.

- `http.Handler` and `http.HandlerFunc`:
  - Core server abstraction in `net/http`.
  - Signature: `ServeHTTP(http.ResponseWriter, *http.Request)`.
  - `http.HandlerFunc` lets ordinary functions satisfy `http.Handler`.

- Routing with `http.ServeMux`:
  - Default router in stdlib; great for lightweight services.
  - Register patterns with `Handle` and `HandleFunc`.
  - Keep route setup centralized so middleware and auth are clear.

- Request lifecycle on the server:
  - Parse and validate request data early (path/query/body).
  - Apply auth/authorization before business logic.
  - Call service layer, map domain errors to HTTP codes, return structured response.
  - Always write status + body exactly once, and handle early returns carefully.

- Context propagation and cancellation:
  - `r.Context()` is canceled if client disconnects or server times out.
  - Pass context into DB/external calls so work stops when request is canceled.
  - Use `context.WithTimeout` for bounded work per request or downstream call.

- Middleware chain pattern:
  - Wrap handlers for reusable cross-cutting concerns (logging, metrics, auth, panic recovery).
  - Order matters: recovery should be outermost, auth before handler, logging around all.
  - Prefer explicit composition over hidden global behavior.

- JSON handling with `encoding/json`:
  - Decode request body into typed structs.
  - Validate required fields and reject unknown fields when strictness is needed.
  - Encode responses with explicit status codes and consistent envelope/error shape.
  - Close request bodies when manually reading; avoid unbounded reads.

- Input safety and limits:
  - Limit body size (`http.MaxBytesReader`) to prevent memory abuse.
  - Validate query/path/body values before use.
  - Set read/write/idle server timeouts to reduce slowloris-style risk.

- `http.Client` fundamentals (critical in real systems):
  - Do not use a new client per request; reuse clients/transports for connection pooling.
  - Set client timeout and, when needed, granular transport timeouts.
  - Reuse TCP connections with keep-alives for lower latency and fewer resources.
  - Always close response bodies (`defer resp.Body.Close()`).

- Retries and backoff:
  - Retry transient failures (`429`, some `5xx`, network timeouts) with exponential backoff + jitter.
  - Retry only idempotent operations unless protected by idempotency keys.
  - Add max attempts and total deadline to avoid retry storms.

- Production-ready HTTP concerns:
  - Observability: structured logs, request IDs, metrics (latency/error rate), traces.
  - Security: TLS, authn/authz, input validation, secure headers, avoid leaking internal errors.
  - Graceful shutdown: stop accepting new requests, finish in-flight work with deadline.

## Interview Questions

1. What is the difference between `http.Handler` and `http.HandlerFunc`?

   `http.Handler` is an interface with a single method: `ServeHTTP(http.ResponseWriter, *http.Request)`. Any type that implements this method satisfies the interface. `http.HandlerFunc` is a function type (`type HandlerFunc func(ResponseWriter, *Request)`) that has a `ServeHTTP` method defined on it, so it automatically satisfies `http.Handler`. This is the adapter pattern — it lets you pass a plain function anywhere an `http.Handler` is expected without defining a named struct. Use `HandlerFunc` for simple stateless handlers; use a struct implementing `Handler` when the handler needs dependencies (DB, logger, config) injected as fields.

2. Why should you reuse `http.Client` instead of creating one per request?

   `http.Client` (and its underlying `http.Transport`) manages a pool of persistent TCP connections with keep-alive. Creating a new client per request means each request opens a fresh TCP connection (with a new TLS handshake for HTTPS), which adds latency and consumes file descriptors and ports. Reusing a single client lets connections be recycled across requests, reducing connection overhead significantly. The default `http.Transport` has fields like `MaxIdleConns`, `MaxIdleConnsPerHost`, and `IdleConnTimeout` that control pooling behavior. A shared client also gives you a single place to configure timeouts, redirect policy, and custom transports.

3. Which HTTP methods are idempotent, and why does that matter for retries?

   `GET`, `PUT`, `DELETE`, `HEAD`, and `OPTIONS` are idempotent — repeating them produces the same server-side effect as a single call. `POST` and `PATCH` are not idempotent by default. This matters for retries because you can safely retry an idempotent request on transient failure (network timeout, `503`) without worrying about duplicating side effects. Retrying a `POST` (e.g., "create order") could create duplicate resources unless the API uses an idempotency key. When designing retry logic, only auto-retry idempotent methods (or non-idempotent ones protected by an idempotency key).

4. How do you structure middleware in `net/http`, and why does middleware order matter?

   Middleware in `net/http` follows a function-wrapping pattern: a middleware is a function that takes an `http.Handler` and returns a new `http.Handler` that adds behavior before/after calling the inner handler.

   ```go
   func Logging(next http.Handler) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           log.Printf("%s %s", r.Method, r.URL.Path)
           next.ServeHTTP(w, r)
       })
   }
   ```

   You compose middleware by nesting: `Recovery(Logging(Auth(handler)))`. The outermost middleware runs first on the way in and last on the way out. Order matters because:
   - **Panic recovery** should be outermost so it catches panics from any inner middleware or handler.
   - **Logging/metrics** should wrap auth so you log both allowed and rejected requests.
   - **Auth** should run before the business handler so unauthenticated requests are rejected early.
   - **Request ID** injection should be early so all downstream middleware and handlers can reference it.

5. How does request cancellation propagate in Go HTTP servers?

   Each incoming request carries a `context.Context` accessible via `r.Context()`. The server cancels this context when the client disconnects (TCP close / reset) or when a server-side timeout (e.g., `http.TimeoutHandler` or `WriteTimeout`) fires. If you pass `r.Context()` into downstream calls — database queries, HTTP client requests, channel operations — those calls will observe the cancellation via `ctx.Done()` and return early with `context.Canceled` or `context.DeadlineExceeded`. This prevents wasted work when the caller is no longer waiting for a response.

6. How do you prevent large request bodies from exhausting memory?

   Use `http.MaxBytesReader` to wrap `r.Body` before decoding:

   ```go
   r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB limit
   ```

   If the body exceeds the limit, reads return an error and the server automatically responds with `413 Request Entity Too Large`. Additionally, set `http.Server.ReadTimeout` and `ReadHeaderTimeout` to bound how long the server waits for the full request, preventing slowloris-style attacks that slowly trickle data. For streaming endpoints, use `io.LimitReader` or read in fixed-size chunks rather than buffering the entire body.

7. When should an API return `400` vs `422` vs `500`?

   - **400 Bad Request**: The request is structurally malformed — invalid JSON syntax, wrong Content-Type, missing required headers, unparseable query parameters. The server can't even parse the input.
   - **422 Unprocessable Entity**: The request is syntactically valid but semantically wrong — the JSON parses fine but a required field is empty, a value fails validation (email format, negative quantity), or business rules reject the data.
   - **500 Internal Server Error**: Something unexpected broke on the server — a nil pointer, database connection failure, unhandled error. The client's request was potentially fine; the server failed to process it.

   The key distinction: `400` = can't parse, `422` = parsed but invalid, `500` = server's fault. Some teams simplify by using `400` for both parse and validation errors, which is also acceptable and common.

8. How do you map domain/business errors to HTTP status codes cleanly?

   Define typed domain errors (sentinel errors or custom error types) in the service layer without any HTTP awareness:

   ```go
   var ErrNotFound = errors.New("not found")
   var ErrConflict = errors.New("conflict")
   type ValidationError struct { Field, Message string }
   ```

   Then in the HTTP handler (or a shared error-rendering helper), use `errors.Is` / `errors.As` to map them:

   ```go
   func writeError(w http.ResponseWriter, err error) {
       switch {
       case errors.Is(err, ErrNotFound):
           http.Error(w, "not found", http.StatusNotFound)
       case errors.Is(err, ErrConflict):
           http.Error(w, "conflict", http.StatusConflict)
       default:
           http.Error(w, "internal error", http.StatusInternalServerError)
       }
   }
   ```

   This keeps the service layer transport-agnostic. The HTTP layer is the only place that knows about status codes. If you later add gRPC, you write a separate mapper from domain errors to gRPC status codes.

9. What timeouts should you consider on `http.Server` and `http.Client`?

   **`http.Server` (inbound):**
   - `ReadTimeout` — max time to read the full request (headers + body). Prevents slow clients from holding connections.
   - `ReadHeaderTimeout` — max time to read just the headers. Defends against slowloris.
   - `WriteTimeout` — max time from end of request read to end of response write. Prevents handlers that hang forever.
   - `IdleTimeout` — how long keep-alive connections sit idle before closing.
   - `http.TimeoutHandler` — wraps a handler with a per-request deadline, returning `503` if exceeded.

   **`http.Client` (outbound):**
   - `Client.Timeout` — end-to-end timeout covering DNS, connect, TLS, request write, response read.
   - `Transport.DialContext` timeout or `net.Dialer.Timeout` — TCP connection timeout.
   - `Transport.TLSHandshakeTimeout` — TLS negotiation timeout.
   - `Transport.ResponseHeaderTimeout` — time waiting for response headers after writing the request.
   - Per-request `context.WithTimeout` — scoped deadline for individual calls.

   A good default: set `Client.Timeout` as a catch-all, use `context.WithTimeout` per call for finer control, and always set server-side `ReadTimeout` + `WriteTimeout`.

10. How do you design a consistent error response format for APIs?

    Define a standard JSON error envelope used across all endpoints:

    ```json
    {
      "error": {
        "code": "VALIDATION_FAILED",
        "message": "email is required",
        "details": [
          {"field": "email", "reason": "must not be empty"}
        ],
        "request_id": "abc-123"
      }
    }
    ```

    Key principles:
    - Use a machine-readable `code` (constant string or enum) that clients can switch on, separate from the human-readable `message`.
    - Include a `request_id` to correlate with server logs.
    - Use `details` for field-level validation errors so the client can highlight specific form fields.
    - Never leak stack traces, SQL errors, or internal paths in production error responses.
    - Write a shared `writeErrorResponse` helper that all handlers call, so the format is consistent everywhere.
    - Document the error codes in your API docs.

11. What risks come from retrying non-idempotent requests?

    The main risk is **duplicate side effects**. If a `POST /orders` request times out but the server actually processed it, retrying creates a second order. Similarly, retrying a `PATCH` that increments a counter would increment it twice. Specific risks include:
    - Duplicate resource creation (double charges, duplicate records).
    - Data corruption from applying a mutation twice.
    - Inconsistent state across systems if some downstream calls succeeded on the first attempt.

    Mitigations:
    - **Idempotency keys**: The client sends a unique key (e.g., UUID) in a header; the server deduplicates by storing the key and returning the cached response on retry.
    - **Conditional requests**: Use `If-Match` with ETags so the server rejects stale updates.
    - **Design for idempotency**: Make `POST` endpoints check for existing resources before creating (e.g., upsert semantics).

12. How can you handle partial failures when one endpoint calls multiple downstream services?

    Strategies depend on consistency requirements:
    - **Best-effort with error reporting**: Call all services, collect successes and failures, return a response indicating which parts succeeded and which failed. Good for non-critical aggregation endpoints.
    - **Fail fast**: If any downstream call fails, abort remaining calls and return an error. Simple but can leave partial state.
    - **Saga / compensating transactions**: If service A succeeds but service B fails, call a compensating action on service A to undo. Complex but preserves consistency.
    - **Outbox pattern**: Write the intent to a local database, then process downstream calls asynchronously with retries. Guarantees eventual consistency.

    Practical tips:
    - Use `errgroup.Group` to fan out concurrent calls with shared context cancellation — if one fails, cancel the rest.
    - Set per-call timeouts via `context.WithTimeout` so a slow downstream doesn't stall the entire request.
    - Return structured partial-failure responses so clients know what happened and can take action.

13. What is the purpose of context deadlines in external HTTP calls?

    Context deadlines bound how long you're willing to wait for an external service to respond. Without a deadline, a hanging downstream service blocks your goroutine (and the upstream request) indefinitely, eventually exhausting server resources. With `context.WithTimeout`, the HTTP client cancels the request when the deadline expires, the goroutine unblocks, and you can return a timely error to the caller.

    Deadlines also propagate: if the incoming request has 5 seconds remaining on its context and you set a 3-second timeout for a downstream call, the effective deadline is `min(5s, 3s) = 3s`. This ensures downstream calls never outlive the parent request. Use `context.WithTimeout` for absolute time budgets and `context.WithDeadline` when coordinating with an upstream deadline.

14. How do you implement graceful shutdown for an HTTP service in Go?

    ```go
    srv := &http.Server{Addr: ":8080", Handler: mux}

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("listen: %v", err)
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    // Give in-flight requests time to finish
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("shutdown: %v", err)
    }
    ```

    `srv.Shutdown` stops accepting new connections, waits for in-flight requests to complete (up to the context deadline), then returns. Key considerations:
    - Register `SIGINT` and `SIGTERM` so both Ctrl+C and container orchestrator signals trigger shutdown.
    - Choose a shutdown timeout that's shorter than your orchestrator's grace period (e.g., Kubernetes default is 30s, so use 25s).
    - Close other resources (DB pools, message consumers) after the HTTP server shuts down so in-flight handlers can still use them.

15. What metrics would you track first for a new HTTP API?

    The **RED method** (Rate, Errors, Duration) is the standard starting point:
    - **Request rate**: Requests per second, broken down by endpoint and method. Shows traffic patterns and helps capacity planning.
    - **Error rate**: Percentage of responses with `4xx` and `5xx` status codes, separated by code class. `5xx` rate is your primary reliability signal.
    - **Duration (latency)**: Response time histograms (p50, p95, p99) per endpoint. Percentiles reveal tail latency that averages hide.

    Additional high-value metrics:
    - **In-flight requests**: Current concurrency — detects saturation.
    - **Response size**: Bytes per response — detects unexpectedly large payloads.
    - **Downstream latency**: Time spent in external calls (DB, APIs) — identifies which dependency is slow.
    - **Connection pool stats**: Active/idle connections for DB and HTTP clients — detects pool exhaustion.

    Instrument these using Prometheus counters/histograms in a middleware so every endpoint is covered automatically.

Practice prompt:
- Build a small CRUD API in stdlib `net/http` with middleware for logging, panic recovery, auth stub, request ID, timeout, and consistent JSON errors.

## Resources

- Go `net/http` package docs: https://pkg.go.dev/net/http
- Go `context` package docs: https://pkg.go.dev/context
- Go `encoding/json` package docs: https://pkg.go.dev/encoding/json
- HTTP Semantics (RFC 9110): https://www.rfc-editor.org/rfc/rfc9110
- REST API resource naming and style guides:
  - Microsoft REST API Guidelines: https://github.com/microsoft/api-guidelines
  - Google API Design Guide: https://cloud.google.com/apis/design
- Resilience patterns (timeouts, retries, backoff):
  - AWS Exponential Backoff and Jitter: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
