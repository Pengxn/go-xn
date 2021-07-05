package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/util/log"
)

//go:embed templates
var html embed.FS

func HTML() *template.Template {
	templatesDir, err := fs.Sub(html, "templates")
	if err != nil {
		log.Errorf("HTML fs.Sub error: %+v", err)
	}

	t, err := template.ParseFS(templatesDir, "*.html", "**/*.html")
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

//go:embed assets
var assets embed.FS

func FS() http.FileSystem {
	assetsDir, err := fs.Sub(assets, "assets")
	if err != nil {
		log.Errorf("FS fs.Sub error: %+v", err)
	}

	return http.FS(assetsDir)
}

//go:embed robots.txt
//go:embed icons
var other embed.FS

func OtherFS() http.FileSystem { return http.FS(other) }
