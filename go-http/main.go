package main

import (
	"fmt"
	"net/http"
)

type Controller struct{}

func (p *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("server"))
}

// 自封装-http请求
func main() {
	addr := "127.0.0.1" + ":" + fmt.Sprint(8080)
	handler := InitHandler(addr)
	handler.Mount(map[string]http.Handler{
		"/":    &Controller{},
		"/api": &Controller{},
	})

	// 测试api 浏览器|curl:
	// curl http://127.0.0.1:8080/
	// curl http://127.0.0.1:8080/api
	// 启动:
	handler.Run()
}

