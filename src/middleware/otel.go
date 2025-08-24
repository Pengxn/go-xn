package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Otel(ctx context.Context) gin.HandlerFunc {
	return otelgin.Middleware("go-xn")
}
