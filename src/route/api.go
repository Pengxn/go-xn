package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/api"
	"github.com/Pengxn/go-xn/src/controller"
	"github.com/Pengxn/go-xn/src/middleware"
)

func apiRoutes(g *gin.Engine) {
	adminAPI := api.NewAdminAPI()

	// API without JWT authentication
	g.GET("/api/token", adminAPI.Token)

	api := g.Group("/api").Use(middleware.JWT())
	api.POST("/admin", adminAPI.RegisterAdmin)
	{ // option
		api.GET("/options", controller.ListOptions)
		api.GET("/option/:name", controller.GetOption)
		api.POST("/option", controller.InsertOption)
		api.PUT("/option/:name", controller.UpdateOption)
		api.DELETE("/option/:name", controller.DeleteOption)
	}
}
