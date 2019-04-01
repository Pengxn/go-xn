package route

import (
	"github.com/gin-gonic/gin"
)

// InitStatic set static files and folders
func InitStatic(g *gin.Engine) {
	g.Static("/asset", "web")
}
