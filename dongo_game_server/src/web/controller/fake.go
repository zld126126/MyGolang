package controller

import (
	"dongo_game_server/src/global_const"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FakeHdl struct{}

// curl -X POST "127.0.0.1:9090/debug/manager/create" -d "name=dongbao&password=123456&tp=3" -H "Fake-Id: YWRtaW4="
func (p *FakeHdl) Mid(c *gin.Context) {
	token := c.Request.Header.Get(global_const.FakeIdDebugKey)
	if token == "" {
		c.String(http.StatusBadRequest, "token非法")
		c.Abort()
		return
	}

	if token != global_const.FakeIdAdmin {
		c.String(http.StatusBadRequest, "token非法")
		c.Abort()
		return
	}
}
