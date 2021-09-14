package grpc

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/config"
	"dongo_game_server/src/database"
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// 通过一个结构体，实现proto中定义的所有服务
type RpcApp struct {
	DB     *database.DB
	Config *config.RpcConfig
}

func (p *RpcApp) Start() {
	listen, err := net.Listen("tcp", p.Config.UserServiceAddr) // Address gRPC服务地址
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`RpcApp listen error`)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到rpc服务器上，
	inf.RegisterUserServiceServer(s, &RpcUserService{DB: p.DB})
	log.Println("rpc serve running")
	if err := s.Serve(listen); err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`RpcApp Serve error`)
		log.Fatal(err)
	}
}
