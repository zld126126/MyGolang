package controller

import (
	"dongo_game_server/src/web/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectHdl struct {
	Service *service.ProjectService
}

type ProjectCreateForm struct {
	Name         string `form:"name" json:"name"`                   // 项目名称
	ResourcePath string `form:"resource_path" json:"resource_path"` // 资源
	RestApi      string `form:"rest_api" json:"rest_api"`           // 项目api
}

// 项目创建
// curl -X POST "127.0.0.1:9090/project/create" -d "name=dongbao&resource_path=/dongbao&rest_api=/dongbao" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ProjectHdl) Create(c *gin.Context) {
	var form ProjectCreateForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	p.Service.Add(form.Name, form.ResourcePath, form.RestApi)
	if err != nil {
		c.String(http.StatusBadRequest, "创建失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

// 项目维护
func (p *ProjectHdl) Update(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 项目删除 一个或多个
func (p *ProjectHdl) Del(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 项目列表
func (p *ProjectHdl) List(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
