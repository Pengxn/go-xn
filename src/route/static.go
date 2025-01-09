package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/web"
)

// staticRoutes registers static routes files and folders.
func staticRoutes(g *gin.Engine) {
	g.SetHTMLTemplate(web.HTML(gin.IsDebugging()))

	g.StaticFS("/assets", web.FS())

	otherFS := web.OtherFS()
	staticFileFromFS(g, "/favicon.ico", "icons/logo.ico", otherFS)
	// Robots.txt, refer to https://www.robotstxt.org/
	staticFileFromFS(g, "/robots.txt", "robots.txt", otherFS)
	// Humans.txt, refer to https://humanstxt.org/
	staticFileFromFS(g, "/humans.txt", "humans.txt", otherFS)
}

func staticFileFromFS(g *gin.Engine, relativePath, filepath string, fs http.FileSystem) {
	handler := func(c *gin.Context) {
		c.FileFromFS(filepath, fs)
	}
	g.GET(relativePath, handler)
	g.HEAD(relativePath, handler)
}
