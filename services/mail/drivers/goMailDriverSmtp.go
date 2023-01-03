package drivers

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type goMailDriver struct {
	username string
	password string
	host     string
	port     string
	from     string
}

func NewGoMailDriver() MailDriver {
	return &goMailDriver{
		username: os.Getenv("MAIL_USERNAME"),
		password: os.Getenv("MAIL_PASSWORD"),
		host:     os.Getenv("MAIL_HOST"),
		port:     os.Getenv("MAIL_PORT"),
		from:     os.Getenv("MAIL_FROM"),
	}
}

func (driver *goMailDriver) SendEmail(emailsTo []string, subject string, body string, attach []string) error {

	// Create a new message.
	message := gomail.NewMessage()

	// Set E-Mail sender parameters.
	message.SetHeader("From", driver.from)
	message.SetHeader("Subject", subject)

	for _, email := range emailsTo {
		message.SetHeader("To", email)
	}

	message.SetBody("text/html", body)

	for _, a := range attach {
		message.Attach(string(a))
	}

	port, _ := strconv.Atoi(driver.port)

	// Settings for SMTP server.
	dialer := gomail.NewDialer(driver.host, port, driver.username, driver.password)

	// Now send E-Mail.
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
