package email

import (
	"fmt"
	"net/smtp"

	"github.com/hashamali/slingshot/config"
)

// SMTPEmail implements Email over SMTP.
type SMTPEmail struct {
	Config config.SMTPConfig
}

// GetURL will get the SMTP URL.
func (email SMTPEmail) GetURL() string {
	return fmt.Sprintf("%s:%v", email.Config.Host, email.Config.Port)
}

// SendEmail sends emails.
func (email SMTPEmail) SendEmail(from string, to []string, body []byte) error {
	return smtp.SendMail(
		email.GetURL(),
		smtp.PlainAuth("", from, email.Config.Password, email.Config.Host),
		from,
		to,
		body,
	)
}
