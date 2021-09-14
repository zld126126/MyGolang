package controller

import (
	"dongo_game_server/src/model"
	"dongo_game_server/src/web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClientHdl struct {
	Service *service.ClientService
}

type TrackCollectForm struct {
	OpenId   string           `form:"openId" json:"open_id"`    // openId
	SourceTp model.SourceType `form:"tp" json:"source_tp"`      // 来源类型 0未知 1微信小程序 2安卓 3IOS 4WEB 5其他
	Token    string           `form:"project_id" json:"token"`  // 项目token
	Messages []string         `form:"messages" json:"messages"` // 打点信息 1个或多个
}

// @Summary 采集数据
// @Tags 埋点
// @Description 采集数据
// @Accept  json
// @Produce  json
// @Param   open_id     query    string     true        "openId"
// @Param   source_tp     query    string     true        "来源类型 0未知 1微信小程序 2安卓 3IOS 4WEB 5其他"
// @Param   token      query    int     true        "项目token"
// @Success 200 {string} string	"ok"
// @Router /web/manager/create/ [post]
// curl -X POST "127.0.0.1:9090/web/manager/create" -d "name=dongbao&password=123456&tp=3" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
// curl -X POST "127.0.0.1:9090/debug/manager/create" -d "name=dongbao&password=123456&tp=3"
func (p *ClientHdl) HttpTrackCollect(c *gin.Context) {
	var form TrackCollectForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	go p.Service.CollectByHttp(form.OpenId, form.SourceTp, form.Token, form.Messages)

	c.String(http.StatusOK, "ok")
}
