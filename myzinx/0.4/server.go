package main

import (
	"fmt"
	"myzinx/0.4/zinx/ziface"
	"myzinx/0.4/zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	//创建一个server句柄
	s := znet.NewServer()

	//配置路由
	s.AddRouter(&PingRouter{})

	//开启服务
	s.Serve()
}
