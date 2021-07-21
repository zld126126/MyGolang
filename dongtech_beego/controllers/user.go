package controllers

import (
	"dongtech/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Redirect("/Error", 400)
		return
	}

	o := orm.NewOrm()
	u := models.User{Id: id}
	err = o.Read(&u)
	if err != nil {
		c.Ctx.WriteString("db error")
		return
	}
	c.Data["json"] = u
	c.ServeJSON()
}
