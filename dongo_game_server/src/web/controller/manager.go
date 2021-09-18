package controller

import (
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"dongo_game_server/src/web/base"
	"dongo_game_server/src/web/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zld126126/dongo_utils"
)

type ManagerHdl struct {
	Service *service.ManagerService
}

type ManagerCreateForm struct {
	Name     string            `form:"name" json:"name"`         // 用户名称
	Password string            `form:"password" json:"password"` // 用户密码
	Tp       model.ManagerType `form:"tp" json:"tp"`             // 用户类型
}

// @Summary 创建管理用户
// @Tags 管理用户
// @Description 创建管理用户
// @Accept  json
// @Produce  json
// @Param   name     query    string     true        "用户名"
// @Param   password     query    string     true        "密码"
// @Param   tp      query    int     true        "用户类型"
// @Success 200 {string} string	"ok"
// @Router /web/manager/create/ [post]
// curl -X POST "127.0.0.1:9090/web/manager/create" -d "name=dongbao&password=123456&tp=3" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
// curl -X POST "127.0.0.1:9090/debug/manager/create" -d "name=dongbao&password=123456&tp=3"
func (p *ManagerHdl) Create(c *gin.Context) {
	var form ManagerCreateForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	err = p.Service.Add(form.Name, form.Password, form.Tp)
	if err != nil {
		c.String(http.StatusBadRequest, "创建失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

type ManagerListForm struct {
	Name     string `form:"name" json:"name"`         // 用户名称
	PageSize int    `form:"pageSize" json:"pageSize"` // 多少条 10
	Page     int    `form:"page" json:"page"`         // 第几页 1
}

// @Summary 获取所有管理用户
// @Tags 管理用户
// @Description 获取所有管理用户
// @Accept  json
// @Produce  json
// @Param   name     query    string     true        "用户名"
// @Param   pageSize     query    int     true        "条数"
// @Param   page      query    int     true        "页数"
// @Success 200 object base.ListResponse
// @Router /web/manager/list/ [get]
// curl -X GET "http://127.0.0.1:9090/web/manager/list?name=dongbao&page=1&pageSize=10" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ManagerHdl) List(c *gin.Context) {
	var form ManagerListForm
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

func (p *ManagerHdl) GetKey(id int64) string {
	return fmt.Sprintf(global_const.ManagerKey, id)
}

func (p *ManagerHdl) Mid(c *gin.Context) {
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

// @Summary 获取管理员信息
// @Tags 管理用户
// @Description 获取管理员信息
// @Accept  json
// @Produce json
// @Param   id path int true "管理员id"
// @Success 200 object base.Response
// @Router /web/manager/{id} [get]
// curl -X GET "http://127.0.0.1:9090/web/manager/1" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ManagerHdl) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Manager)

	c.JSON(http.StatusOK, &base.Response{Data: m})
}

type ManagerUpdateForm struct {
	Name     string            `form:"name" json:"name"`         // 用户名称
	Password string            `form:"password" json:"password"` // 用户密码
	Tp       model.ManagerType `form:"tp" json:"tp"`             // 用户类型
}

// @Summary 编辑管理员信息
// @Tags 管理用户
// @Description 编辑管理员信息
// @Accept  json
// @Produce json
// @Param   name     query    string     true        "用户名"
// @Param   password     query    string     true        "密码"
// @Param   tp      query    int     true        "用户类型"
// @Param   id path int true "管理员id"
// @Success 200 {string} string	"ok"
// @Router /web/manager/{id}/edit [post]
// curl -X POST "127.0.0.1:9090/web/manager/1/edit" -d "name=dongbao2&password=123456&tp=3" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ManagerHdl) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	var form ManagerUpdateForm
	err = c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Manager)

	err = p.Service.Update(m.Id, form.Name, form.Password, form.Tp)
	if err != nil {
		c.String(http.StatusBadRequest, "更新失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

// @Summary 删除管理员
// @Tags 管理用户
// @Description 删除管理员
// @Accept  json
// @Produce json
// @Param   id path int true "管理员id"
// @Success 200 {string} string	"ok"
// @Router /web/manager/{id}/del [post]
// curl -X POST "127.0.0.1:9090/web/manager/1/del" -d "" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ManagerHdl) Del(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Manager)

	err = p.Service.Del(m.Id)
	if err != nil {
		c.String(http.StatusBadRequest, "删除失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

type ManagerLoginForm struct {
	Name     string `form:"name" json:"name"`         // 用户名称
	Password string `form:"password" json:"password"` // 用户密码
}

// @Summary 登陆用户
// @Tags 管理用户
// @Description 登陆用户
// @Accept  json
// @Produce json
// @Param   name     query    string     true        "用户名"
// @Param   password     query    string     true        "密码"
// @Success 200 {string} string	"token:XXXXXXXX"
// @Router /web/manager/{id}/edit [post]
// curl -X POST 127.0.0.1:9090/web/manager/login -d "name=dongbao&password=123456" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
// curl -X POST 127.0.0.1:9090/base/manager/login -d "name=dongbao&password=123456"
func (p *ManagerHdl) Login(c *gin.Context) {
	var form ManagerLoginForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	u, err := p.Service.Login(form.Name, form.Password)
	if err != nil {
		c.String(http.StatusBadRequest, "登陆失败")
		return
	}

	token, err := p.Service.EncodeToken(u)
	if err != nil {
		c.String(http.StatusBadRequest, "登陆失败")
		return
	}

	c.String(http.StatusOK, token)
}

// @Summary 登出用户
// @Tags 管理用户
// @Description 登出用户
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /web/manager/{id}/edit [post]
// curl -X POST 127.0.0.1:9090/web/manager/logout -d "" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
// curl -X POST 127.0.0.1:9090/base/manager/logout -d ""
func (p *ManagerHdl) Logout(c *gin.Context) {
	// TODO 暂时前端简单处理 返回index页面
	c.String(http.StatusOK, "ok")
}

// @Summary 刷新令牌
// @Tags 管理用户
// @Description 刷新令牌
// @Accept  json
// @Produce json
// @Param   name     query    string     true        "用户名"
// @Param   password     query    string     true        "密码"
// @Success 200 {string} string	"token:XXXXXXXX"
// @Router /web/manager/{id}/edit [post]
// curl -X POST 127.0.0.1:9090/web/manager/login -d "name=dongbao&password=123456" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ManagerHdl) Refresh(c *gin.Context) {
	u := GetCurrentManager(c.Request, p.Service)
	if u == nil {
		c.String(http.StatusBadRequest, "token无效")
		return
	}

	token, err := p.Service.EncodeToken(u)
	if err != nil {
		c.String(http.StatusBadRequest, "登陆失败")
		return
	}

	c.String(http.StatusOK, token)
}

// context登陆校验用户
// curl -X GET "http://127.0.0.1:9090/web/manager/list?name=dongbao&page=1&pageSize=10" -H "ManagerWebHeaderKey: MXx8MTYzMTYxMjUxMTA0MQ=="
func (p *ManagerHdl) MidLogin(c *gin.Context) {
	token := c.Request.Header.Get(global_const.ManagerWebHeaderKey)
	if token == "" {
		c.String(http.StatusBadRequest, "token非法")
		c.Abort()
		return
	}

	_, err := p.Service.DecodeToken(token)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
}

func (p *ManagerHdl) PathList(c *gin.Context) {

}

func (p *ManagerHdl) PathBind(c *gin.Context) {

}

func (p *ManagerHdl) PathUnBind(c *gin.Context) {

}

// context登陆校验用户
// TODO fake-id情况处理
func (p *ManagerHdl) IsSuperManager(c *gin.Context) {
	token := c.Request.Header.Get(global_const.ManagerWebHeaderKey)
	if token == "" {
		c.String(http.StatusBadRequest, "token非法")
		c.Abort()
		return
	}

	m, err := p.Service.DecodeToken(token)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	if !m.IsSuper() {
		c.String(http.StatusBadRequest, "超管权限")
		c.Abort()
		return
	}
}

// context获取登陆用户
func GetCurrentManager(h *http.Request, service *service.ManagerService) *model.Manager {
	token := h.Header.Get(global_const.ManagerWebHeaderKey)
	if token != "" {
		m, err := service.DecodeToken(token)
		if err == nil {
			return m
		}
	}
	return nil
}
