package drivers

import (
	"net/smtp"
	"strings"
)

type plainMailDriver struct {
	identity string
	username string
	password string
	host     string
	port     string
	from     string
}

func NewPlainMailDriver() MailDriver {
	return &plainMailDriver{
		identity: "",
		username: "iam.gt2rz@gmail.com",
		password: "duddcuuyzlwawqdn",
		host:     "smtp.gmail.com",
		port:     "587",
		from:     "iam.gt2rz@gmail.com",
	}
}

func (p *plainMailDriver) SendEmail(emailsTo []string, subject string, body string, attach []string) error {
	auth := smtp.PlainAuth(p.identity, p.username, p.password, p.host)

	msg := []byte("To:" + strings.Join(emailsTo, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	url := p.host + ":" + p.port // smtp.gmail.com:587

	err := smtp.SendMail(url, auth, p.from, emailsTo, msg)

	return err
}
