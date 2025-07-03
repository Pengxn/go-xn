package otel

import (
	"context"
	"log"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func NewLogger(ctx context.Context, c *config) *slog.Logger {
	loggerProvider := InitLog(ctx, c)

	logger := otelslog.NewLogger("go-xn",
		otelslog.WithLoggerProvider(loggerProvider),
	)

	return logger
}

func InitLog(ctx context.Context, c *config) *sdklog.LoggerProvider {
	var exporterFn func(context.Context, config) (sdklog.Exporter, error)
	switch c.ClientType {
	case "grpc":
		slog.Debug("init otel log", slog.String("type", c.ClientType))
		exporterFn = newGRPCLogExporter
	case "http":
		slog.Debug("init otel log", slog.String("type", c.ClientType))
		exporterFn = newHTTPLogExporter
	case "stdout":
		slog.Debug("init otel log", slog.String("type", c.ClientType))
		exporterFn = newStdoutLogExporter
	default:
		slog.Warn("unknown otel log type", slog.String("type", c.ClientType))
		slog.Debug("init otel log with default stdout")
		exporterFn = newStdoutLogExporter
	}
	exporter, err := exporterFn(ctx, *c)
	if err != nil {
		log.Fatalf("failed to create log exporter: %s", err)
	}

	// create the resource
	resources, err := commonResource(ctx)
	if err != nil {
		log.Fatalf("failed to set resources: %s", err)
	}

	// Initialize the logger provider
	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(resources),
	)

	return loggerProvider
}

// newStdoutExporter creates a new stdout exporter for OpenTelemetry logs.
// https://opentelemetry.io/docs/languages/go/exporters/#console-logs
func newStdoutLogExporter(_ context.Context, _ config) (sdklog.Exporter, error) {
	return stdoutlog.New(
		stdoutlog.WithPrettyPrint(),
		stdoutlog.WithWriter(log.Writer()), // TODO: use custom writer
	)
}

// newGRPCExporter creates a new gRPC exporter for OpenTelemetry logs.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-logs-over-grpc-experimental
func newGRPCLogExporter(ctx context.Context, c config) (sdklog.Exporter, error) {
	return otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint(c.Endpoint),
		otlploggrpc.WithHeaders(c.Headers),
	)
}

// newHTTPExporter creates a new HTTP exporter for OpenTelemetry logs.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-logs-over-http-experimental
func newHTTPLogExporter(ctx context.Context, c config) (sdklog.Exporter, error) {
	return otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(c.Endpoint),
		otlploghttp.WithHeaders(c.Headers),
	)
}
