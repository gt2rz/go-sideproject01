package mail

import (
	"bytes"
	"html/template"
	"microtwo/services/mail/drivers"
	"microtwo/services/mail/mailables"
)

type EmailService struct {
	Mailer drivers.MailDriver
}

func NewEmailService(driver drivers.MailDriver) (*EmailService, error) {
	return &EmailService{
		Mailer: driver,
	}, nil
}

func (e *EmailService) SendResetPasswordEmail(email *mailables.ResetPasswordEmail) error {

	t := template.New(email.TemplateName)

	t, err := t.ParseFiles(email.TemplateFile)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, email.Data); err != nil {
		return err
	}

	result := tpl.String()

	var emailsTo []string
	emailsTo = append(emailsTo, email.Email)

	return e.Mailer.SendEmail(emailsTo, email.Subject, result, email.Attach)
}
