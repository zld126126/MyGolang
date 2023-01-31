package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线玩家容器
	OnlinePlayerMap map[string]*Player

	// 容器锁，对容器进行操作时进行加锁
	ServerLock sync.RWMutex

	// 消息广播的管道
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:              ip,
		Port:            port,
		OnlinePlayerMap: make(map[string]*Player),
		Message:         make(chan string),
	}

	return server
}

// Start 启动服务
func (p *Server) Start() {
	// socket监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", p.Ip, p.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// 程序退出时，关闭监听，注意defer关键字的用途
	defer listener.Close()

	// 启动一个协程来执行ListenMessage
	go p.ListenMessage()

	// 注意for循环不加条件，相当于while循环
	for {
		// Accept，此处会阻塞，当有客户端连接时才会往后执行
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// 启动一个协程去处理
		go p.Handler(conn)
	}
}

func (p *Server) Handler(conn net.Conn) {
	// 构造player对象，NewPlayer全局方法在player.go脚本中
	player := NewPlayer(conn, p)

	// 玩家上线
	player.Online()

	// 启动一个协程
	go func() {
		buf := make([]byte, 4096)
		for {
			// 从Conn中读取消息
			len, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			if len == 0 {
				// 玩家下线
				player.Offline()
				return
			}

			// 玩家针对msg进行消息处理
			player.DoMessage(buf, len)
		}
	}()
}

func (p *Server) BroadCast(player *Player, msg string) {
	//sendMsg := "[" + player.Addr + "]: " + msg
	//p.Message <- sendMsg
	p.Message <- player.Addr + ":" + msg
}

func (p *Server) ListenMessage() {
	for {
		// 从Message管道中读取消息
		msg := <-p.Message

		// 加锁
		p.ServerLock.Lock()
		// 遍历在线玩家，把广播消息同步给在线玩家
		for _, player := range p.OnlinePlayerMap {
			// 把要广播的消息写到玩家管道中
			player.Channel <- msg
		}
		// 解锁
		p.ServerLock.Unlock()
	}
}
