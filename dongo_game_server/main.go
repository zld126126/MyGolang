package main

import (
	"dongo_game_server/src/boot"
	"dongo_game_server/src/goLua"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

func Application() {
	// GoRpc init
	rpcApp, err := boot.InitRpc()
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`rpcApp init error`)
		log.Fatal(err)
	}

	// GoRpc start
	go rpcApp.Start()

	// support init
	supportApp, err := boot.InitSupport()
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`supportApp init error`)
		log.Fatal(err)
	}

	// support init
	go supportApp.Start()

	// web init
	webApp, err := boot.InitWeb()
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`webApp init error`)
		log.Fatal(err)
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
