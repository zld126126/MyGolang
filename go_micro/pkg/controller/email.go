package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"go_micro/pkg/util"
)

func SendEmail(emailConfig *util.EmailConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var email util.EmailMsg
		email.From = []string{"zld126126@sina.com"}
		// 10分钟邮箱
		email.To = []string{"QWYHIN@10min.club"}
		email.Subject = fmt.Sprintf("%s,%d", "test", util.ParseSecondTimeToInt64())
		email.Body = fmt.Sprintf("this is a test letter from dongbao.")
		email.From = []string{emailConfig.Username}

		err := util.SendEmail(&email, emailConfig)
		if err != nil {
			logrus.WithError(err).Println("send email err")
			c.String(http.StatusBadRequest, "send email err")
			return
		}
		c.String(http.StatusOK, "send email ok")
	}
}
