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
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTrace(ctx context.Context, c Config) func(context.Context) error {
	var client otlptrace.Client
	switch c.ClientType {
	case "grpc":
		slog.Debug("init grpc otel client")
		client = newGRPCClient(c.Endpoint, c.Headers)
	case "http":
		slog.Debug("init https otel client")
		client = newHTTPClient(c.Endpoint, c.Headers)
	default:
		slog.Warn("unknown otel client", slog.String("client", c.ClientType))
	}

	return initOTELTracer(ctx, client)
}

type Config struct {
	ClientType string
	Endpoint   string
	Headers    map[string]string
}

func newGRPCClient(endpoint string, headers map[string]string) otlptrace.Client {
	return otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithHeaders(headers),
	)
}

func newHTTPClient(endpoint string, headers map[string]string) otlptrace.Client {
	return otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithHeaders(headers),
	)
}

func initOTELTracer(ctx context.Context, client otlptrace.Client) func(context.Context) error {
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatalf("failed to create exporter: %s", err)
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", "go-xn"),
			attribute.String("service.os", runtime.GOOS),
			attribute.String("service.arch", runtime.GOARCH),
		),
	)
	if err != nil {
		log.Fatalf("failed to set resources: %s", err)
	}

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
