package otel

import (
	"context"
	"log"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitTrace initializes the OpenTelemetry trace exporter with the given config.
// It returns a function to shut down the exporter when done.
func InitTrace(ctx context.Context, c Config) func(context.Context) error {
	var exporterFn exporterFunc
	switch c.ClientType {
	case "grpc":
		slog.Debug("init otel trace", slog.String("type", c.ClientType))
		exporterFn = newGRPCTraceExporter
	case "http":
		slog.Debug("init otel trace", slog.String("type", c.ClientType))
		exporterFn = newHTTPTraceExporter
	case "stdout":
		slog.Debug("init otel trace", slog.String("type", c.ClientType))
		exporterFn = newStdoutTraceExporter
	default:
		slog.Warn("unknown otel trace type", slog.String("type", c.ClientType))
		slog.Debug("init otel trace with default stdout")
		exporterFn = newStdoutTraceExporter
	}

	return initOTELTracer(ctx, c, exporterFn)
}

// exporterFunc is a function type that takes a context and config,
// it's used to create a new OpenTelemetry trace exporter.
type exporterFunc func(context.Context, Config) (trace.SpanExporter, error)

// newStdoutExporter creates a new stdout exporter for OpenTelemetry traces.
// https://opentelemetry.io/docs/languages/go/exporters/#console-traces
func newStdoutTraceExporter(_ context.Context, _ Config) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(log.Writer()), // TODO: use custom writer
	)
}

// newGRPCExporter creates a new gRPC exporter for OpenTelemetry traces.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-traces-over-grpc
func newGRPCTraceExporter(ctx context.Context, c Config) (trace.SpanExporter, error) {
	return otlptrace.New(ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithEndpoint(c.Endpoint),
			otlptracegrpc.WithHeaders(c.Headers),
		),
	)
}

// newHTTPExporter creates a new HTTP exporter for OpenTelemetry traces.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-traces-over-http
func newHTTPTraceExporter(ctx context.Context, c Config) (trace.SpanExporter, error) {
	return otlptrace.New(ctx,
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(c.Endpoint),
			otlptracehttp.WithHeaders(c.Headers),
		),
	)
}

// initOTELTracer initializes the OpenTelemetry tracer with the given client.
func initOTELTracer(ctx context.Context, c Config, fn exporterFunc) func(context.Context) error {
	// create the exporter
	exporter, err := fn(ctx, c)
	if err != nil {
		log.Fatalf("failed to create tracer exporter: %s", err)
	}

	// create the resource
	resources, err := commonResource(ctx)
	if err != nil {
		log.Fatalf("failed to set resources: %s", err)
	}

	// set the global OpenTelemetry tracer provider
	otel.SetTracerProvider(
		trace.NewTracerProvider(
			trace.WithSampler(trace.AlwaysSample()),
			trace.WithSpanProcessor(trace.NewBatchSpanProcessor(exporter)),
			trace.WithSyncer(exporter),
			trace.WithResource(resources),
		),
	)

	return exporter.Shutdown
}
