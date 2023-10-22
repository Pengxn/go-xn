package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
	"github.com/Pengxn/go-xn/src/middleware"
)

// othersRoutes registers other routes.
func othersRoutes(g *gin.Engine) {
	// RSS 2.0, https://www.rssboard.org/rss-specification
	g.GET("/rss", controller.RSS)
	// Atom, https://datatracker.ietf.org/doc/html/rfc4287
	g.GET("/atom", controller.Atom)
	// JSON Feed Version 1, https://jsonfeed.org/version/1
	g.GET("/feed", controller.Feed)

	// Well-Known URIs specification, refer to:
	// https://datatracker.ietf.org/doc/html/rfc8615
	// https://www.iana.org/assignments/well-known-uris/well-known-uris.xhtml
	wuris := g.Group("/.well-known")
	// WebDAV well-known URIs, redirect to /dav.
	wuris.Any("/webdav", controller.Redirect("/dav"))
	wuris.Any("/caldav", controller.Redirect("/dav"))
	wuris.Any("/carddav", controller.Redirect("/dav"))

	// WebDAV, server for WebDAV service.
	webdav := g.Group("/dav").Use(middleware.BasicAuth())
	for _, v := range controller.WebdavMethods {
		webdav.Handle(v, "/*webdav", controller.WebDAV)
	}

	g.GET("/md", controller.Mdcat)

	// Get domain whois information.
	g.GET("/whois", controller.GwtWhoisInfo)

	// Register routes and methods for uPic, more information
	// to https://blog.svend.cc/upic/tutorials/custom
	g.POST("/upload/upic", controller.UploadFileForUPic)
	g.Static("/upic", "data/uPic")
}
