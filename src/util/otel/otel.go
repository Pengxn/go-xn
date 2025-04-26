package otel

import (
	"context"
	"log"
	"log/slog"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
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

// Config is the configuration for [OpenTelemetry].
// It contains the client type, endpoint, and headers for the exporter.
// The client type can be either "grpc" or "http".
// The endpoint is the URL of the [OpenTelemetry Collector], default not including `v1/trace`.
// refer to https://opentelemetry.io/docs/languages/sdk-configuration/otlp-exporter/
//
// [OpenTelemetry]: https://opentelemetry.io/
// [OpenTelemetry Collector]: https://opentelemetry.io/docs/collector/
type Config struct {
	ClientType string
	Endpoint   string
	Headers    map[string]string
}

type exporterFunc func(context.Context, Config) (trace.SpanExporter, error)

// newStdoutExporter creates a new stdout exporter for OpenTelemetry traces.
// https://opentelemetry.io/docs/languages/go/exporters/#console-traces
func newStdoutTraceExporter(_ context.Context, _ Config) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)
}

// newGRPCExporter creates a new gRPC exporter for OpenTelemetry traces.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-traces-over-grpc
func newGRPCTraceExporter(ctx context.Context, c Config) (trace.SpanExporter, error) {
	return otlptrace.New(ctx,
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(c.Endpoint),
			otlptracehttp.WithHeaders(c.Headers),
		),
	)
}

// newHTTPExporter creates a new HTTP exporter for OpenTelemetry traces.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-traces-over-http
func newHTTPTraceExporter(ctx context.Context, c Config) (trace.SpanExporter, error) {
	return otlptrace.New(ctx,
		otlptracegrpc.NewClient(
			otlptracegrpc.WithEndpoint(c.Endpoint),
			otlptracegrpc.WithHeaders(c.Headers),
		),
	)
}

// initOTELTracer initializes the OpenTelemetry tracer with the given client.
func initOTELTracer(ctx context.Context, c Config, fn exporterFunc) func(context.Context) error {
	// create the exporter
	exporter, err := fn(ctx, c)
	if err != nil {
		log.Fatalf("failed to create exporter: %s", err)
	}

	// create the resource
	resources, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", "go-xn"),
			attribute.String("service.os", runtime.GOOS),
			attribute.String("service.arch", runtime.GOARCH),
		),
	)
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
