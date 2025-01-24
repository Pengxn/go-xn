package route

import (
	"github.com/gin-gonic/gin"
)

func authRoutes(g *gin.Engine) {
	// OAuth2
	g.POST("/oauth/:provider", oauth2)
	g.POST("/oauth/callback", oauth2)
	g.POST("/oauth/logout", oauth2)
}

func oauth2(c *gin.Context) {
	c.String(200, "not implemented")
}
