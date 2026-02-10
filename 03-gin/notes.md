# Gin

## Key Concepts

- How Gin wraps `net/http`:
  - `gin.Engine` implements `http.Handler` via its `ServeHTTP` method — it plugs into the standard `http.Server`.
  - Each request gets a `gin.Context` pulled from a `sync.Pool`, reset, used, then returned. This avoids allocations per request.
  - Gin's `ResponseWriter` wraps `http.ResponseWriter` with extras like status tracking and body-size recording.
  - Because `Engine` is just an `http.Handler`, you can use it as a sub-handler, mount it behind a reverse proxy, or combine it with stdlib middleware.

- Router — radix-tree based:
  - Gin stores one radix (compressed trie) tree per HTTP method (`GET`, `POST`, etc.).
  - Lookup is O(length of path), not O(number of routes). This matters when you have hundreds of endpoints.
  - Path parameters (`:id`) and wildcards (`*filepath`) are parsed during tree traversal with zero allocation.
  - Route conflicts (e.g., `/users/:id` vs `/users/new`) are caught at registration time, not at request time.
  - Static segments are preferred over parameter segments when both could match, so `/users/new` wins over `/users/:id` for the path `/users/new`.

- Route groups and versioning:
  - `router.Group("/api/v1")` returns a `RouterGroup` that shares a common prefix and middleware set.
  - Groups nest: `v1.Group("/users")` produces `/api/v1/users`.
  - Middleware attached to a group applies to all routes within it and any nested sub-groups.
  - Common pattern for API versioning:
    ```go
    v1 := router.Group("/api/v1")
    v1.Use(authMiddleware())
    {
        v1.GET("/users", listUsers)
        v1.POST("/users", createUser)
    }

    v2 := router.Group("/api/v2")
    v2.Use(authMiddleware(), rateLimiter())
    {
        v2.GET("/users", listUsersV2)
    }
    ```
  - Use groups to scope middleware: public routes without auth, admin routes with auth + RBAC, internal routes with service-to-service token checks.

- Parameter binding — query, path, JSON body:
  - `c.Param("id")` reads path parameters (`:id`).
  - `c.Query("page")` and `c.DefaultQuery("page", "1")` read query strings.
  - `c.ShouldBindJSON(&req)` decodes the JSON body into a struct and returns an error on failure.
  - `c.ShouldBindQuery(&filter)` binds query params to a struct using `form` tags.
  - `c.ShouldBindUri(&params)` binds path params using `uri` tags.
  - **`ShouldBind` vs `Bind`**: `ShouldBind*` returns the error for you to handle. `Bind*` (a.k.a. `MustBind`) automatically responds with 400 and aborts — this couples binding to response writing, so prefer `ShouldBind*` for explicit control.
  - `c.ShouldBindBodyWith(&req, binding.JSON)` caches the body bytes so you can bind the same body more than once (normally the body is consumed on first read).

- Validation with struct tags:
  - Gin uses `go-playground/validator/v10` under the hood.
  - Tag syntax: `binding:"required,min=1,max=100"`.
  - Common validators: `required`, `email`, `url`, `min`, `max`, `len`, `oneof=active inactive`, `gt=0`.
  - Example:
    ```go
    type CreateUserReq struct {
        Name  string `json:"name"  binding:"required,min=2,max=50"`
        Email string `json:"email" binding:"required,email"`
        Role  string `json:"role"  binding:"required,oneof=admin member"`
    }
    ```
  - Custom validators can be registered on the validator engine:
    ```go
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("nowhitespace", noWhitespaceValidator)
    }
    ```
  - Validation errors are returned as `validator.ValidationErrors`, which you can iterate to build field-level error responses.

- Middleware chain (`HandlersChain`):
  - A route's handlers = group middleware + route-specific middleware + final handler, stored as a single `[]HandlerFunc` slice.
  - Execution flows left-to-right. Each middleware calls `c.Next()` to pass to the next handler, then continues after `c.Next()` returns (onion model).
  - `c.Abort()` stops further handlers from executing (sets the index past the end of the chain). Already-running middleware still finishes its post-`Next()` code.
  - Built-in middleware: `gin.Logger()` (request logging), `gin.Recovery()` (panic recovery with 500 response).
  - `gin.Default()` = `gin.New()` + Logger + Recovery. For production, use `gin.New()` and add your own.
  - Middleware ordering matters:
    1. Recovery (outermost — catches panics from everything below)
    2. Request-ID / correlation-ID injection
    3. Logging / metrics
    4. Auth / authz
    5. Route handler (innermost)
  - Middleware communicates via `c.Set(key, value)` / `c.Get(key)` — this is the request-scoped key-value store on the context.

- Custom middleware patterns:
  - **Logging**: record method, path, status, latency, request ID.
    ```go
    func RequestLogger() gin.HandlerFunc {
        return func(c *gin.Context) {
            start := time.Now()
            c.Next()
            slog.Info("request",
                "method", c.Request.Method,
                "path", c.Request.URL.Path,
                "status", c.Writer.Status(),
                "latency", time.Since(start),
            )
        }
    }
    ```
  - **Auth**: extract token, validate, set user identity in context, abort if invalid.
  - **Recovery**: catch panics, log stack trace, return 500 without crashing the process.
  - **Timeout**: wrap the handler in `context.WithTimeout` so downstream calls respect a deadline.

- Error handling patterns:
  - `c.AbortWithStatusJSON(code, obj)` — stops the chain and writes a JSON error response in one call. Use for guard-style checks in middleware.
  - `c.Error(err)` — appends an error to `c.Errors` without stopping the chain. Useful for collecting errors that a later error-handling middleware formats into a response.
  - Error middleware pattern: a final middleware that checks `c.Errors`, maps domain errors to status codes, and writes a consistent JSON error envelope.
    ```go
    func ErrorHandler() gin.HandlerFunc {
        return func(c *gin.Context) {
            c.Next()
            if len(c.Errors) > 0 {
                err := c.Errors.Last().Err
                // map domain errors to HTTP codes
                switch {
                case errors.Is(err, ErrNotFound):
                    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
                default:
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
                }
            }
        }
    }
    ```
  - Keep handlers thin: return domain errors via `c.Error(err)`, let middleware translate to HTTP.

- Gin vs stdlib `net/http` — when Gin helps, when it doesn't:
  - **Use Gin when**: you need path parameters, route groups, structured binding/validation, and middleware chaining without writing boilerplate. Most CRUD APIs and microservices benefit.
  - **Stick with stdlib when**: you have very few routes, need minimal dependencies, are building a proxy/gateway with custom connection handling, or the project policy favors zero third-party deps.
  - Gin adds ~5 MB to binary size. The performance difference vs stdlib is negligible for most workloads — the win is developer ergonomics, not raw speed.
  - Gin's `Context` is **not** `context.Context`. Pass `c.Request.Context()` to service/repo layers, never `*gin.Context` itself. This keeps your business logic framework-agnostic.

- Production considerations:
  - Set `gin.SetMode(gin.ReleaseMode)` to disable debug logging in production.
  - Use `gin.New()` instead of `gin.Default()` and register your own structured-logging and recovery middleware.
  - Graceful shutdown: Gin's engine is an `http.Handler`, so use `http.Server.Shutdown(ctx)` the same way as with stdlib.
  - Do not store mutable per-request state on the `Engine` or `RouterGroup` — those are shared across goroutines.
  - `gin.Context` must not be used after the request handler returns. If you need it in a goroutine, copy it with `c.Copy()`.

## Interview Questions

1. How does Gin's router differ from the stdlib `http.ServeMux`, and what advantage does a radix tree give?
2. What is the difference between `c.ShouldBindJSON` and `c.BindJSON`? Which do you prefer and why?
3. How does middleware execution work in Gin? What happens when you call `c.Abort()` in the middle of a chain?
4. Why should you pass `c.Request.Context()` to service-layer functions instead of passing `*gin.Context`?
5. How would you implement a consistent error response format across all endpoints using Gin middleware?
6. What is the purpose of `c.Set` / `c.Get`, and when would you use them vs. adding fields to a request-scoped struct?
7. How do you handle validation errors from `ShouldBind` to return field-level feedback to the client?
8. How would you version an API using Gin route groups, and how does middleware scoping play into that?
9. When would you choose stdlib `net/http` over Gin for a new Go service?
10. How do you implement graceful shutdown when using Gin?
11. Why should you avoid using `*gin.Context` after the handler returns, and what does `c.Copy()` do?
12. How would you add a request timeout middleware in Gin that cancels downstream work if the deadline is exceeded?

Practice prompt:
- Build a CRUD API with Gin using route groups (`/api/v1`), struct binding with validation, an error-handling middleware that maps domain errors to HTTP codes, and custom middleware for request logging and auth.

## Resources

- Gin GitHub repository and docs: https://github.com/gin-gonic/gin
- Gin API reference: https://pkg.go.dev/github.com/gin-gonic/gin
- go-playground/validator docs: https://pkg.go.dev/github.com/go-playground/validator/v10
- Gin examples directory: https://github.com/gin-gonic/gin/tree/master/examples
- Go `net/http` package docs (Gin builds on this): https://pkg.go.dev/net/http
