package util

import "github.com/sirupsen/logrus"

//捕获异常 error
func Chk(err error) {
	if err != nil {
		logrus.WithError(err).Println("chk error")
		panic(err)
	}
}
