package controller

import (
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"dongo_game_server/src/web/base"
	"dongo_game_server/src/web/service"
	"fmt"
	"github.com/zld126126/dongo_utils/dongo_utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHdl struct {
	Service *service.ProjectService
}

type ProjectCreateForm struct {
	Name         string `form:"name" json:"name"`                   // 项目名称
	ResourcePath string `form:"resource_path" json:"resource_path"` // 资源
	RestApi      string `form:"rest_api" json:"rest_api"`           // 项目api
	Port         int64  `form:"port"' json:"port"`                  // 端口号
}

// @Summary 项目创建
// @Tags 管理项目
// @Description 项目创建
// @Accept  json
// @Produce  json
// @Param   name     query    string     true        "项目名称"
// @Param   resource_path     query    string     true        "资源"
// @Param   rest_api      query    string     true        "项目api"
// @Param   port      query    int     true        "端口号"
// @Success 200 {string} string	"ok"
// @Router /web/project/create/ [post]
// curl -X POST "127.0.0.1:9090/project/create" -d "name=dongbao&resource_path=/dongbao&rest_api=/dongbao&port=10086" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ProjectHdl) Create(c *gin.Context) {
	var form ProjectCreateForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	err = p.Service.Add(form.Name, form.ResourcePath, form.RestApi, form.Port)
	if err != nil {
		c.String(http.StatusBadRequest, "创建失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

type ProjectUpdateForm struct {
	Name         string `form:"name" json:"name"`                   // 项目名称
	ResourcePath string `form:"resource_path" json:"resource_path"` // 资源
	RestApi      string `form:"rest_api" json:"rest_api"`           // 项目api
	Port         int64  `form:"port"' json:"port"`                  // 端口号
}

// @Summary 编辑项目信息
// @Tags 管理项目
// @Description 编辑项目信息
// @Accept  json
// @Produce json
// @Param   name     query    string     true        "项目名称"
// @Param   resource_path     query    string     true        "资源"
// @Param   rest_api      query    string     true        "项目api"
// @Param   port      query    int     true        "端口号"
// @Param   id path int true "项目id"
// @Success 200 {string} string	"ok"
// @Router /web/project/{id}/edit [post]
func (p *ProjectHdl) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	var form ProjectUpdateForm
	err = c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Project)

	err = p.Service.Update(m.Id, form.Name, form.ResourcePath, form.RestApi, form.Port)
	if err != nil {
		c.String(http.StatusBadRequest, "更新失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

// @Summary 获取项目信息
// @Tags 管理项目
// @Description 获取项目信息
// @Accept  json
// @Produce json
// @Param   id path int true "项目id"
// @Success 200 object base.Response
// @Router /web/project/{id} [get]
func (p *ProjectHdl) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Project)

	c.JSON(http.StatusOK, &base.Response{Data: m})
}

// 项目维护
func (p *ProjectHdl) Mid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	m, err := p.Service.Get(id)
	if err != nil {
		c.String(http.StatusBadRequest, "不存在")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	dongo_utils.ParisMap_Put(key, m)
}

func (p *ProjectHdl) GetKey(id int64) string {
	return fmt.Sprintf(global_const.ProjectKey, id)
}

// @Summary 删除项目
// @Tags 管理项目
// @Description 删除项目
// @Accept  json
// @Produce json
// @Param   id path int true "项目id"
// @Success 200 {string} string	"ok"
// @Router /web/project/{id}/del [post]
func (p *ProjectHdl) Del(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Project)

	err = p.Service.Del(m.Id)
	if err != nil {
		c.String(http.StatusBadRequest, "删除失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

type ProjectListForm struct {
	Name     string `form:"name" json:"name"`         // 项目名称
	PageSize int    `form:"pageSize" json:"pageSize"` // 多少条 10
	Page     int    `form:"page" json:"page"`         // 第几页 1
}

// @Summary 获取所有项目
// @Tags 管理项目
// @Description 获取所有项目
// @Accept  json
// @Produce  json
// @Param   name     query    string     true        "项目名"
// @Param   pageSize     query    int     true        "条数"
// @Param   page      query    int     true        "页数"
// @Success 200 object base.ListResponse
// @Router /web/project/list/ [get]
func (p *ProjectHdl) List(c *gin.Context) {
	var form ProjectListForm
	err := c.BindQuery(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	total, l, err := p.Service.List(form.Name, form.Page, form.PageSize)
	if err != nil {
		c.String(http.StatusBadRequest, "查询出错")
		return
	}

	res := &base.ListResponse{
		Total:    total,
		Data:     l,
		Page:     form.Page,
		PageSize: form.PageSize,
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 刷新项目token
// @Tags 管理项目
// @Description 刷新项目token
// @Accept  json
// @Produce json
// @Param   id path int true "项目id"
// @Success 200 {string} string	"ok"
// @Router /web/project/{id}/refreshToken [post]
func (p *ProjectHdl) RefreshToken(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Project)

	err = p.Service.RefreshToken(m.Id)
	if err != nil {
		c.String(http.StatusBadRequest, "刷新token失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

// TODO 重启socket服务
func (p *ProjectHdl) RestartSocket(c *gin.Context) {

}
