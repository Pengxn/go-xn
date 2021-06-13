package web

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/util/log"
)

//go:embed templates
//go:embed error.html
var html embed.FS

func HTML() *template.Template {
	t, err := template.ParseFS(html, "*.html", "**/*.html", "**/**/*.html")
	if err != nil {
		log.Errorf("template.ParseFS error: %+v", err)
	}

	if !gin.IsDebugging() {
		for _, tmpl := range t.Templates() {
			log.Info(tmpl.Name())
		}
	}

	return t
}

//go:embed assets
var assets embed.FS

func FS() http.FileSystem {
	assetsDir, err := fs.Sub(assets, "assets")
	if err != nil {
		log.Errorf("fs.Sub error: %+v", err)
	}

	return http.FS(assetsDir)
}

//go:embed robots.txt
//go:embed icons
var other embed.FS

func OtherFS() http.FileSystem { return http.FS(other) }
