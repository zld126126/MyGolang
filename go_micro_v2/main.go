package main

import (
	micro "github.com/micro/go-micro/v2"

	user "go_micro_v2/proto/user"
	userService "go_micro_v2/service"
	web2 "go_micro_v2/web"

	"github.com/micro/go-micro/v2/web"
)

func main() {
	// TODO 自动注入
	var app App
	app.init()
	CheckError(app.Run())
}

type App struct {
	SrvService micro.Service
	WebService web.Service
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *App) Run() error {
	go func() {
		CheckError(p.SrvService.Run())
	}()
	return p.WebService.Run()
}

func (p *App) init() {
	// Create a new service
	p.SrvService = micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Address(":10010"),
	)
	// Register Service
	user.RegisterUserHandler(p.SrvService.Server(), new(userService.User))
	p.SrvService.Init()

	// Create a new web service
	p.WebService = web.NewService(
		web.Name("web.gin"),
		web.Version("latest"),
		web.Address(":10086"))
	// TODO 自动注入
	handler := web2.Handler{
		Srv: p.SrvService,
	}
	// Register Handler
	p.WebService.Handle("/", handler.Api())
	p.WebService.Init()
}
