// metrics.go — MeterProvider setup with stdout exporter, Counter, Histogram, and UpDownCounter.
//
// Run: go mod init example && go get go.opentelemetry.io/otel go.opentelemetry.io/otel/sdk/metric go.opentelemetry.io/otel/exporters/stdout/stdoutmetric && go run metrics.go

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	otelmetric "go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// initMeter sets up the OTel metrics pipeline and returns a shutdown function.
func initMeter() (func(context.Context) error, error) {
	// stdout metric exporter prints aggregated metrics as JSON to the terminal.
	// In production you would use an OTLP exporter or a Prometheus exporter.
	exporter, err := stdoutmetric.New()
	if err != nil {
		return nil, fmt.Errorf("creating stdout metric exporter: %w", err)
	}

	// Resource identifies the service producing metrics. Backends group and
	// filter metrics by resource attributes, so always set service.name.
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("example-service"),
		semconv.ServiceVersion("0.1.0"),
	)

	// PeriodicReader collects metrics from all registered instruments at a
	// fixed interval and pushes them to the exporter. The interval controls
	// how often you see output. 3 seconds is fast for demo purposes;
	// production defaults are typically 30-60 seconds.
	reader := sdkmetric.NewPeriodicReader(exporter,
		sdkmetric.WithInterval(3*time.Second),
	)

	// MeterProvider is analogous to TracerProvider — created once at startup,
	// registered globally, and used to obtain Meters.
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(reader),
		sdkmetric.WithResource(res),
	)

	// Register as global so any package calling otel.Meter("...") gets a real
	// meter rather than a no-op.
	otel.SetMeterProvider(mp)

	return mp.Shutdown, nil
}

func main() {
	shutdown, err := initMeter()
	if err != nil {
		log.Fatalf("failed to init meter: %v", err)
	}
	// Shutdown flushes any pending metrics. Without this you may lose the
	// last collection interval's data.
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown meter provider: %v", err)
		}
	}()

	// Obtain a Meter scoped to this component. Like Tracer, the name should
	// match the instrumentation library or package.
	meter := otel.Meter("example-service")

	// --- Counter (monotonic) ---
	// Counters only go up. Use them for things you count: requests, errors,
	// bytes sent. Interview point: Counter.Add() panics if you pass a
	// negative value — that is what UpDownCounter is for.
	requestCounter, err := meter.Int64Counter("http_requests_total",
		otelmetric.WithDescription("Total number of HTTP requests received"),
		otelmetric.WithUnit("{request}"),
	)
	if err != nil {
		log.Fatalf("failed to create counter: %v", err)
	}

	// --- Histogram ---
	// Histograms capture the distribution of values (latency, payload size).
	// The SDK aggregates values into buckets so you can compute percentiles
	// (p50, p95, p99) on the backend. Interview point: histograms are more
	// expensive than counters because they track distribution, not just sum.
	requestDuration, err := meter.Float64Histogram("http_request_duration_seconds",
		otelmetric.WithDescription("Duration of HTTP requests in seconds"),
		otelmetric.WithUnit("s"),
	)
	if err != nil {
		log.Fatalf("failed to create histogram: %v", err)
	}

	// --- UpDownCounter ---
	// Unlike Counter, UpDownCounter can increase and decrease. Use it for
	// values that fluctuate: active connections, queue depth, in-flight
	// requests. Interview point: UpDownCounter accepts negative values in
	// Add(), Counter does not.
	activeConnections, err := meter.Int64UpDownCounter("active_connections",
		otelmetric.WithDescription("Number of currently active connections"),
		otelmetric.WithUnit("{connection}"),
	)
	if err != nil {
		log.Fatalf("failed to create up-down counter: %v", err)
	}

	ctx := context.Background()

	// Simulate traffic in a loop so the PeriodicReader has data to export.
	// Each iteration represents an incoming HTTP request.
	fmt.Println("Generating metrics... (wait ~3 seconds for first export)")
	for i := 0; i < 10; i++ {
		// Attributes (also called labels) add dimensions to measurements.
		// They let you break down metrics by method, status, endpoint, etc.
		// Keep cardinality low — high-cardinality attributes (e.g., user ID)
		// create too many time series and blow up storage costs.
		method := "GET"
		if i%3 == 0 {
			method = "POST"
		}
		status := "200"
		if i%5 == 0 {
			status = "500"
		}

		attrs := otelmetric.WithAttributes(
			attribute.String("http.method", method),
			attribute.String("http.status_code", status),
			attribute.String("http.route", "/api/users"),
		)

		// Record a request.
		requestCounter.Add(ctx, 1, attrs)

		// Record a latency measurement. Different values show how the
		// histogram captures distribution — the backend can derive p50,
		// p95, p99 from the bucket counts.
		duration := 0.05 + rand.Float64()*0.45 // 50ms to 500ms
		requestDuration.Record(ctx, duration, attrs)

		// Simulate connection lifecycle: open a connection, then close it
		// after a brief pause. The UpDownCounter reflects the current count.
		activeConnections.Add(ctx, 1, otelmetric.WithAttributes(
			attribute.String("protocol", "http"),
		))

		time.Sleep(200 * time.Millisecond)

		// Close the connection — subtract 1.
		activeConnections.Add(ctx, -1, otelmetric.WithAttributes(
			attribute.String("protocol", "http"),
		))
	}

	// Wait for the PeriodicReader to export at least once more after the loop.
	fmt.Println("Waiting for final metric export...")
	time.Sleep(4 * time.Second)
	fmt.Println("Done")
}
