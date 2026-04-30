package otel

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

// commonResource returns a common resource with service name, OS, and arch.
func commonResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", "go-xn"),
			attribute.String("os.type", runtime.GOOS),
			attribute.String("os.arch", runtime.GOARCH),
		),
		resource.WithHost(),
	)
}
