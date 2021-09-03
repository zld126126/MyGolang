package main

import (
	"dongo_game_server/src/boot"
	"dongo_game_server/src/util"
	"log"
)

// TODO 扩展成cobra-viper/src/cmd 命令启动
func main() {
	// grpc init
	grpcApp, err := boot.InitGrpc()
	if err != nil {
		log.Fatal(err)
		util.Chk(err)
	}

	// grpc start
	go grpcApp.Start()

	// support init
	supportApp, err := boot.InitSupport()
	if err != nil {
		log.Fatal(err)
		util.Chk(err)
	}

	// support init
	go supportApp.Start()

	// web init
	webApp, err := boot.InitWeb()
	if err != nil {
		log.Fatal(err)
		util.Chk(err)
	}

	// web start
	webApp.Start()
}
