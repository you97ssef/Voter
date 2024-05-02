package services

import (
	"fmt"
	"net/smtp"
	"strings"
)

func NewMailer(username, password, host, port string) *Mailer {
	return &Mailer{
		username: username,
		password: password,
		host:     host,
		port:     port,
	}
}

type Mailer struct {
	username string
	password string
	host     string
	port     string
}

func (m *Mailer) SendEmail(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	err := smtp.SendMail(m.host+":"+m.port, auth, m.username, to, m.buildMessage(to, subject, body))
	if err != nil {
		return err
	}

	return nil
}

func (m *Mailer) buildMessage(to []string, subject, body string) []byte {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", m.username)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	return []byte(msg)
}
