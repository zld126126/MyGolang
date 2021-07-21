package controllers

import "github.com/astaxie/beego"

type InfoController struct {
	beego.Controller
}

func (c *InfoController) Get() {
	c.Data["Email"] = "liandong.zhang@outlook.com"
	c.Data["Name"] = "dongtech"
	c.TplName = "user.tpl"
}
