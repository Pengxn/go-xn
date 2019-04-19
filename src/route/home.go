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
		c.HTML(404, "404.html", gin.H{
			"code": 404,
			"data": "The Page Could Not be Found",
		})
	})

	// No Method
	g.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"data": "The Method Could Not be Found",
		})
	})
}
