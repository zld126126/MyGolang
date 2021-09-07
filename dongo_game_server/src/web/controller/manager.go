package controller

import (
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"dongo_game_server/src/web/base"
	"dongo_game_server/src/web/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zld126126/dongo_utils/dongo_utils"
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
// @Router /manager/create/ [post]
// curl -X POST "127.0.0.1:9090/manager/create" -d "name=dongbao&password=123456&tp=3"
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
// @Success 200 object base.Response
// @Router /manager/list/ [get]
// curl -X GET "http://127.0.0.1:9090/manager/list?name=dongbao&page=1&pageSize=10"
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

	res := &base.Response{
		Total:    total,
		Data:     l,
		Page:     form.Page,
		PageSize: form.PageSize,
	}

	c.JSON(http.StatusOK, res)
}

func (p *ManagerHdl) GetKey(id int) string {
	return fmt.Sprintf(global_const.ManagerKey, id)
}

func (p *ManagerHdl) Mid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

// @获取指定用户
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "managerId"
// @Success 200 {string} string	"ok"
// @Router /manager/{some_id} [get]
// curl -X GET "http://127.0.0.1:9090/manager/1"
func (p *ManagerHdl) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Manager)

	c.JSON(http.StatusOK, m)
}

type ManagerUpdateForm struct {
	Name     string            `form:"name" json:"name"`         // 用户名称
	Password string            `form:"password" json:"password"` // 用户密码
	Tp       model.ManagerType `form:"tp" json:"tp"`             // 用户类型
}

// curl -X POST "127.0.0.1:9090/manager/1/edit" -d "name=dongbao2&password=123456&tp=3"
func (p *ManagerHdl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

// curl -X POST "127.0.0.1:9090/manager/1/del" -d ""
func (p *ManagerHdl) Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

// 登陆
// curl -X POST 127.0.0.1:9090/manager/login -d "name=dongbao&password=123456"
func (p *ManagerHdl) Login(c *gin.Context) {
	var form ManagerLoginForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	err = p.Service.Login(form.Name, form.Password)
	if err != nil {
		c.String(http.StatusBadRequest, "登陆失败")
		return
	}

	claims := &JWTClaims{
		UserID:      1,
		Username:    form.Name,
		Password:    form.Password,
		FullName:    form.Name,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := getToken(claims)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, signedToken)
}

// 登出
func (p *ManagerHdl) Logout(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 验证
func (p *ManagerHdl) Verify(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 刷新令牌
func (p *ManagerHdl) Refresh(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
