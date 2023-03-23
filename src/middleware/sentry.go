package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/util/log"
)

func Sentry() gin.HandlerFunc {
	sentryConfig := config.Config.Sentry
	// Initialize Sentry's handler
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:   sentryConfig.DSN,
		Debug: sentryConfig.Debug,
	}); err != nil {
		log.Errorf("Sentry initialization failed: %v", err)
	}

	return sentrygin.New(sentrygin.Options{})
}
