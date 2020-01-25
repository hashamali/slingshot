package email

// EmailSender sends emails.
type EmailSender interface {
	SendEmail(from string, to string, body string) error
}
