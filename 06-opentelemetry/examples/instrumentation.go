// instrumentation.go — Reference example: wiring OTel into a Gin + GORM app with otelgin,
// otelhttp, otelgorm, and baggage propagation.
//
// NOTE: This is a WIRING REFERENCE that shows how all the OTel instrumentation pieces
// connect in a realistic Gin+GORM service. It requires the full dependency set below
// to compile and run. Run it, then curl http://localhost:8080/users/42 to see the
// connected trace printed to stdout.
//
// Run:
//   go mod init example && go get \
//     go.opentelemetry.io/otel \
//     go.opentelemetry.io/otel/sdk \
//     go.opentelemetry.io/otel/exporters/stdout/stdouttrace \
//     go.opentelemetry.io/otel/exporters/stdout/stdoutmetric \
//     go.opentelemetry.io/otel/sdk/metric \
//     go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin \
//     go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp \
//     github.com/uptrace/opentelemetry-go-extra/otelgorm \
//     github.com/gin-gonic/gin \
//     gorm.io/gorm \
//     gorm.io/driver/sqlite \
//   && go run instrumentation.go

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

// ──────────────────────────────────────────────────────────────────────────────
// OTel setup — TracerProvider + MeterProvider with stdout exporters
// ──────────────────────────────────────────────────────────────────────────────

func initOTel() (func(context.Context) error, error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("user-service"),
		semconv.ServiceVersion("0.1.0"),
		semconv.DeploymentEnvironment("development"),
	)

	// --- Traces ---
	traceExp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(traceExp),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	// --- Metrics ---
	metricExp, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}
	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExp, metric.WithInterval(5*time.Second))),
		metric.WithResource(res),
	)
	otel.SetMeterProvider(mp)

	// --- Propagation ---
	// CompositeTextMapPropagator ensures both W3C trace context (traceparent)
	// and W3C baggage are injected/extracted from HTTP headers. Without this,
	// distributed traces break at service boundaries and baggage is lost.
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, // W3C traceparent/tracestate
		propagation.Baggage{},     // W3C baggage header
	))

	// Combined shutdown flushes both traces and metrics.
	shutdown := func(ctx context.Context) error {
		if err := tp.Shutdown(ctx); err != nil {
			return err
		}
		return mp.Shutdown(ctx)
	}
	return shutdown, nil
}

// ──────────────────────────────────────────────────────────────────────────────
// Database layer — GORM with otelgorm plugin
// ──────────────────────────────────────────────────────────────────────────────

// User is a simple GORM model for the demo.
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// otelgorm registers GORM callbacks that automatically create spans for
	// every query (Create, Query, Update, Delete, Row, Raw). Each span
	// includes attributes like db.system, db.statement, and db.sql.table.
	// Interview point: this is why you must always use db.WithContext(ctx) —
	// without the context, the plugin cannot link DB spans to the parent trace.
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		return nil, err
	}

	// Seed some data.
	db.AutoMigrate(&User{})
	db.Create(&User{ID: 42, Name: "Alice"})

	return db, nil
}

// UserRepo encapsulates database access.
type UserRepo struct {
	db *gorm.DB
}

// FindByID looks up a user. Note how ctx flows from the handler through the
// service into the repo — this is the critical context propagation chain that
// connects all spans into a single trace tree.
func (r *UserRepo) FindByID(ctx context.Context, id string) (*User, error) {
	var user User
	// WithContext passes the trace context to otelgorm so the DB span becomes
	// a child of the calling span. Forgetting this is a common mistake that
	// creates orphaned spans.
	result := r.db.WithContext(ctx).First(&user, id)
	return &user, result.Error
}

// ──────────────────────────────────────────────────────────────────────────────
// HTTP client — wrapped with otelhttp for outbound call tracing
// ──────────────────────────────────────────────────────────────────────────────

// newTracedHTTPClient creates an http.Client whose transport automatically:
//  1. Creates a client span for every outbound request.
//  2. Injects the current trace context into request headers (via the global
//     propagator) so the downstream service can continue the trace.
//
// Interview point: this is how distributed tracing works across services.
// otelhttp.NewTransport wraps the RoundTripper to handle injection.
func newTracedHTTPClient() *http.Client {
	return &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
}

// ──────────────────────────────────────────────────────────────────────────────
// Service layer — business logic with manual spans
// ──────────────────────────────────────────────────────────────────────────────

type UserService struct {
	repo   *UserRepo
	client *http.Client
	tracer trace.Tracer
}

// GetUser demonstrates the full instrumented request flow:
// Gin handler -> Service (manual span) -> Repo (otelgorm span) -> HTTP client (otelhttp span).
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	// Create a manual span for business logic. Instrumentation libraries
	// (otelgin, otelgorm, otelhttp) auto-create spans for framework
	// operations, but you need manual spans for your own logic.
	ctx, span := s.tracer.Start(ctx, "UserService.GetUser")
	defer span.End()

	// Read baggage that was set by the handler. Baggage propagates across
	// service boundaries via HTTP headers, unlike span attributes which
	// are local to one span.
	bag := baggage.FromContext(ctx)
	tenantID := bag.Member("tenant.id").Value()
	span.SetAttributes(attribute.String("tenant.id", tenantID))

	// Database call — otelgorm creates a child span automatically because
	// we pass ctx through.
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Outbound HTTP call — otelhttp creates a client span and injects trace
	// context into headers. The downstream service (if instrumented) will
	// extract it and continue the trace.
	s.enrichFromExternalAPI(ctx, id)

	return user, nil
}

// enrichFromExternalAPI simulates calling a third-party API. The traced HTTP
// client automatically creates a span and propagates context.
func (s *UserService) enrichFromExternalAPI(ctx context.Context, userID string) {
	ctx, span := s.tracer.Start(ctx, "UserService.enrichFromExternalAPI")
	defer span.End()

	// The request is made with ctx so otelhttp can extract the parent span
	// and inject traceparent/baggage headers into the outgoing request.
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/get", nil)
	resp, err := s.client.Do(req)
	if err != nil {
		span.SetAttributes(attribute.String("enrichment.error", err.Error()))
		return
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)

	span.SetAttributes(attribute.Int("enrichment.status_code", resp.StatusCode))
}

// ──────────────────────────────────────────────────────────────────────────────
// Gin handlers — with otelgin middleware and baggage
// ──────────────────────────────────────────────────────────────────────────────

func main() {
	shutdown, err := initOTel()
	if err != nil {
		log.Fatalf("failed to init otel: %v", err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("otel shutdown error: %v", err)
		}
	}()

	db, err := initDB()
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	svc := &UserService{
		repo:   &UserRepo{db: db},
		client: newTracedHTTPClient(),
		tracer: otel.Tracer("user-service"),
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// otelgin.Middleware is the key integration point for Gin. It:
	//  - Extracts incoming trace context from request headers (so this
	//    service continues a trace started by an upstream caller).
	//  - Creates a server span for each request with attributes like
	//    http.method, http.route, http.status_code.
	//  - Sets span status to Error for 5xx responses.
	//  - Records http.server.request.duration metrics.
	// Interview point: without this middleware, every incoming request
	// starts a new trace even if the caller sent a traceparent header.
	router.Use(otelgin.Middleware("user-service"))

	router.GET("/users/:id", func(c *gin.Context) {
		// The otelgin middleware already created a server span and stored
		// it in c.Request.Context(). All subsequent operations that
		// receive this context become children of that span.
		ctx := c.Request.Context()

		// --- Baggage ---
		// Set baggage on the context. Baggage key-value pairs propagate
		// to downstream services via the W3C baggage header. Use cases:
		// tenant ID, request priority, feature flags.
		// Interview point: baggage is NOT the same as span attributes.
		// Attributes are local to one span; baggage travels across
		// service boundaries. Use sparingly — every member is sent with
		// every outbound request.
		member, _ := baggage.NewMember("tenant.id", "acme-corp")
		bag, _ := baggage.New(member)
		ctx = baggage.ContextWithBaggage(ctx, bag)

		// Pass ctx through the full call chain: handler -> service -> repo + HTTP client.
		// This is the context propagation discipline that makes distributed tracing work.
		user, err := svc.GetUser(ctx, c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   user.ID,
			"name": user.Name,
		})
	})

	// Start the server. Try: curl http://localhost:8080/users/42
	// The stdout exporter will print the full trace tree showing:
	//   otelgin server span
	//     -> UserService.GetUser (manual span)
	//       -> GORM query span (otelgorm)
	//       -> UserService.enrichFromExternalAPI (manual span)
	//         -> HTTP GET httpbin.org (otelhttp client span)
	fmt.Println("Server running on :8080 — try: curl http://localhost:8080/users/42")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
