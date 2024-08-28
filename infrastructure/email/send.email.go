package email

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gmail "gopkg.in/gomail.v2"
)

type Email struct {
	User string
	Password string
}

func NewEmail(user, password string) *Email {
	return &Email{User: user, Password: password}
}

func (email *Email)SendVerificationEmail(to, subject, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	username := os.Getenv("USER_EMAIL")
	password := os.Getenv("EMAIL_PASS")
	log.Println(username , password)
	m := gmail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gmail.NewDialer("smtp.gmail.com", 587, username, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}