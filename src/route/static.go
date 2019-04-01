package route

import (
	"github.com/gin-gonic/gin"
)

// StaticRoutes set static files and folders
func StaticRoutes(g *gin.Engine) {
	g.Static("/asset", "web")
}
