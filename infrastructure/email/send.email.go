package email

import (
	"fmt"

	gmail "gopkg.in/gomail.v2"
)

type Email struct {
	User string
	Password string
}

func (email *Email)SendVerificationEmail(to, subject, body string) error {
	m := gmail.NewMessage()
	m.SetHeader("From", email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gmail.NewDialer("smtp.gmail.com", 587, email.User, email.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}