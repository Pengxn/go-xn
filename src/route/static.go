package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/web"
)

// staticRoutes registers static routes files and folders.
func staticRoutes(g *gin.Engine) {
	g.SetHTMLTemplate(web.HTML())

	g.StaticFS("/assets", web.FS())

	otherFS := web.OtherFS()
	staticFileFromFS(g, "/favicon.ico", "icons/logo.ico", otherFS)
	staticFileFromFS(g, "/robots.txt", "robots.txt", otherFS)
}

func staticFileFromFS(g *gin.Engine, relativePath, filepath string, fs http.FileSystem) {
	handler := func(c *gin.Context) {
		c.FileFromFS(filepath, fs)
	}
	g.GET(relativePath, handler)
	g.HEAD(relativePath, handler)
}
