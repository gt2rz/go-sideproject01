package mailables

type ResetPasswordEmailData struct {
	Token string
}

// Mailable for the reset password email
type ResetPasswordEmail struct {
	TemplateName string
	TemplateFile string
	Email        string
	Data         ResetPasswordEmailData
	Subject      string
	Attach       []string
}

// NewResetPasswordEmail returns a new instance of the reset password email
func NewResetPasswordEmail(email string, token string) *ResetPasswordEmail {
	return &ResetPasswordEmail{
		TemplateName: "resetTokenPassword.html",
		TemplateFile: "./services/mail/views/resetTokenPassword.html",
		Email:        email,
		Data: ResetPasswordEmailData{
			Token: token,
		},
		Subject: "Reset Password",
	}
}
