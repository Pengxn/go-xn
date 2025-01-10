package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/Pengxn/go-xn/src/util/log"
)

//go:embed templates
//go:embed assets
var EmbedFS embed.FS

// FS returns the `assets` file system.
func FS() http.FileSystem {
	assets, err := fs.Sub(EmbedFS, "assets")
	if err != nil {
		log.Errorf("FS fs.Sub error: %+v", err)
	}

	return http.FS(assets)
}

//go:embed robots.txt
//go:embed humans.txt
//go:embed icons
var other embed.FS

// OtherFS returns the `other` file system.
func OtherFS() http.FileSystem {
	return http.FS(other)
}

var (
	//go:embed templates/mail/mail-login.html
	MailLogin string
	//go:embed templates/mail/mail-change.html
	MailChange string
	//go:embed templates/mail/mail-verify.html
	MailVerify string
)
