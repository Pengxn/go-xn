package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// othersRoutes registers other routes.
func othersRoutes(g *gin.Engine) {
	// RSS 2.0, https://www.rssboard.org/rss-specification
	g.GET("/rss", controller.RSS)
	// Atom, https://datatracker.ietf.org/doc/html/rfc4287
	g.GET("/atom", controller.Atom)
	// JSON Feed Version 1, https://jsonfeed.org/version/1
	g.GET("/feed", controller.Feed)

	g.GET("/md", controller.Mdcat)

	// Get domain whois information.
	g.GET("/whois", controller.GwtWhoisInfo)

	// Register routes and methods for uPic, more information
	// to https://blog.svend.cc/upic/tutorials/custom
	g.POST("/upload/upic", controller.UploadFileForUPic)
	g.Static("/upic", "data/uPic")
}
