package web

import (
	"embed"
)

var (
	//go:embed templates
	//go:embed assets
	EmbedFS embed.FS

	//go:embed robots.txt
	//go:embed humans.txt
	//go:embed security.txt
	//go:embed icons
	Other embed.FS

	//go:embed templates/mail/mail-login.html
	MailLogin string
	//go:embed templates/mail/mail-change.html
	MailChange string
	//go:embed templates/mail/mail-verify.html
	MailVerify string
)
