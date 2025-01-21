package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/api"
	"github.com/Pengxn/go-xn/src/middleware"
)

func apiRoutes(g *gin.Engine) {
	adminAPI := api.NewAdminAPI()

	// API without JWT authentication
	g.GET("/api/token", adminAPI.Token)

	api := g.Group("/api").Use(middleware.JWT())
	api.POST("/admin", adminAPI.RegisterAdmin)
}
