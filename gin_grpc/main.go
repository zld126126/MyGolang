package main

import (
	"log"

	"gin_grpc/pkg"
	"gin_grpc/util"
)

func main() {
	userService, err := pkg.InitUserService()
	if err != nil {
		log.Fatal(err)
		util.Catch(err)
	}
	go userService.CreateUserService()
	web, err := pkg.InitWeb()
	if err != nil {
		log.Fatal(err)
		util.Catch(err)
	}
	web.Start()
}
