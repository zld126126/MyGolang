package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// http://localhost:9090/test
func Test(c *gin.Context) {
	c.String(http.StatusOK, "mychatgpt-dongtech 运行正常")
}

// http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "/templates/index.html", gin.H{
		`WEBSITE_TITLE`: `ChatGPT提问程序`,
	})
}

type ChatGPTParam struct {
	Ask string `form:"ask" json:"ask"`
}

// http://localhost:9090/chatgpt/where is china?
func ChatGPT(c *gin.Context) {
	var param ChatGPTParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请输入正确的问题!",
		})
		return
	}

	if &param == nil || param.Ask == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请输入正确的问题!",
		})
		return
	}

	openAI := GetOpenAIService()
	resp, err := openAI.GetTextResult(param.Ask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "获取ChatGPT结果失败!",
		})
	}
	c.JSON(http.StatusOK, gin.H{`Ask`: param.Ask, `Answer`: resp})
}
