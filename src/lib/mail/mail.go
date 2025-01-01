package mail

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"

	"github.com/go-mail/mail"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/util/log"
	"github.com/Pengxn/go-xn/web"
)

var dialer *mail.Dialer

func init() {
	smtp := config.Config.SMTP
	dialer = mail.NewDialer(smtp.Host, smtp.Port, smtp.Username, smtp.Password)
	if smtp.SkipTLS {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
}

// Mail is a struct for mail.
type Mail struct {
	from   string
	To     string
	ToName string
	Cc     []string
}

// Send sends a mail by smtp.
func Send(email Mail, mailType MailType) error {
	m := mail.NewMessage()
	if email.from == "" {
		email.from = m.FormatAddress(config.Config.SMTP.Username, "Feng.YJ")
	}
	m.SetHeader("From", email.from)
	if email.ToName != "" {
		email.To = m.FormatAddress(email.To, email.ToName)
	}
	m.SetHeader("To", email.To)
	if len(email.Cc) > 0 {
		m.SetHeader("Cc", email.Cc...)
	}
	subject, content := MailContent(mailType, "")
	if subject == "" || content == "" {
		return errors.New("mail subject or content is empty")
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	return dialer.DialAndSend(m)
}

// MailType is a type for mail type.
type MailType int

const (
	MailLogin  MailType = iota // MailLogin is a mail type for login.
	MailChange                 // MailChange is a mail type for change email.
	MailVerify                 // MailVerify is a mail type for verify email.
)

// MailContent is a struct for mail content.
func MailContent(mailType MailType, magicLink string) (subject, content string) {
	var (
		temp *template.Template
		err  error
	)
	switch mailType {
	case MailLogin:
		subject = "Login - "
		temp, err = template.New("mail").Parse(web.MailLogin)
	case MailChange:
		subject = "Confirm Your Email Change - "
		temp, err = template.New("mail").Parse(web.MailChange)
	case MailVerify:
		subject = "Verify Your Email - "
		temp, err = template.New("mail").Parse(web.MailVerify)
	}
	if err != nil {
		log.Errorf("Fail to parse mail template %d, error: %v", mailType, err)
		return "", content
	}

	site := map[string]string{
		"name":          "Feng",
		"logo_url":      "https://fengyj.cn/icons/avatar.png",
		"contact_email": "support@fengyj.cn",
	}

	var cbuf bytes.Buffer
	if err = temp.Execute(&cbuf, map[string]any{
		"site":       site,
		"magic_link": "https://magic-link.com/test?code=xxxx",
	}); err != nil {
		log.Errorf("Fail to execute mail template %d, error: %v", mailType, err)
		return "", content
	}

	subject += site["name"]
	content = cbuf.String()
	return
}
