package main

import (
	"fmt"
	"net"
)

const socketTestPort = "20200"

// curl -X POST "127.0.0.1:9090/project/create" -d "name=dongbao&resource_path=/dongbao&rest_api=/dongbao"
func SocketTest() {
	conn, err := net.Dial("tcp", ":"+socketTestPort) //连接服务端
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connect to localhost:" + socketTestPort + " success")
	defer conn.Close()
	for {
		//一直循环读入用户数据，发送到服务端处理
		fmt.Print("Please input send data :")
		var a string
		fmt.Scan(&a)
		if a == "exit" {
			break
		} //添加一个退出机制，用户输入exit，退出
		_, err := conn.Write([]byte(a))
		if err != nil {
			fmt.Println(err)
			return
		}
		data := make([]byte, 2048)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Response data :", string(data[:n]))
	}
}
