package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	proto "go_micro/proto"
)

func InterceptController(service proto.QueryService, db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(ServeRecover)
	group := router.Group("")
	{
		group.GET("/hello", SayHello)
		group.GET("/activity", GetActivity(service))
		group.GET("/user/:id", GetUser(db))
		group.GET("/err", TestErr)
	}
	return router
}

// http://localhost:9090/hello
func SayHello(c *gin.Context) {
	c.String(http.StatusOK, "hello,world!")
}

// http://localhost:9090/activity
func GetActivity(service proto.QueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		rsp, err := service.GetUser(context.TODO(), &proto.UserId{Id: 666})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, rsp)
	}
}

// http://localhost:9090/err
func TestErr(c *gin.Context) {
	panic(errors.New("我故意的error"))
	c.String(http.StatusBadRequest, "")
}
