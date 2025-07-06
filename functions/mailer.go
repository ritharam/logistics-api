package functions

import (
	"bytes"
	"html/template"
	// "log"
	"net/smtp"
	"os"
)

func SendMail(to, subject string, data map[string]string) error {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer
	err = tpl.Execute(&body, data)
	if err != nil {
		return err
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	sender := os.Getenv("SENDER")
	password := os.Getenv("APP_PASSWORD")

	auth := smtp.PlainAuth("", sender, password, smtpHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body.String())

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{to}, msg)
}