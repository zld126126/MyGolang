package handler

import (
	"dongtech_go/config"
	"dongtech_go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
	"net/http"
)

type EmailMsg struct {
	From       []string
	To         []string
	Cc         []string
	Subject    string
	Body       string
	Attachment string
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func SendEmail(c *gin.Context) {
	var email EmailMsg
	email.From = []string{"zld126126@sina.com"}
	//10分钟邮箱
	email.To = []string{"QWYHIN@10min.club"}
	email.Subject = fmt.Sprintf("%s,%d", "test", util.ParseSecondTimeToInt64())
	email.Body = fmt.Sprintf("this is a test letter from dongbao.")

	config, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Println("get config err")
		util.Catch(err)
	}

	email.Host = config.Email.Host
	email.Port = config.Email.Port
	email.Username = config.Email.Username
	email.Password = config.Email.Password
	email.From = []string{config.Email.Username}

	err = sendEmailDefault(email)
	if err != nil {
		logrus.WithError(err).Println("send email err")
		c.String(http.StatusBadRequest, "send email err")
		return
	}
	c.String(http.StatusOK, "send email ok")
}

func sendEmailDefault(msg EmailMsg) error {
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
	err := gomail.NewDialer(msg.Host, msg.Port, msg.Username, msg.Password).DialAndSend(m)
	if err != nil {
		logrus.WithError(err).WithField("email", msg).Println("send email err")
		return err
	}
	return nil
}
