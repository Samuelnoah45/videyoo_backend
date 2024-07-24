package utilService

import (
	"bytes"
	"net/smtp"
	"os"
)

func SendEmail(email string, body bytes.Buffer) (string, error) {
	from := os.Getenv("MOVIE_SERVER_GMAIL")
	to := email

	//1. google
	smtpServer := os.Getenv("SMTP_SERVER")
	auth := smtp.PlainAuth(os.Getenv("MOVIE_SERVER_NAME"), os.Getenv("MOVIE_SERVER_GMAIL"), os.Getenv("MOVIE_SERVER_PASSWORD"), smtpServer)
	err := smtp.SendMail(os.Getenv("SMTP_AUTH_SERVER"), auth, from, []string{to}, body.Bytes())

	//2. mailtrap service
	// auth := smtp.PlainAuth("", "bb0fbe593f233b", "2f47796776dd86", "sandbox.smtp.mailtrap.io")
	// err := smtp.SendMail("sandbox.smtp.mailtrap.io:2525", auth, from, []string{to}, body.Bytes())

	return "Email sent", err
}
