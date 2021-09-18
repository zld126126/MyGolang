package controller

import (
	"dongo_game_server/src/global_const"
	"dongo_game_server/src/model"
	"dongo_game_server/src/web/base"
	"dongo_game_server/src/web/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zld126126/dongo_utils"
	"net/http"
	"strconv"
)

type ManagerPathHdl struct {
	Service *service.ManagerPathService
	Manager *service.ManagerService
}

type ManagerPathCreateForm struct {
	Name       string `form:"name" json:"name"`               // 路径名称
	OptionPath string `form:"option_path" json:"option_path"` // 路径
}

func (p *ManagerPathHdl) Create(c *gin.Context) {
	var form ManagerPathCreateForm
	err := c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	err = p.Service.Add(form.Name, form.OptionPath)
	if err != nil {
		c.String(http.StatusBadRequest, "创建失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

type ManagerPathListForm struct {
	OptionPath string `form:"option_path" json:"option_path"` // 路径
	PageSize   int    `form:"pageSize" json:"pageSize"`       // 多少条 10
	Page       int    `form:"page" json:"page"`               // 第几页 1
}

func (p *ManagerPathHdl) List(c *gin.Context) {
	var form ManagerPathListForm
	err := c.BindQuery(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	total, l, err := p.Service.List(form.OptionPath, form.Page, form.PageSize)
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

func (p *ManagerPathHdl) GetKey(id int64) string {
	return fmt.Sprintf(global_const.ManagerPathKey, id)
}

func (p *ManagerPathHdl) Mid(c *gin.Context) {
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

func (p *ManagerPathHdl) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.ManagerPath)

	c.JSON(http.StatusOK, &base.Response{Data: m})
}

type ManagerPathUpdateForm struct {
	Name       string `form:"name" json:"name"`               // 路径名称
	OptionPath string `form:"option_path" json:"option_path"` // 路径
}

func (p *ManagerPathHdl) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		c.Abort()
		return
	}

	var form ManagerPathUpdateForm
	err = c.Bind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	key := p.GetKey(id)
	m := dongo_utils.ParisMap_Get(key).(*model.Manager)

	err = p.Service.Update(m.Id, form.Name, form.OptionPath)
	if err != nil {
		c.String(http.StatusBadRequest, "更新失败")
		return
	}

	c.String(http.StatusOK, "ok")
}

func (p *ManagerPathHdl) PathList(c *gin.Context) {
	var result []string
	result = global_const.Paths

	token := c.Request.Header.Get(global_const.ManagerWebHeaderKey)
	if token != "" {
		path, err := p.Manager.GetPathByToken(token)
		if err != nil {
			c.String(http.StatusBadRequest, "查询错误")
			return
		}

		if len(path) > 0 {
			result = append(result, path...)
			c.JSON(http.StatusOK, result)
			return
		}
	}

	c.JSON(http.StatusOK, result)
	return
}
