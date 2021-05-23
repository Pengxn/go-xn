package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
	"github.com/Pengxn/go-xn/src/middleware"
)

// InitRoutes initializes all routes.
func InitRoutes(port string) error {
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()
	g.Use(middleware.Logger(), gin.Recovery())
	g.GET("/", controller.HomePage)

	// Page number handler
	g.GET("/page/:pageNum", controller.ArticlesPage)

	errorRoute(g)
	adminRoutes(g)
	othersRoutes(g)
	staticRoutes(g)
	optionsRoutes(g)
	articlesRoutes(g)

	return g.Run(":" + port)
}
