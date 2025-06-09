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
	// JSON Feed Version 1.1, https://www.jsonfeed.org/version/1.1
	g.GET("/feed.json", controller.Feed)

	// Well-Known URIs specification, refer to:
	// https://datatracker.ietf.org/doc/html/rfc8615
	// https://www.iana.org/assignments/well-known-uris/well-known-uris.xhtml
	wuris := g.Group("/.well-known")
	// WebDAV well-known URIs, redirect to /dav.
	wuris.Match(controller.WebdavMethods, "/webdav", controller.Redirect("/dav"))
	wuris.Match(controller.WebdavMethods, "/caldav", controller.Redirect("/dav"))
	wuris.Match(controller.WebdavMethods, "/carddav", controller.Redirect("/dav"))

	// WebAuthn well-known URIs, note: its status is still proposal, refer to:
	// https://github.com/w3c/webauthn/wiki/Explainer:-Related-origin-requests#proposal
	wuris.Any("/webauthn", controller.WebAuthnWellKnown)

	// WebDAV, server for WebDAV service.
	webdav := g.Group("/dav").Use(middleware.BasicAuth())
	webdav.Match(controller.WebdavMethods, "/*webdav", controller.WebDAV)

	// Implement MetaWeblog XML-RPC API, refer to:
	// https://codex.wordpress.org/XML-RPC_MetaWeblog_API
	g.Any("/metaweblog", controller.MetaWeblog)

	g.GET("/md", controller.Mdcat)

	// Get domain whois information.
	g.GET("/whois", controller.GwtWhoisInfo)

	// Register routes and methods for uPic, more information
	// to https://blog.svend.cc/upic/tutorials/custom
	g.POST("/upload/upic", controller.UploadFileForUPic)
	g.Static("/upic", "data/uPic")

	// Bitcoin BIP15 aliases, the status of standard BIP15 aliases is deferred, refer to:
	// https://github.com/bitcoin/bips/blob/master/bip-0015.mediawiki#https-web-service
	g.Any("/bitcoin-alias", controller.BitcoinAliases)

	// Metrics for Prometheus.
	g.GET("/metrics", controller.PrometheusMetrics)
}
