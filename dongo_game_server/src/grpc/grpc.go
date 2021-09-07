package grpc

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/config"
	"dongo_game_server/src/database"
	"log"
	"net"

	"github.com/zld126126/dongo_utils/dongo_utils"
	"google.golang.org/grpc"
)

// 通过一个结构体，实现proto中定义的所有服务
type GrpcApp struct {
	DB              *database.DB
	GrpcUserService *config.GrpcConfig
}

func (p *GrpcApp) Start() {
	listen, err := net.Listen("tcp", p.GrpcUserService.UserServiceAddr) // Address gRPC服务地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		dongo_utils.Chk(err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	inf.RegisterUserServiceServer(s, &Grpc_UserService{DB: p.DB})
	log.Println("grpc serve running")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
		dongo_utils.Chk(err)
	}
}
