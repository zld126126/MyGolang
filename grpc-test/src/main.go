package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-test/proto"
	"log"
	"net/http"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"
const dateTimeFormat = "20060102150405"

func main() {
	router := gin.Default()
	router.GET("/test", Test)
	router.GET("/grpc", Grpc)
	router.Run(":9090")
}

func Test(c *gin.Context) {
	c.String(http.StatusOK, "HelloWorld!")
}

func Grpc(c *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())
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
	c.String(http.StatusOK, "get grpc success,serve1:"+timeFormat(res1.Time, dateFormat))
}

func timeFormat(timeStamp int64, format ...string) string {
	t := time.Unix(timeStamp/1e3, 0)
	defaultFormat := dateTimeFormat
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	return t.Format(defaultFormat)
}
