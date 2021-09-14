package controller

import (
	"dongo_game_server/src/web/service"
	"github.com/gin-gonic/gin"
)

type ClientHdl struct {
	Service *service.ClientService
}

func (p *ClientHdl) HttpTrackCollect(c *gin.Context) {

}
