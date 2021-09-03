package controller

import (
	"dongo_game_server/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocketHdl struct {
	DB *database.DB
}

// 获取Socket对应连接
func (p *SocketHdl) Create(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
