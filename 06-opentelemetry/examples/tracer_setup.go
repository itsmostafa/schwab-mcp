// tracer_setup.go — TracerProvider setup with stdout exporter, spans, attributes, events, and error handling.
//
// Run: go mod init example && go get go.opentelemetry.io/otel go.opentelemetry.io/otel/sdk go.opentelemetry.io/otel/exporters/stdout/stdouttrace && go run tracer_setup.go

package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

// initTracer sets up the OTel tracing pipeline and returns a shutdown function.
// Interview point: the API vs SDK separation means libraries instrument against
// the API package (go.opentelemetry.io/otel) while only the application owner
// imports the SDK. If no SDK is registered, all API calls are safe no-ops.
func initTracer() (func(context.Context) error, error) {
	// stdout exporter prints span JSON to the terminal — great for local dev.
	// In production you would use an OTLP exporter (otlptracegrpc or otlptracehttp)
	// that sends spans to an OTel Collector.
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, fmt.Errorf("creating stdout exporter: %w", err)
	}

	// Resource describes the entity producing telemetry. Semantic conventions
	// (semconv) define standard attribute keys so backends can interpret them
	// uniformly across services. Always set at least service.name.
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("example-service"),
		semconv.ServiceVersion("0.1.0"),
		semconv.DeploymentEnvironment("development"),
	)

	// TracerProvider is the central factory — it holds configuration (resource,
	// sampler, span processors) and creates Tracers.
	tp := sdktrace.NewTracerProvider(
		// WithSyncer wraps the exporter in a SimpleSpanProcessor, which
		// exports each span synchronously when it ends. This blocks the
		// caller, so it is suitable ONLY for development and testing.
		// Production: use sdktrace.WithBatcher(exporter) instead — it
		// batches spans and sends them asynchronously for much better
		// throughput and lower latency impact.
		sdktrace.WithSyncer(exporter),
		sdktrace.WithResource(res),
	)

	// Register as the global TracerProvider so that any package calling
	// otel.Tracer("...") receives a real tracer rather than a no-op.
	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}

func main() {
	shutdown, err := initTracer()
	if err != nil {
		log.Fatalf("failed to init tracer: %v", err)
	}
	// Always shut down the TracerProvider on exit so that buffered spans
	// are flushed to the exporter. For BatchSpanProcessor this is critical —
	// without it you may lose the last batch of spans.
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown tracer provider: %v", err)
		}
	}()

	// Obtain a Tracer scoped to this component. The name conventionally
	// matches the instrumentation library or package. TracerProvider is the
	// factory; Tracer is per-library/component.
	tracer := otel.Tracer("example-service")

	ctx := context.Background()

	// --- Root span ---
	// tracer.Start returns a new context carrying the span. This context is
	// what links child spans to their parent, forming the trace tree.
	ctx, rootSpan := tracer.Start(ctx, "ProcessOrder")
	// Always defer span.End() immediately after Start. If you forget, the
	// span is never exported and you get a broken trace.
	defer rootSpan.End()

	// Add attributes to the span — key-value metadata. Use attributes for
	// dimensions you want to filter/group by in your tracing backend.
	rootSpan.SetAttributes(
		attribute.String("order.id", "ORD-12345"),
		attribute.String("customer.id", "CUST-42"),
		attribute.Int("order.item_count", 3),
	)

	// Add an event — a timestamped log-like annotation within the span's
	// lifetime. Events are useful for marking milestones without creating
	// a separate child span.
	rootSpan.AddEvent("order validated",
		// Events can also carry attributes via trace.WithAttributes.
		trace.WithAttributes(attribute.String("validation.result", "passed")),
	)

	// --- Child span: demonstrates nesting ---
	// Because we pass the ctx that carries rootSpan, this new span becomes
	// a child. Spans form a tree via context — this is how distributed
	// traces show the call hierarchy.
	if err := chargePayment(ctx, tracer, "CUST-42", 99.99); err != nil {
		// RecordError adds an exception event to the span with the error
		// message and stack trace. This makes errors visible in the trace
		// UI without manually extracting error details.
		rootSpan.RecordError(err)

		// SetStatus marks the span as errored. Tracing backends use this
		// to highlight failed spans. Only set Error explicitly; Ok is
		// optional and cannot be downgraded once set.
		rootSpan.SetStatus(codes.Error, err.Error())
		fmt.Println("Order processing failed:", err)
		return
	}

	// Ship the order — another child span.
	shipOrder(ctx, tracer, "ORD-12345")

	rootSpan.AddEvent("order processing complete")
	rootSpan.SetStatus(codes.Ok, "")
	fmt.Println("Order processed successfully")
}

// chargePayment creates a child span and simulates a payment failure.
func chargePayment(ctx context.Context, tracer trace.Tracer, customerID string, amount float64) error {
	// Child span: linked to the parent via ctx. We use _ because this function
	// has no deeper calls that need the child context. In real code you would
	// pass it down to further operations.
	_, span := tracer.Start(ctx, "ChargePayment")
	defer span.End() // Always defer End immediately.

	span.SetAttributes(
		attribute.String("payment.customer_id", customerID),
		attribute.Float64("payment.amount", amount),
		attribute.String("payment.currency", "USD"),
	)

	span.AddEvent("contacting payment gateway")

	// Simulate a failure to demonstrate error recording.
	err := errors.New("payment gateway timeout")

	// Record the error on the child span itself so the trace shows exactly
	// which operation failed and why.
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())

	return err
}

// shipOrder creates a child span for a successful operation.
func shipOrder(ctx context.Context, tracer trace.Tracer, orderID string) {
	// Even though chargePayment failed in this demo flow, this function
	// illustrates what a successful child span looks like.
	_, span := tracer.Start(ctx, "ShipOrder")
	defer span.End()

	span.SetAttributes(
		attribute.String("shipping.order_id", orderID),
		attribute.String("shipping.carrier", "FastShip"),
	)

	span.AddEvent("shipping label created")
	span.SetStatus(codes.Ok, "")
}
