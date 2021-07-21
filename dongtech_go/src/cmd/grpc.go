package cmd

import (
	"context"
	pb "dongtech_go/proto"
	"dongtech_go/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func Grpc(addr string) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalln(err)
		}
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
		c.String(http.StatusOK, "get grpc success,serve1:"+timeFormat(res1.Time, util.DateFormat))
	}
}

func timeFormat(timeStamp int64, format ...string) string {
	t := time.Unix(timeStamp/1e3, 0)
	defaultFormat := util.DateFormat
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	return t.Format(defaultFormat)
}
