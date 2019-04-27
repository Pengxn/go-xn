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
	articlesRoutes(g)
}

func errorRoute(g *gin.Engine) {
	g.LoadHTMLFiles("web/error.html")

	// No route
	g.NoRoute(func(c *gin.Context) {
		if c.ContentType() == "application/json" {
			c.JSON(404, gin.H{
				"code": 404,
				"data": "The JSON Could Not be Found",
			})
		} else {
			c.HTML(404, "error.html", gin.H{
				"code": 404,
				"data": "The Page Could Not be Found",
			})
		}
	})

	// No Method
	g.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"code": 405,
			"data": "The Method Could Not be Found",
		})
	})
}
