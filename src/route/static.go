package route

import (
	"github.com/gin-gonic/gin"
)

// staticRoutes registers static routes files and folders.
func staticRoutes(g *gin.Engine) {
	g.LoadHTMLFiles("web/error.html")

	g.StaticFile("/favicon.ico", "web/icons/xn-02f.png")
	g.StaticFile("/robots.txt", "web/robots.txt")

	g.Static("/css", "web/css")
	g.Static("/js", "web/js")
	g.Static("/image", "web/image")
}
