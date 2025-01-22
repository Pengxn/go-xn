package route

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"

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

	localFS := http.FS(os.DirFS("./data/.well-known"))
	// Security.txt, refer to: https://securitytxt.org/
	// and https://www.iana.org/assignments/security-txt-fields/security-txt-fields.xhtml
	// and https://datatracker.ietf.org/doc/html/rfc9116
	staticFile(g, "/.well-known/security.txt", "security.txt", localFS, otherFS)
	// keybase.txt, refer to: https://keybase.io/docs/keybase_well_known
	staticFile(g, "/.well-known/keybase.txt", "keybase.txt", localFS, otherFS)
}

func staticFile(g *gin.Engine, relativePath, filepath string, localfs, embedfs http.FileSystem) {
	handler := func(c *gin.Context) {
		if _, err := localfs.Open(filepath); err == nil {
			c.FileFromFS(filepath, localfs)
		} else {
			c.FileFromFS(filepath, embedfs)
		}
	}
	g.GET(relativePath, handler)
	g.HEAD(relativePath, handler)
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
