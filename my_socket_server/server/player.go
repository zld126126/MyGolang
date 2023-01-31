package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Player struct {
	Name    string      // 昵称，默认与Addr相同
	Addr    string      // 地址
	Channel chan string // 消息管道
	conn    net.Conn    // 连接
	server  *Server     // 缓存Server的引用
}

func NewPlayer(conn net.Conn, server *Server) *Player {
	playerAddr := conn.RemoteAddr().String()

	player := &Player{
		Name:    playerAddr,
		Addr:    playerAddr,
		Channel: make(chan string),
		conn:    conn,
		server:  server,
	}

	// 启动协程，监听Channel管道消息
	go player.ListenMessage()

	return player
}

func (p *Player) Online() {
	// 玩家上线，将玩家加入到OnlineMap中，注意加锁操作
	p.server.ServerLock.Lock()
	p.server.OnlinePlayerMap[p.Name] = p
	p.server.ServerLock.Unlock()

	// 广播当前玩家上线消息
	p.server.BroadCast(p, "上线啦O(∩_∩)O")
	fmt.Println("player Online")
}

func (p *Player) Offline() {
	// 玩家下线，将玩家从OnlineMap中删除，注意加锁
	p.server.ServerLock.Lock()
	delete(p.server.OnlinePlayerMap, p.Name)
	p.server.ServerLock.Unlock()

	// 广播当前玩家下线消息
	p.server.BroadCast(p, "下线了o(╥﹏╥)o")
	fmt.Println("player Offline")
}

func (p *Player) DoMessage(buf []byte, len int) {
	msg := string(buf[:len])
	fmt.Println("DoMessage: ", msg)
	// 调用Server的BroadCast方法
	p.server.BroadCast(p, msg)
}

func (p *Player) ListenMessage() {
	for {
		msg := <-p.Channel
		fmt.Println("Send msg to client: ", msg, ", len: ", int16(len(msg)))
		bytes := bytes.NewBuffer([]byte{})
		// TODO 前两个字节写入消息长度 (前后端编程语言一致时不需要)
		binary.Write(bytes, binary.BigEndian, int16(len(msg)))
		// 写入消息数据
		binary.Write(bytes, binary.BigEndian, []byte(msg))
		// 发送消息给客户端
		p.conn.Write(bytes.Bytes())
	}
}
