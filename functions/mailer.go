package functions

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

func SendMail(to, subject string, data map[string]string) error {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer
	tpl.Execute(&body, data)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	sender := os.Getenv("SENDER")
	password := os.Getenv("APP_PASSWORD")

	auth := smtp.PlainAuth("", sender, password, smtpHost)
	msg := []byte("Subject: " + subject + "\r\n" + "\r\n" + body.String())

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{to}, msg)
}