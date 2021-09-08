package main

import (
	"log"
	"my_wx_app/boot"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func main() {
	// web init
	webApp, err := boot.InitWeb()
	if err != nil {
		log.Fatal(err)
		dongo_utils.Chk(err)
	}

	// web start
	webApp.Start()
}
