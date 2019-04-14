package route

import (
	"github.com/gin-gonic/gin"
)

// staticRoutes register static routes files and folders
func staticRoutes(g *gin.Engine) {
	g.StaticFile("/favicon.ico", "web/icons/xn-02f.png")

	g.Static("/asset", "web")
}
