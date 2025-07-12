package route

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/controller"
	"github.com/Pengxn/go-xn/src/middleware"
)

// InitRoutes initializes all routes.
func InitRoutes(ctx context.Context, c config.ServerConfig) error {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), middleware.RequestID(), middleware.Sentry())
	g.GET("/", controller.HomePage)

	// Page number handler
	g.GET("/page/:pageNum", controller.ArticlesPage)

	// Health check
	g.GET("/health", func(c *gin.Context) { c.String(200, "ok") })

	if c.Debug {
		debugRoutes(g)
	}
	apiRoutes(g)
	authRoutes(g)
	errorRoute(g)
	adminRoutes(g)
	othersRoutes(g)
	staticRoutes(g)
	articlesRoutes(g)

	if c.TLS {
		return g.RunTLS(":"+c.Port, c.CertFile, c.KeyFile)
	}
	return g.Run(":" + c.Port)
}
