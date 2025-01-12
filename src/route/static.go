package route

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/util/log"
	"github.com/Pengxn/go-xn/web"
)

// staticRoutes registers static routes files and folders.
func staticRoutes(g *gin.Engine) {
	g.SetHTMLTemplate(templateFromFS(web.EmbedFS))

	g.StaticFS("/assets", assetsFS())

	otherFS := http.FS(web.Other)
	// logo and favicon
	staticFileFromFS(g, "/favicon.ico", "icons/logo.ico", otherFS)
	staticFileFromFS(g, "/logo.svg", "icons/xn-02f.svg", otherFS)
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

// HTML returns the HTML template.
func templateFromFS(fsys fs.FS) *template.Template {
	templates, err := fs.Sub(fsys, "templates")
	if err != nil {
		log.Errorf("HTML fs.Sub error: %+v", err)
	}

	t, err := template.ParseFS(templates, "*.html", "**/*.html")
	if err != nil {
		log.Errorf("HTML template.ParseFS error: %+v", err)
	}

	if gin.IsDebugging() {
		for _, tmpl := range t.Templates() {
			fmt.Println(tmpl.Name())
		}
	}

	return t
}

// assetsFS returns the assets file system.
func assetsFS() http.FileSystem {
	assets, err := fs.Sub(web.EmbedFS, "assets")
	if err != nil {
		log.Errorf("FS fs.Sub error: %+v", err)
	}

	return http.FS(assets)
}
