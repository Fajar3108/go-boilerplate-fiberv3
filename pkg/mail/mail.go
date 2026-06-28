package mail

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	host     string
	port     int
	user     string
	password string
	sender   string
}

func NewMailer(host string, port int, user string, password string, sender string) *Mailer {
	return &Mailer{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		sender:   sender,
	}
}

func (m *Mailer) SendMail(to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.sender)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	d := gomail.NewDialer(m.host, m.port, m.user, m.password)
	if err := d.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email to %s: %w", to, err)
	}

	return nil
}
