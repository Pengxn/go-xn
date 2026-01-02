package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

func authRoutes(g *gin.Engine) {
	// OAuth2
	g.GET("/oauth/:provider", controller.OAuth2Redirect)
	g.GET("/oauth/callback", controller.OAuth2Callback)
	g.POST("/oauth/logout", oauth2)
}

func oauth2(c *gin.Context) {
	c.String(200, "not implemented")
}
