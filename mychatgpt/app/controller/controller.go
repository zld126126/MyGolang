package controller

import (
	"github.com/gin-gonic/gin"
	"mychatgpt/app/service"
	"net/http"
)

type Controller struct {
	OpenAI *service.OpenAIService
}

// Test http://localhost:9090/test
// @Summary 测试api接口
// @Tags 测试
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /test [GET]
func (p *Controller) Test(c *gin.Context) {
	c.String(http.StatusOK, "mychatgpt-dongtech 运行正常")
}

// Index http://localhost:9090/index
// @Summary 主页html
// @Tags 业务
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /index [GET]
func (p *Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "/templates/index.html", gin.H{
		`WEBSITE_TITLE`: `ChatGPT提问程序`,
	})
}

// ChatGPTReq ChatGPT请求参数
type ChatGPTReq struct {
	Ask string `form:"ask" json:"ask"`
}

// ChatGPTResp ChatGPT返回参数
type ChatGPTResp struct {
	Ask    string `form:"ask" json:"ask"`       //问题
	Answer string `form:"answer" json:"answer"` //回答
}

// ChatGPT http://localhost:9090/chatgpt/haha
// @Summary 获取ChatGPT答案
// @Tags 业务
// @Param Param body ChatGPTReq true "ChatGPT请求参数"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} ChatGPTResp
// @Router /chatgpt [POST]
func (p *Controller) ChatGPT(c *gin.Context) {
	var req ChatGPTReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请输入正确的问题!",
		})
		return
	}

	if &req == nil || req.Ask == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请输入正确的问题!",
		})
		return
	}

	answer, err := p.OpenAI.GetTextResult(req.Ask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "获取ChatGPT结果失败!",
		})
	}

	result := &ChatGPTResp{
		Ask:    req.Ask,
		Answer: answer,
	}
	c.JSON(http.StatusOK, result)
}
