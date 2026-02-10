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
2. Why should you reuse `http.Client` instead of creating one per request?
3. Which HTTP methods are idempotent, and why does that matter for retries?
4. How do you structure middleware in `net/http`, and why does middleware order matter?
5. How does request cancellation propagate in Go HTTP servers?
6. How do you prevent large request bodies from exhausting memory?
7. When should an API return `400` vs `422` vs `500`?
8. How do you map domain/business errors to HTTP status codes cleanly?
9. What timeouts should you consider on `http.Server` and `http.Client`?
10. How do you design a consistent error response format for APIs?
11. What risks come from retrying non-idempotent requests?
12. How can you handle partial failures when one endpoint calls multiple downstream services?
13. What is the purpose of context deadlines in external HTTP calls?
14. How do you implement graceful shutdown for an HTTP service in Go?
15. What metrics would you track first for a new HTTP API?

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
