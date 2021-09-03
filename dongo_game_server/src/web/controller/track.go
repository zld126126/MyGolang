package controller

import (
	"dongo_game_server/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrackHdl struct {
	DB *database.DB
}

// 采集打点信息 TODO : http/proto/socket多平台实现
func (p *TrackHdl) Collect(c *gin.Context) {
	// go func() 多goruntine
	c.String(http.StatusOK, "ok")
}

// 采集数据列表
func (p *TrackHdl) List(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 采集数据统计
func (p *TrackHdl) Stat(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
