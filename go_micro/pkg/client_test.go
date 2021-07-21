package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/micro/go-micro"

	proto "go_micro/proto"
)

func TestMain(m *testing.M) {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("landon.assist.client"), micro.Address(":9091"))
	service.Init()

	// 创建客户端
	assist := proto.NewQueryService("landon.assist.service", service.Client())

	// 调用greeter服务
	rsp, err := assist.GetUser(context.TODO(), &proto.UserId{Id: 888888})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp)
}
