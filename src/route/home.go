package route

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/controller"
)

// InitRoutes will init all routes.
func InitRoutes(g *gin.Engine) {
	g.GET("/", controller.HomePage)

	// Page number handler
	Page := g.Group("/page")
	Page.GET("/:pageNum", controller.ByPageNum)

	staticRoutes(g)
	articlesRoutes(g)

	g.LoadHTMLFiles("web/404.html")

	// No route
	g.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	// No Method
	g.NoMethod(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
}
