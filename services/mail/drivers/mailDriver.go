package drivers

type MailDriver interface {
	SendEmail(emailsTo []string, subject string, body string, attach []string) error
}
