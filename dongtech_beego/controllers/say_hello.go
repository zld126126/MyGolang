package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type SayHelloController struct {
	beego.Controller
}

func (c *SayHelloController) Get() {
	name := c.GetString("name")
	var helloWorld struct {
		Time int64  `json:"time"`
		Name string `json:"name"`
	}
	helloWorld.Time = time.Now().UnixNano() / 1e6
	helloWorld.Name = name
	c.Data["json"] = helloWorld
	c.ServeJSON()
}
