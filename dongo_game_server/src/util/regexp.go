package util

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

func RegexpToken(token string) (bool, error) {
	logrus.WithField("token", token).WithField("length", len(token)).Println("======")
	if matched, err := regexp.MatchString(`^[A-Za-z0-9/.]{243}$`, token); err != nil {
		return false, err
	} else {
		return matched, nil
	}
}
