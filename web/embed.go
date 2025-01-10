package web

import (
	"embed"
	"net/http"
)

//go:embed templates
//go:embed assets
var EmbedFS embed.FS

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
