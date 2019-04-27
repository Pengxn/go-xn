package route

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/controller"
)

// InitRoutes will init all routes.
func InitRoutes(g *gin.Engine) {
	g.GET("/", controller.HomePage)
	g.GET("/about", controller.AboutPage)

	// Page number handler
	g.GET("/page/:pageNum", controller.ByPageNum)

	errorRoute(g)
	staticRoutes(g)
	optionsRoutes(g)
	articlesRoutes(g)
}
