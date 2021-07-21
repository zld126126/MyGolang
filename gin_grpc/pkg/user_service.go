package pkg

import (
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gin_grpc/database"
	"gin_grpc/pkg/model"
	"gin_grpc/service/inf"
	"gin_grpc/util"
)

// 通过一个结构体，实现proto中定义的所有服务
type UserServiceHandler struct {
	Config *Config
	DB     *database.DB
}

type UserGrpcService struct {
	Config *Config
	DB     *database.DB
}

func (p UserGrpcService) GetUser(c context.Context, req *inf.UserReq) (*inf.UserResp, error) {
	var user model.User
	err := p.DB.Gorm.Table(`users u`).Where(`u.id = ?`, req.Id).Scan(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("id:%d,get user err", req.Id))
	}
	return &inf.UserResp{Name: user.Name, Time: util.ParseTimeToInt64()}, nil
}

func (p UserGrpcService) PushUser(context.Context, *inf.UserReq) (*empty.Empty, error) {
	return nil, nil
}

func (p *UserServiceHandler) CreateUserService() {
	listen, err := net.Listen("tcp", p.Config.UserService.Addr) // Address gRPC服务地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		util.Catch(err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	inf.RegisterUserServiceServer(s, UserGrpcService{DB: p.DB, Config: p.Config})
	log.Println("grpc serve running")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
		util.Catch(err)
	}
}
