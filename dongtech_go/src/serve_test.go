package main

import (
	pb "dongtech_go/proto" // 引入编译生成的包
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

//通过一个结构体，实现proto中定义的所有服务
type ServeRoute struct{}

func (h ServeRoute) GetUser(ctx context.Context, in *pb.Id) (*pb.User, error) {
	resp := &pb.User{
		Time: time.Now().UnixNano() / 1e6,
		Name: fmt.Sprintf("%d,dongtech", in.Id),
	}
	fmt.Println(resp)
	return resp, nil
}

func (h ServeRoute) GetActivity(ctx context.Context, in *pb.Name) (*pb.Activity, error) {
	tp := pb.Tp(rand.Intn(4))
	resp := &pb.Activity{
		Name: in.Name,
		Tp:   tp,
	}
	fmt.Println(resp)
	return resp, nil
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10086") // Address gRPC服务地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	pb.RegisterServeRouteServer(s, ServeRoute{})
	log.Println("grpc serve running")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
