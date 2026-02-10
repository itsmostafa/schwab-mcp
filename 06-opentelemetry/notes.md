# OpenTelemetry

## Key Concepts

### The three pillars: traces, metrics, logs

- OpenTelemetry (OTel) is a vendor-neutral observability framework that standardises how telemetry is generated, collected, and exported.
- Three signal types:
  - **Traces** — end-to-end request flow across services, composed of spans.
  - **Metrics** — numeric measurements (counters, histograms, gauges) aggregated over time.
  - **Logs** — structured event records correlated with traces via trace/span IDs.
- In Go the trace and metric APIs are stable (v1.x); the log bridge API is newer and still evolving.
- Interview framing: traces tell you *where* time is spent, metrics tell you *how much* is happening, logs tell you *what* happened at a specific moment.

### API vs SDK — instrumentation vs export

- OTel separates the **API** (interfaces you instrument against) from the **SDK** (concrete implementations that process and export data).
- Library authors depend only on the API package (`go.opentelemetry.io/otel`). Application owners wire up the SDK at startup.
- This split means instrumentation code has zero vendor lock-in and stays stable even when export backends change.

| Layer | Package | Purpose |
|-------|---------|---------|
| API | `go.opentelemetry.io/otel` | `TracerProvider`, `Tracer`, `Span` interfaces; global getters/setters |
| API | `go.opentelemetry.io/otel/metric` | `MeterProvider`, `Meter`, instrument interfaces |
| SDK | `go.opentelemetry.io/otel/sdk/trace` | Concrete `TracerProvider`, span processors, samplers |
| SDK | `go.opentelemetry.io/otel/sdk/metric` | Concrete `MeterProvider`, readers, views |
| Export | `go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc` | OTLP gRPC trace exporter |
| Export | `go.opentelemetry.io/otel/exporters/stdout/stdouttrace` | Stdout trace exporter (dev/debug) |

- No-op default: if no SDK is registered, API calls are safe no-ops. This is why libraries can instrument without forcing a dependency on the SDK.

### TracerProvider, Tracer, and Span

- **TracerProvider**: top-level object that holds configuration (sampler, span processors, resource). Created once at application startup.
- **Tracer**: obtained from a `TracerProvider` and scoped to an instrumentation library/package name. Use `otel.Tracer("my/package")`.
- **Span**: represents a single unit of work. Created with `tracer.Start(ctx, "operation-name")`, which returns a new `context.Context` carrying the span.

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/codes"
)

var tracer = otel.Tracer("myapp/service/sync")

func SyncUser(ctx context.Context, userID string) error {
    ctx, span := tracer.Start(ctx, "SyncUser")
    defer span.End()

    span.SetAttributes(attribute.String("user.id", userID))

    if err := fetchFromAPI(ctx, userID); err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, err.Error())
        return err
    }
    span.SetStatus(codes.Ok, "")
    return nil
}
```

- **Attributes**: key-value metadata on a span (e.g. `http.method`, `db.statement`, custom business fields).
- **Events**: timestamped annotations within a span's lifetime; useful for recording exceptions or notable occurrences.
- **Status**: `Unset` (default), `Ok`, or `Error`. Only set `Error` explicitly; `Ok` is optional and cannot be downgraded once set.
- **Span kind**: `Server`, `Client`, `Producer`, `Consumer`, `Internal` — tells backends how to interpret the span.

### Parent-child relationships and trace context

- When you call `tracer.Start(ctx, ...)`, the new span automatically becomes a child of whatever span is already in `ctx`.
- This creates the trace tree: an incoming HTTP request span is the root; downstream DB calls and outbound HTTP calls are children.
- `trace.SpanFromContext(ctx)` retrieves the current span from a context.
- `trace.SpanContextFromContext(ctx)` retrieves the immutable `SpanContext` (trace ID, span ID, trace flags) without needing the full span.

### Context propagation across service boundaries

- Within a single process, span context travels via `context.Context`.
- Across services (HTTP, gRPC), span context must be serialized into headers and extracted on the other side.
- OTel uses **Propagators** for this. The standard is W3C `traceparent`/`tracestate` headers.

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/propagation"
)

// Set globally at startup
otel.SetTextMapPropagator(
    propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{},  // W3C traceparent
        propagation.Baggage{},       // W3C baggage
    ),
)
```

- **Injection**: before an outbound HTTP call, the propagator writes trace context into request headers.
- **Extraction**: on the receiving side, middleware extracts headers into a `context.Context` so child spans link correctly.
- The `otelhttp` transport and `otelgin` middleware handle this automatically.

### Setting up the SDK — TracerProvider and shutdown

- At application startup you wire together: resource, exporter, span processor, sampler, and `TracerProvider`.

```go
import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func initTracer(ctx context.Context) (func(context.Context) error, error) {
    exp, err := otlptracegrpc.New(ctx)
    if err != nil {
        return nil, err
    }

    res, err := resource.Merge(
        resource.Default(),
        resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceName("integration-sync-service"),
            semconv.ServiceVersion("1.0.0"),
        ),
    )
    if err != nil {
        return nil, err
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(res),
        sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(0.1))),
    )
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{},
        propagation.Baggage{},
    ))
    return tp.Shutdown, nil
}
```

- **BatchSpanProcessor** (via `WithBatcher`) is recommended for production — batches and sends asynchronously.
- **SimpleSpanProcessor** sends each span immediately — useful for development/testing only.
- Always call `tp.Shutdown(ctx)` on application exit to flush pending spans.

### Resource detection

- A **Resource** describes the entity producing telemetry (service name, version, host, container, cloud provider).
- `resource.Default()` detects environment automatically (process info, OS, `OTEL_RESOURCE_ATTRIBUTES` env var).
- `resource.Merge` combines default detection with explicit attributes.
- Semantic conventions (`semconv`) define standard attribute keys (`service.name`, `service.version`, `deployment.environment`, etc.).

### Samplers

- Sampling controls which traces are recorded and exported to manage cost and volume.
- Common samplers:

| Sampler | Behavior |
|---------|----------|
| `AlwaysSample` | Record every span (dev/test) |
| `NeverSample` | Drop everything |
| `TraceIDRatioBased(0.1)` | Sample 10% of traces deterministically |
| `ParentBased(root)` | Respect parent's sampling decision; apply `root` sampler for new traces |

- Production pattern: `ParentBased(TraceIDRatioBased(fraction))` — this respects upstream sampling decisions while applying ratio-based sampling to locally-originated traces.
- Head sampling (decide at span creation) vs tail sampling (decide after the trace completes, done in the collector) — tail sampling is more powerful but requires an OTel Collector.

### Exporters — OTLP, stdout, Prometheus

- **OTLP** (OpenTelemetry Protocol): the standard wire protocol. Supports gRPC and HTTP transports.
  - `otlptracegrpc` / `otlptracehttp` for traces.
  - `otlpmetricgrpc` / `otlpmetrichttp` for metrics.
  - Typically exports to an **OTel Collector**, which then fans out to Jaeger, Datadog, Grafana Tempo, etc.
- **Stdout**: prints telemetry to stdout; useful for local debugging.
- **Prometheus**: for metrics, use a Prometheus exporter that exposes a `/metrics` HTTP endpoint for scraping.
- Environment variable configuration: `OTEL_EXPORTER_OTLP_ENDPOINT`, `OTEL_EXPORTER_OTLP_PROTOCOL`, `OTEL_EXPORTER_OTLP_HEADERS` allow runtime configuration without code changes.

### Metrics — MeterProvider and instruments

- **MeterProvider**: analogous to `TracerProvider` for metrics. Created once, registered globally.
- **Meter**: obtained from `MeterProvider`, scoped to an instrumentation name.
- **Instruments**: the metric types you record with.

| Instrument | Use case | Example |
|------------|----------|---------|
| Counter | Monotonically increasing value | Total HTTP requests, bytes sent |
| UpDownCounter | Value that increases and decreases | Active connections, queue depth |
| Histogram | Distribution of values | Request latency, payload size |
| Gauge | Point-in-time measurement | CPU usage, temperature |

```go
import "go.opentelemetry.io/otel/metric"

var meter = otel.Meter("myapp/service/sync")

func initMetrics() error {
    syncCounter, err := meter.Int64Counter("sync.operations.total",
        metric.WithDescription("Total sync operations"),
    )
    if err != nil {
        return err
    }

    syncDuration, err := meter.Float64Histogram("sync.duration.seconds",
        metric.WithDescription("Duration of sync operations"),
        metric.WithUnit("s"),
    )
    if err != nil {
        return err
    }

    // Usage:
    // syncCounter.Add(ctx, 1, metric.WithAttributes(attribute.String("status", "success")))
    // syncDuration.Record(ctx, elapsed.Seconds())
    return nil
}
```

- **Readers**: control how metrics are collected.
  - `PeriodicReader` pushes metrics at an interval (used with OTLP exporters).
  - Prometheus exporter acts as a pull-based reader.
- **Views**: customize aggregation, rename instruments, or drop unwanted metrics.

### Gin middleware instrumentation (`otelgin`)

- The `otelgin` package from `opentelemetry-go-contrib` provides drop-in middleware for Gin.
- Automatically creates a server span for each incoming request, extracts propagated context from headers, and records HTTP metrics.

```go
import "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

router := gin.New()
router.Use(otelgin.Middleware("integration-sync-service"))
```

- Auto-recorded span attributes: `http.method`, `http.route`, `http.status_code`, `http.scheme`, `net.host.name`.
- Auto-recorded metrics: `http.server.request.duration`, `http.server.request.body.size`, `http.server.response.body.size`.
- Sets span status to `Error` for 5xx responses.
- Options: `WithTracerProvider`, `WithPropagators`, `WithFilter` (skip health checks), `WithSpanNameFormatter`.

### HTTP client instrumentation (`otelhttp`)

- Wraps `http.RoundTripper` to create client spans for outgoing HTTP calls and inject trace context into headers.

```go
import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

client := &http.Client{
    Transport: otelhttp.NewTransport(http.DefaultTransport),
}

// Every request made with this client now produces a client span
// and propagates trace context in headers.
resp, err := client.Do(req.WithContext(ctx))
```

- This is how distributed traces connect: the Gin middleware extracts incoming context, your service creates child spans, and the HTTP client injects context into outgoing requests.
- For the practice project, this instruments the third-party API calls with automatic span creation and context forwarding.

### GORM instrumentation

- The `otelgorm` plugin (from `github.com/uptrace/opentelemetry-go-extra/otelgorm`) adds tracing to GORM database operations.

```go
import "github.com/uptrace/opentelemetry-go-extra/otelgorm"

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
    return err
}
if err := db.Use(otelgorm.NewPlugin()); err != nil {
    return err
}
```

- Creates spans for every query with attributes like `db.system`, `db.statement`, `db.sql.table`.
- Pairs well with `db.WithContext(ctx)` — always pass context so spans link correctly in the trace tree.

### Baggage

- Baggage carries key-value pairs across service boundaries alongside trace context.
- Use cases: passing tenant ID, request priority, or feature flags without adding them to every function signature.
- Baggage is not the same as span attributes — it propagates to downstream services, while attributes are local to a span.
- Use sparingly: baggage is sent with every outbound request and adds overhead.

```go
import "go.opentelemetry.io/otel/baggage"

member, _ := baggage.NewMember("tenant.id", "acme-corp")
bag, _ := baggage.New(member)
ctx = baggage.ContextWithBaggage(ctx, bag)

// Downstream service:
bag := baggage.FromContext(ctx)
tenantID := bag.Member("tenant.id").Value()
```

### OTel Collector — the deployment pattern

- In production, applications export to an **OTel Collector** rather than directly to backends.
- The Collector receives, processes, and exports telemetry data.
- Pipeline: receivers → processors → exporters.
- Benefits:
  - Decouples application from backend (switch vendors without redeploying).
  - Offloads batching, retry, and enrichment from the application.
  - Enables tail sampling, attribute filtering, and fan-out to multiple backends.
- Deployment modes:
  - **Agent** (sidecar/daemonset): runs alongside each service, low-latency local collection.
  - **Gateway**: centralized collector that receives from agents, handles heavy processing.

### Putting it all together — instrumentation checklist

For the Integration Sync Service practice project, the instrumentation plan:

1. **Startup**: initialize `TracerProvider` and `MeterProvider` with OTLP exporters and proper resource attributes.
2. **Gin handlers**: add `otelgin.Middleware` for automatic server span creation.
3. **HTTP client**: wrap transport with `otelhttp.NewTransport` for outbound API call spans.
4. **GORM**: register `otelgorm` plugin for database query spans.
5. **Service layer**: create manual spans for business logic (e.g. `SyncUser`, `PublishEvent`).
6. **Metrics**: define counters for sync operations, histograms for latency, gauges for queue depth.
7. **Shutdown**: call `tp.Shutdown(ctx)` and `mp.Shutdown(ctx)` on graceful termination.
8. **Context discipline**: always pass `ctx` through every layer so spans connect into a single trace.

## Interview Questions

1. What are the three pillars of observability, and how does OpenTelemetry address each?
2. Why does OpenTelemetry separate the API from the SDK? What problem does this solve?
3. Walk through the lifecycle of a span: creation, enrichment, and export.
4. How does context propagation work across service boundaries in OTel?
5. What is W3C `traceparent`, and what information does it carry?
6. Explain the difference between `BatchSpanProcessor` and `SimpleSpanProcessor`. When would you use each?
7. What is a Resource in OTel, and why is it important?
8. Compare head sampling vs tail sampling. What are the tradeoffs?
9. How would you configure `ParentBased(TraceIDRatioBased(0.1))`, and what behavior does it produce?
10. What is the role of the OTel Collector, and why not export directly to a backend?
11. Describe how you would instrument a Gin-based service with both tracing and metrics.
12. How does `otelhttp.NewTransport` create distributed traces across HTTP calls?
13. What metric instruments does OTel provide, and when would you use each?
14. How does baggage differ from span attributes?
15. What happens if you forget to pass `ctx` when creating a span or making a DB call?
16. How would you set up OTel in a Go service to export traces to Jaeger and metrics to Prometheus simultaneously?
17. What is the purpose of semantic conventions, and why should you use them?

Practice prompts:
- Instrument the Integration Sync Service: add tracing to Gin handlers, outbound HTTP calls, and GORM queries so a single request produces a connected trace tree.
- Set up a `MeterProvider` with a Prometheus exporter and define custom metrics for sync success rate and latency distribution.
- Configure an OTel Collector pipeline that receives OTLP, applies tail sampling (keep all error traces, sample 10% of successful ones), and exports to Jaeger.

## Resources

- OpenTelemetry Go getting started: https://opentelemetry.io/docs/languages/go/getting-started/
- OpenTelemetry Go instrumentation guide: https://opentelemetry.io/docs/languages/go/instrumentation/
- OTel Go API reference: https://pkg.go.dev/go.opentelemetry.io/otel
- OTel Go SDK trace package: https://pkg.go.dev/go.opentelemetry.io/otel/sdk/trace
- OTel Go SDK metric package: https://pkg.go.dev/go.opentelemetry.io/otel/sdk/metric
- OTel Go contrib (otelgin, otelhttp): https://github.com/open-telemetry/opentelemetry-go-contrib
- otelgin middleware package: https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
- otelhttp package: https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
- otelgorm plugin: https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelgorm
- OpenTelemetry Collector docs: https://opentelemetry.io/docs/collector/
- W3C Trace Context specification: https://www.w3.org/TR/trace-context/
- Semantic conventions: https://opentelemetry.io/docs/specs/semconv/
- DeepWiki (opentelemetry-go): https://deepwiki.com/open-telemetry/opentelemetry-go
- DeepWiki (opentelemetry-go-contrib): https://deepwiki.com/open-telemetry/opentelemetry-go-contrib
