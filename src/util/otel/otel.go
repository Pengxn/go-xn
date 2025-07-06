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

	// Initialize OpenTelemetry tracing
	traceShutdown := InitTrace(ctx, cfg)

	// Initialize OpenTelemetry metrics
	metricShutdown := InitMetric(ctx, cfg)

	// Initialize OpenTelemetry logging
	_ = NewLogger(ctx, cfg)

	return func(ctx context.Context) {
		err := traceShutdown(ctx)
		if err != nil {
			slog.Error("failed to shutdown otel trace", slog.Any("error", err))
		}
		err = metricShutdown(ctx)
		if err != nil {
			slog.Error("failed to shutdown otel metric", slog.Any("error", err))
		}
	}
}
