package route

import (
	"github.com/gin-gonic/gin"
)

// errorRoute registers bad routes and methods.
func errorRoute(g *gin.Engine) {
	g.LoadHTMLFiles("web/error.html")

	// No route
	g.NoRoute(func(c *gin.Context) {
		if c.ContentType() == "text/html" {
			c.HTML(404, "error.html", gin.H{
				"code": 404,
				"data": "The Page Could Not be Found",
			})
		} else {
			c.JSON(404, gin.H{
				"code": 404,
				"data": "The Route Could Not be Found",
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
