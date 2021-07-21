package main

import (
	pb "dongtech_go/proto" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	rpc := pb.NewServeRouteClient(conn)
	reqBody1 := &pb.Id{Id: 1}
	res1, err := rpc.GetUser(context.Background(), reqBody1) //就像调用本地函数一样，通过serve1得到返回值
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from serve: ", res1.Name)

	reqBody2 := &pb.Name{Name: "HaHa"}
	res2, err := rpc.GetActivity(context.Background(), reqBody2) //就像调用本地函数一样，通过serve2得到返回值
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from serve: ", res2.Name)
}
