package controller

import (
	"dongo_game_server/src/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocketHdl struct {
	Service *service.SocketService
}

// 获取Socket对应连接
func (p *SocketHdl) Create(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (p *SocketHdl) InitSocket() {
	p.Service.InitPort()
}
