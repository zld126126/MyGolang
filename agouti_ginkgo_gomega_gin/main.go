package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("../templates/*")
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", Index)
	router.POST("/login", Login)
	router.Run(":9090")
}

//http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": fmt.Sprintf("%v,%v", "DongTech", time.Now().Local()),
	})
}

type Param struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//http://localhost:9090/Login
func Login(c *gin.Context) {
	var p Param
	if err := c.Bind(&p); err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}
	if p.Username == "dongtech" && p.Password == "123456" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "登录成功",
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "登录失败",
	})
}
