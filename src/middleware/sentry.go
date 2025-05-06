package middleware

import (
	"log/slog"

	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/config"
)

func Sentry() gin.HandlerFunc {
	sentryConfig := config.Config.Sentry
	// Initialize Sentry's handler
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:   sentryConfig.DSN,
		Debug: sentryConfig.Debug,
	}); err != nil {
		slog.Error("Sentry initialization failed", slog.Any("error", err))
	}

	return sentrygin.New(sentrygin.Options{})
}
