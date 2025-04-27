package otel

import (
	"context"
	"log"
	"log/slog"
	"runtime"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func InitLog(ctx context.Context, c Config) *slog.Logger {
	var exporter sdklog.Exporter
	var err error
	switch c.ClientType {
	case "grpc":
		slog.Debug("init otel log grpc client")
		exporter, err = newGRPCExporter(ctx, c)
	case "http":
		slog.Debug("init otel log http client")
		exporter, err = newHTTPExporter(ctx, c)
	default:
		slog.Warn("unknown otel log client", slog.String("client", c.ClientType))
		slog.Debug("init otel log client with default grpc")
		exporter, err = newGRPCExporter(ctx, c)
	}
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

	// Initialize the logger provider
	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(resources),
	)

	logger := otelslog.NewLogger("go-xn",
		otelslog.WithLoggerProvider(loggerProvider),
	)

	return logger
}

// newGRPCExporter creates a new gRPC exporter for OpenTelemetry logs.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-logs-over-grpc-experimental
func newGRPCExporter(ctx context.Context, c Config) (*otlploggrpc.Exporter, error) {
	return otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint(c.Endpoint),
		otlploggrpc.WithHeaders(c.Headers),
	)
}

// newHTTPExporter creates a new HTTP exporter for OpenTelemetry logs.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-logs-over-http-experimental
func newHTTPExporter(ctx context.Context, c Config) (*otlploghttp.Exporter, error) {
	return otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(c.Endpoint),
		otlploghttp.WithHeaders(c.Headers),
	)
}
