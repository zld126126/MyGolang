package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

type Client struct {
	Ip   string
	Port int

	// 连接
	conn net.Conn

	// 容器锁，对容器进行操作时进行加锁
	ClientLock sync.RWMutex

	//发送消息
	Message chan string
}

func NewClient(ip string, port int) *Client {
	return &Client{
		Ip:      ip,
		Port:    port,
		Message: make(chan string),
	}
}

func (p *Client) Start() {
	address := p.Ip + ":" + fmt.Sprint(p.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	go p.Receive()
	go p.Send()
	p.conn = conn
	defer conn.Close()
	for {
		fmt.Println("请输入信息，回车结束输入")
		reader := bufio.NewReader(os.Stdin)
		//终端读取用户回车，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		line = strings.TrimSpace(line)
		p.Message <- line
	}
}

func (p *Client) Receive() {
	buf := make([]byte, 4096)
	for {
		len, err := p.conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Conn Read err:", err)
			return
		}

		if len == 0 {
			return
		}

		data := buf[:len]
		fmt.Printf("收到的内容:%s\n", string(data))
	}
}

func (p *Client) Send() {
	for {
		// 从Message管道中读取消息
		msg := <-p.Message

		// 加锁
		p.ClientLock.Lock()

		bytes := bytes.NewBuffer([]byte{})
		// TODO 前两个字节写入消息长度 (前后端编程语言一致时不需要)
		binary.Write(bytes, binary.BigEndian, int16(len(msg)))
		// 写入消息数据
		binary.Write(bytes, binary.BigEndian, []byte(msg))
		// 发送消息给客户端
		p.conn.Write(bytes.Bytes())
		fmt.Printf("发送的内容:%s\n", msg)

		// 解锁
		p.ClientLock.Unlock()
	}
}
