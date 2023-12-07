package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// adminRoutes registers routes about admin.
func adminRoutes(g *gin.Engine) {
	admin := g.Group("/admin")
	admin.Any("/", controller.Redirect("/admin/login"))
	admin.GET("/login", controller.LoginPage)
	admin.POST("/login/begin", controller.BeginLogin)
	admin.POST("/login/finish", controller.FinishLogin)
	admin.GET("/register", controller.RegisterPage)
	admin.POST("/register/begin", controller.BeginRegister)
	admin.POST("/register/finish", controller.FinishRegister)
}
