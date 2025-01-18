package route

import (
	"github.com/Pengxn/go-xn/src/api"
	"github.com/gin-gonic/gin"
)

func apiRoutes(g *gin.Engine) {
	adminAPI := api.NewAdminAPI()

	api := g.Group("/api")
	api.POST("/admin", adminAPI.RegisterAdmin)
}
