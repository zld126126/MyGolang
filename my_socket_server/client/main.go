package main

import (
	"fmt"
	"time"
)

func main() {
	StartClient()

	fmt.Println("这是一个Go客户端，实现了Socket消息广播功能")

	// 防止主线程退出
	for {
		time.Sleep(1 * time.Second)
	}
}

func StartClient() {
	client := NewClient("127.0.0.1", 8888)
	client.Start()
}
