package otel

import (
	"context"
	"log/slog"
	"runtime"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

type otelVersionCtx struct{}

// OtelVersionKey is the context key for storing the service version for OpenTelemetry.
var OtelVersionKey = otelVersionCtx{}

// commonResource returns a common resource with service name, OS, and arch.
func commonResource(ctx context.Context) (*resource.Resource, error) {
	version, ok := ctx.Value(OtelVersionKey).(string)
	if !ok {
		slog.Debug("otel version not found in context", slog.String("version", "unknown"))
		version = "unknown"
	}

	return resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithAttributes(
			attribute.String("service.name", "go-xn"),
			attribute.String("service.version", version),
			attribute.String("os.type", runtime.GOOS),
			attribute.String("os.arch", runtime.GOARCH),
		),
	)
}
