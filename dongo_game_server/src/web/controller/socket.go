package controller

import (
	"bytes"
	"dongo_game_server/src/web/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type SocketHdl struct {
	Service *service.SocketService
	Project *service.ProjectService
}

type SocketCreateForm struct {
	ProjectId int64 `form:"project_id" json:"project_id"` // 项目id
}

func (p *SocketHdl) HandleSocket(conn net.Conn) {
	defer conn.Close() //关闭连接
	logrus.Println("Connect :", conn.RemoteAddr())
	var dataBuffer bytes.Buffer
	for {
		//只要客户端没有断开连接，一直保持连接，读取数据
		data := make([]byte, 2048)
		n, err := conn.Read(data)
		//数据长度为0表示客户端连接已经断开
		if n == 0 {
			logrus.Printf("%s has disconnect", conn.RemoteAddr())
			break
		}
		if err != nil {
			logrus.Println(err)
			continue
		}
		logrus.Printf("Receive data [%s] from [%s]", string(data[:n]), conn.RemoteAddr())
		//转大写
		rspData := string(data[:n])

		_, err = dataBuffer.Write([]byte(rspData))
		if err != nil {
			logrus.Println(err)
			continue
		}
	}
	fmt.Printf(dataBuffer.String())
	// 重置buffer
	if dataBuffer.Len() != 0 {
		dataBuffer.Reset()
	}
}

func (p *SocketHdl) AcceptSocket(port int64) {
	listener, err := net.Listen("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		logrus.Println(err)
		return
	}
	logrus.Println("Start listen localhost" + fmt.Sprint(port))

	for {
		//开始循环接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//一旦收到客户端连接，开启一个新的gorutine去处理这个连接
		go p.HandleSocket(conn)
	}
}

// 获取Socket对应连接
// curl -X POST "127.0.0.1:9090/socket" -d "project_id=1"
func (p *SocketHdl) Create(c *gin.Context) {
	var form SocketCreateForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	proj, err := p.Project.Get(form.ProjectId)
	if err != nil {
		c.String(http.StatusBadRequest, "项目不存在")
		return
	}

	usePort, err := p.Service.GetInUsePort(form.ProjectId)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}
	if usePort != 0 {
		c.JSON(http.StatusOK, gin.H{"port": usePort})
		return
	}

	port, err := p.Service.GetAvailablePort(form.ProjectId)
	if err != nil {
		c.String(http.StatusBadRequest, "无可用Socket端口")
		return
	}

	err = p.Service.UsePort(port, proj.Id)
	if err != nil {
		c.String(http.StatusBadRequest, "创建socket失败")
		return
	}

	p.AcceptSocket(port)

	c.JSON(http.StatusOK, gin.H{"port": port})
}

// 初始化socket
func (p *SocketHdl) InitSocket() {
	p.Service.InitPort()
}
