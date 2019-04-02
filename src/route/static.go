package route

import (
	"github.com/gin-gonic/gin"
)

// StaticRoutes set static routes files and folders
func StaticRoutes(g *gin.Engine) {
	g.StaticFile("/favicon.ico", "web/icons/xn-02f.png")

	g.Static("/asset", "web")
}
