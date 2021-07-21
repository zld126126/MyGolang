package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	user "go_micro_v2/proto/user"
)

// 该模块演示的是在常规的web服务中调用微服务，比如在gin，beego，echo搭建的服务中调用微服务
type Handler struct {
	Srv micro.Service
}

//
func (p *Handler) Api() *gin.Engine {
	router := gin.Default()
	router.GET("/greeter", p.Anything)
	router.GET("/greeter/:name", p.Hello)
	return router
}

func (p *Handler) Anything(c *gin.Context) {
	logger.Info("Received Say.Anything API request")
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (p *Handler) Hello(c *gin.Context) {
	logger.Info("Received Say.Hello API request")

	name := c.Param("name")

	client := user.NewUserService("go.micro.srv.user", p.Srv.Client())
	response, err := client.QueryUserByName(context.TODO(), &user.Request{
		UserName: name,
	})

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, response)
}
