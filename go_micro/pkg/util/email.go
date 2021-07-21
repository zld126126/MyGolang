package util

import (
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
)

type EmailMsg struct {
	From       []string
	To         []string
	Cc         []string
	Subject    string
	Body       string
	Attachment string
}

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func SendEmail(msg *EmailMsg, emailConfig *EmailConfig) error {
	m := gomail.NewMessage()
	if len(msg.From) > 0 {
		m.SetHeader("From", msg.From...)
	}
	if len(msg.To) > 0 {
		m.SetHeader("To", msg.To...)
	}
	if len(msg.Cc) > 0 {
		for _, cc := range msg.Cc {
			m.SetAddressHeader("Cc", cc, "")
		}
	}
	if msg.Subject != "" {
		m.SetHeader("Subject", msg.Subject)
	}
	if msg.Body != "" {
		m.SetBody("text/html", msg.Body)
	}
	if msg.Attachment != "" {
		m.Attach(msg.Attachment)
	}

	// Send the email to Bob, Cora and Dan.
	err := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.Username, emailConfig.Password).DialAndSend(m)
	if err != nil {
		logrus.WithError(err).WithField("email", msg).Println("send email err")
		return err
	}
	return nil
}
