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
//go:embed assets
var embedFS embed.FS

func HTML() *template.Template {
	templates, err := fs.Sub(embedFS, "templates")
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

func FS() http.FileSystem {
	assets, err := fs.Sub(embedFS, "assets")
	if err != nil {
		log.Errorf("FS fs.Sub error: %+v", err)
	}

	return http.FS(assets)
}

//go:embed robots.txt
//go:embed icons
var other embed.FS

func OtherFS() http.FileSystem { return http.FS(other) }
