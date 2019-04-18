package route

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/controller"
)

// InitRoutes will init all routes.
func InitRoutes(g *gin.Engine) {
	g.LoadHTMLGlob("web/template/*")

	g.GET("/", controller.Home)

	staticRoutes(g)
	articlesRoutes(g)
}
