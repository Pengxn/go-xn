package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// InitRoutes initializes all routes.
func InitRoutes(port string) error {
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()
	g.GET("/", controller.HomePage)

	// Page number handler
	g.GET("/page/:pageNum", controller.ArticlesPage)

	errorRoute(g)
	othersRoutes(g)
	staticRoutes(g)
	optionsRoutes(g)
	articlesRoutes(g)

	return g.Run(":" + port)
}
