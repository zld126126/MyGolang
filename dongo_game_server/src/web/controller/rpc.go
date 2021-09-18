package controller

import (
	"dongo_game_server/service/inf"
	"fmt"
	"github.com/zld126126/dongo_utils"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type RpcHdl struct {
	UserService inf.UserServiceClient
}

// http://localhost:9090/client/rpc/user/2
func (p *RpcHdl) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}
	userResp, err := p.UserService.GetUser(context.Background(), &inf.UserReq{Id: int32(userId)})
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`RpcHdl GetUser err`)
		dongo_utils.Chk(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user_name": userResp.Name,
	})
}
