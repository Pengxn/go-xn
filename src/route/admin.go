package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// adminRoutes registers routes about admin.
func adminRoutes(g *gin.Engine) {
	admin := g.Group("/admin")
	admin.Any("/", controller.Redirect("/admin/login"))
	admin.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Code": 200,
			"Data": "Hi",
		})
	})

	admin.GET("/register", controller.RegisterPage)
	admin.POST("/register/begin", controller.BeginRegister)
	admin.POST("/register/finish", controller.FinishRegister)
}
