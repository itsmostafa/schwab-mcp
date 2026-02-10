// Run: go mod init example && go get github.com/gin-gonic/gin && go run middleware.go
//
// This example demonstrates custom Gin middleware patterns: request logging,
// token-based auth with c.Set/c.Get, centralized error handling via c.Error(),
// and request timeouts using context.WithTimeout.
//
// Test with curl:
//   curl -s localhost:8080/public/health | jq                                              # no auth required
//   curl -s localhost:8080/api/profile | jq                                                # 401 — missing token
//   curl -s localhost:8080/api/profile -H 'Authorization: Bearer token-user-42' | jq       # 200 — auth passes
//   curl -s localhost:8080/api/slow -H 'Authorization: Bearer token-user-42' | jq          # may timeout (2s limit)
//   curl -s -X POST localhost:8080/api/fail -H 'Authorization: Bearer token-user-42' | jq  # 404 — domain error mapped by error handler

package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// --- Domain errors ---
// Define sentinel errors so the error handler middleware can map them to
// HTTP status codes. This decouples handlers from HTTP concerns.
var (
	ErrNotFound = errors.New("resource not found")
	ErrForbidden = errors.New("access forbidden")
)

func main() {
	// gin.New() gives a bare engine with no middleware. In production you
	// should use this and attach your own structured middleware rather than
	// relying on gin.Default()'s built-in Logger and Recovery.
	r := gin.New()

	// Middleware ordering matters. Recovery is outermost so it catches panics
	// from everything below. Logger wraps all routes. ErrorHandler runs its
	// post-Next() logic last, after handlers have called c.Error().
	r.Use(gin.Recovery())  // catches panics, returns 500
	r.Use(RequestLogger()) // logs method, path, status, latency
	r.Use(ErrorHandler())  // maps domain errors to HTTP status codes

	// Public routes — no auth middleware.
	public := r.Group("/public")
	{
		public.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	// Protected routes — auth middleware applied to the group.
	// Interview note: middleware attached to a group applies to all routes
	// in that group and any nested sub-groups.
	api := r.Group("/api")
	api.Use(AuthMiddleware())
	api.Use(TimeoutMiddleware(2 * time.Second))
	{
		api.GET("/profile", profileHandler)
		api.GET("/slow", slowHandler)
		api.POST("/fail", failHandler)
	}

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}

// ---------------------------------------------------------------------------
// Middleware 1: RequestLogger
// ---------------------------------------------------------------------------

// RequestLogger logs method, path, status code, and latency for every request.
// It calls c.Next() to execute downstream handlers, then logs the result —
// this is the "onion model" where pre-Next code runs on the way in and
// post-Next code runs on the way out.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// c.Next() passes control to the next handler in the chain.
		// Without calling c.Next(), downstream handlers would never execute.
		c.Next()

		// After c.Next() returns, all downstream handlers have finished.
		// c.Writer.Status() reflects the final status code.
		slog.Info("request completed",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", time.Since(start).String(),
			"client_ip", c.ClientIP(),
		)
	}
}

// ---------------------------------------------------------------------------
// Middleware 2: AuthMiddleware
// ---------------------------------------------------------------------------

// AuthMiddleware checks for a Bearer token in the Authorization header.
// On success it stores the extracted user ID with c.Set so downstream
// handlers can retrieve it with c.Get or c.MustGet.
//
// Interview note: c.Set/c.Get is a request-scoped key-value store on
// *gin.Context. It's the idiomatic way to pass data between middleware
// and handlers — e.g., authenticated user identity, request IDs, etc.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// c.AbortWithStatusJSON stops the chain AND writes the response.
			// Interview note: c.Abort() sets the handler index past the end of
			// the chain so no further handlers run. But the current middleware
			// continues executing its remaining code.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization header",
			})
			return
		}

		// Simple token parsing for demonstration. In production you would
		// validate a JWT or call an auth service here.
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid Authorization format, expected 'Bearer <token>'",
			})
			return
		}

		token := parts[1]

		// Extract a fake user ID from the token for demonstration.
		userID := strings.TrimPrefix(token, "token-user-")
		if userID == token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		// Store the user ID so handlers can access it.
		c.Set("userID", userID)

		c.Next()
	}
}

// ---------------------------------------------------------------------------
// Middleware 3: ErrorHandler
// ---------------------------------------------------------------------------

// ErrorHandler collects errors that handlers attached via c.Error(err) and
// maps domain errors to appropriate HTTP status codes.
//
// Interview note: this pattern keeps handlers thin — they call c.Error(err)
// and return, letting the middleware centralize error-to-HTTP translation.
// This guarantees every endpoint returns the same error envelope shape.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// After all handlers have run, check if any errors were collected.
		if len(c.Errors) == 0 {
			return
		}

		// Use the last error to determine the status code.
		// c.Errors is a slice — handlers can attach multiple errors if needed.
		err := c.Errors.Last().Err

		switch {
		case errors.Is(err, ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			// Don't leak internal error details to clients.
			slog.Error("unhandled error", "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}

// ---------------------------------------------------------------------------
// Middleware 4: TimeoutMiddleware
// ---------------------------------------------------------------------------

// TimeoutMiddleware wraps the request context with a deadline. Downstream
// handlers and service calls that check ctx.Done() or pass the context to
// database/HTTP clients will be cancelled when the deadline fires.
//
// Interview note: pass c.Request.Context() to your service layer, NOT
// *gin.Context. This keeps business logic framework-agnostic and ensures
// timeout/cancellation signals propagate through stdlib-compatible interfaces.
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a derived context with a timeout.
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace the request's context so downstream code uses the
		// deadline-aware context. c.Request.WithContext returns a shallow
		// copy of the request with the new context.
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// ---------------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------------

// profileHandler demonstrates reading user identity set by AuthMiddleware.
func profileHandler(c *gin.Context) {
	// c.MustGet panics if the key is missing — safe here because the auth
	// middleware guarantees "userID" is set for this route group.
	userID := c.MustGet("userID").(string)

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"message": fmt.Sprintf("Hello, user %s!", userID),
	})
}

// slowHandler simulates a long-running operation that respects context
// cancellation from the timeout middleware.
func slowHandler(c *gin.Context) {
	ctx := c.Request.Context()

	// Simulate work that checks the context. In real code this would be
	// a database query or HTTP call that accepts a context.
	select {
	case <-time.After(3 * time.Second):
		c.JSON(http.StatusOK, gin.H{"message": "completed (this is slow)"})
	case <-ctx.Done():
		// The timeout middleware's context expired. Return the error via
		// c.Error so the ErrorHandler middleware translates it.
		// Interview note: if you need the *gin.Context in a goroutine,
		// use c.Copy() — the original context is pooled and will be
		// recycled after the handler returns.
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "request timed out"})
	}
}

// failHandler demonstrates the c.Error() pattern — the handler attaches a
// domain error and returns without writing a response. The ErrorHandler
// middleware maps it to the correct HTTP status.
func failHandler(c *gin.Context) {
	// Simulate a service call that returns a domain error.
	err := fmt.Errorf("item xyz: %w", ErrNotFound)

	// c.Error appends to c.Errors without stopping the chain or writing
	// a response. The ErrorHandler middleware picks it up after c.Next().
	_ = c.Error(err)
}
