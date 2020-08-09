package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// optionsRoutes register routes about option.
func optionsRoutes(g *gin.Engine) {
	g.Any("/options", controller.ListOptions)

	option := g.Group("/option")
	option.GET("/:name", controller.GetOption)
	option.POST("/", controller.InsertOption)
	option.PUT("/:name", controller.UpdateOption)
	option.DELETE("/:name", controller.DeleteOption)
}
