package otel

import (
	"context"
	"log"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

// InitTrace initializes the OpenTelemetry trace exporter with the given config.
// It returns a function to shut down the exporter when done.
func InitMetric(ctx context.Context, c *config) func(context.Context) error {
	var metricExporter metricFunc
	switch c.ClientType {
	case "grpc":
		slog.Debug("init otel metric", slog.String("type", c.ClientType))
		metricExporter = newGRPCMetricExporter
	case "http":
		slog.Debug("init otel metric", slog.String("type", c.ClientType))
		metricExporter = newHTTPMetricExporter
	case "stdout":
		slog.Debug("init otel metric", slog.String("type", c.ClientType))
		metricExporter = newStdoutMetricExporter
	default:
		slog.Warn("unknown otel metric type", slog.String("type", c.ClientType))
		slog.Debug("init otel metric with default stdout")
		metricExporter = newStdoutMetricExporter
	}

	exporter, err := metricExporter(ctx, *c)
	if err != nil {
		log.Fatal("failed to create metric exporter: ", err)
	}

	// Create new resource
	resources, err := commonResource(ctx)
	if err != nil {
		log.Fatal("failed to create resource: ", err)
	}

	otel.SetMeterProvider(
		metric.NewMeterProvider(
			metric.WithReader(metric.NewPeriodicReader(exporter)),
			metric.WithResource(resources),
		),
	)

	return exporter.Shutdown
}

// metricFunc is a function that creates an OpenTelemetry exporter.
type metricFunc func(context.Context, config) (metric.Exporter, error)

// newStdoutMetricExporter creates a new stdout exporter for OpenTelemetry metrics.
// https://opentelemetry.io/docs/languages/go/exporters/#console-metrics
func newStdoutMetricExporter(_ context.Context, _ config) (metric.Exporter, error) {
	return stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(),
		stdoutmetric.WithWriter(log.Writer()), // TODO: use custom writer
	)
}

// newGRPCMetricExporter creates a new gRPC exporter for OpenTelemetry metrics.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-metrics-over-grpc
func newGRPCMetricExporter(ctx context.Context, c config) (metric.Exporter, error) {
	return otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(c.Endpoint),
		otlpmetricgrpc.WithHeaders(c.Headers),
	)
}

// newHTTPMetricExporter creates a new HTTP exporter for OpenTelemetry metrics.
// https://opentelemetry.io/docs/languages/go/exporters/#otlp-metrics-over-http
func newHTTPMetricExporter(ctx context.Context, c config) (metric.Exporter, error) {
	return otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(c.Endpoint),
		otlpmetrichttp.WithHeaders(c.Headers),
	)
}
