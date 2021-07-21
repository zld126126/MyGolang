package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.Data["content"] = "未经授权，请求要求验证身份"
	c.TplName = "error/error.tpl"
}
func (c *ErrorController) Error403() {
	c.Data["content"] = "服务器拒绝请求"
	c.TplName = "error/error.tpl"
}
func (c *ErrorController) Error404() {
	c.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	c.TplName = "error/error.tpl"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "server error"
	c.TplName = "error/error.tpl"
}
func (c *ErrorController) Error503() {
	c.Data["content"] = "服务器目前无法使用（由于超载或停机维护）"
	c.TplName = "error/error.tpl"
}

func Error(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("dberror.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/dberror.html")
	data := make(map[string]interface{})
	data["content"] = "server is now down"
	t.Execute(rw, data)
}
