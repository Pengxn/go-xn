package otel

import (
	"context"
	"log/slog"

	commonConfig "github.com/Pengxn/go-xn/src/config"
)

func SetOtel(ctx context.Context, c commonConfig.OtelConfig) func(ctx context.Context) {
	// OpenTelemetry is disabled if ClientType is empty
	if c.ClientType == "" {
		slog.Debug("OpenTelemetry is disabled as clientType is empty")
		return func(ctx context.Context) {}
	}

	cfg := NewConfig(
		WithClientType(c.ClientType),
		WithEndpoint(c.Endpoint),
		WithHeaders(map[string]string{
			c.Header: c.Token,
		}))

	shutdown := []func(context.Context){}

	// Enable and initialize OpenTelemetry tracing
	if c.EnableTrace {
		traceShutdown := InitTrace(ctx, cfg)
		// Add the trace shutdown function to the list
		shutdown = append(shutdown, func(ctx context.Context) {
			err := traceShutdown(ctx)
			if err != nil {
				slog.Error("failed to shutdown otel trace", slog.Any("error", err))
			}
		})
	}

	// Enable and initialize OpenTelemetry metrics
	if c.EnableMetric {
		metricShutdown := InitMetric(ctx, cfg)
		// Add the metric shutdown function to the list
		shutdown = append(shutdown, func(ctx context.Context) {
			err := metricShutdown(ctx)
			if err != nil {
				slog.Error("failed to shutdown otel metric", slog.Any("error", err))
			}
		})
	}

	// Enable and initialize OpenTelemetry logging
	if c.EnableLog {
		logger := NewLogger(ctx, cfg)
		slog.SetDefault(logger)
	}

	return func(ctx context.Context) {
		for _, fn := range shutdown {
			fn(ctx)
		}
	}
}
