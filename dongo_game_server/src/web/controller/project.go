package controller

import (
	"dongo_game_server/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectHdl struct {
	DB *database.DB
}

// 项目创建
func (p *ProjectHdl) Create(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 项目维护
func (p *ProjectHdl) Update(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 项目删除 一个或多个
func (p *ProjectHdl) Del(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 项目列表
func (p *ProjectHdl) List(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
