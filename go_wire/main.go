package main

import (
	"awesomeProject/service"
	"awesomeProject/wire"
	"log"
)

func main() {
	//hand()
	auto()
}

//手动注入
func hand() {
	service.NewUserService(1, "dong").Impl.DoA()
}

//自动注入
func auto() {
	//复杂逻辑注入
	userService, err := wire.NewUserService(1, "dong")
	if err != nil {
		log.Fatal(err)
	}
	userService.Impl.DoB()

	//简单注入
	logService, err := wire.InitLogService("haha")
	if err != nil {
		log.Fatal(err)
	}
	logService.Write()

	//简单注入
	configService, err := wire.InitConfigService([]string{"china", "win", "hello world"}, "8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}
	configService.Get()
}
