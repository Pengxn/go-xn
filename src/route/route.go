package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/controller"
	"github.com/Pengxn/go-xn/src/middleware"
)

// InitRoutes initializes all routes.
func InitRoutes(port string) error {
	g := gin.New()
	g.Use(middleware.Logger(), gin.Recovery(), middleware.Sentry())
	g.GET("/", controller.HomePage)

	// Page number handler
	g.GET("/page/:pageNum", controller.ArticlesPage)

	errorRoute(g)
	adminRoutes(g)
	othersRoutes(g)
	staticRoutes(g)
	optionsRoutes(g)
	articlesRoutes(g)

	serverConfig := config.GetServerConfig()
	if !serverConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	if serverConfig.TLS {
		return g.RunTLS(":"+serverConfig.Port, serverConfig.CertFile, serverConfig.KeyFile)
	}
	return g.Run(":" + port)
}
