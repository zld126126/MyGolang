package controller

import (
	"dongo_game_server/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceHdl struct {
	DB *database.DB
}

// 获取静态资源
func (p *ResourceHdl) Get(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
