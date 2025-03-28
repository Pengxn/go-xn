package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/controller"
	"github.com/Pengxn/go-xn/src/middleware"
)

// InitRoutes initializes all routes.
func InitRoutes() error {
	serverConfig := config.Config.Server
	if !serverConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), middleware.RequestID(), middleware.Sentry())
	g.GET("/", controller.HomePage)

	// Page number handler
	g.GET("/page/:pageNum", controller.ArticlesPage)

	// Health check
	g.GET("/health", func(c *gin.Context) { c.String(200, "ok") })

	if serverConfig.Debug {
		debugRoutes(g)
	}
	apiRoutes(g)
	authRoutes(g)
	errorRoute(g)
	adminRoutes(g)
	othersRoutes(g)
	staticRoutes(g)
	articlesRoutes(g)

	if serverConfig.TLS {
		return g.RunTLS(":"+serverConfig.Port, serverConfig.CertFile, serverConfig.KeyFile)
	}
	return g.Run(":" + serverConfig.Port)
}
