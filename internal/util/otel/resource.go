package otel

import (
	"context"
	"log/slog"
	"runtime"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Service semantic convention keys
const (
	// ServiceNameKey is the attribute Key conforming to the "service.name"
	// semantic conventions. It represents the logical name of the service.
	ServiceNameKey = attribute.Key("service.name")

	// ServiceVersionKey is the attribute Key conforming to the "service.version"
	// semantic conventions. It represents the version string of the service
	// component. The format is not defined by these conventions.
	ServiceVersionKey = attribute.Key("service.version")
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
			ServiceNameKey.String("go-xn"),
			ServiceVersionKey.String(version),
			attribute.String("os.type", runtime.GOOS),
			attribute.String("os.arch", runtime.GOARCH),
		),
	)
}
