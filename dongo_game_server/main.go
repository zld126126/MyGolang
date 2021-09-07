package main

import (
	"dongo_game_server/src/boot"
	"dongo_game_server/src/goLua"
	"log"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func Application() {
	// grpc init
	grpcApp, err := boot.InitGrpc()
	if err != nil {
		log.Fatal(err)
		dongo_utils.Chk(err)
	}

	// grpc start
	go grpcApp.Start()

	// support init
	supportApp, err := boot.InitSupport()
	if err != nil {
		log.Fatal(err)
		dongo_utils.Chk(err)
	}

	// support init
	go supportApp.Start()

	// web init
	webApp, err := boot.InitWeb()
	if err != nil {
		log.Fatal(err)
		dongo_utils.Chk(err)
	}

	// web start
	webApp.Start()
}

// TODO 扩展成cobra-viper/src/cmd 命令启动
func main() {
	Application()

	//Test()
}

// TODO 扩展成Testing.T
func Test() {
	goLua.LuaObject.Example_Go2Lua()
	goLua.LuaObject.Example_Lua2Go()
}
