package app

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mychatgpt/app/controller"
	_ "mychatgpt/docs"
	"net/http"
)

//go:embed templates/*
var f embed.FS

type App struct {
	Server     *gin.Engine
	Controller *controller.Controller
}

func (p *App) Start() {
	// 服务容灾recover
	p.Server.Use(serveRecover)

	// asset加载html
	templates, err := LoadTemplate()
	if err != nil {
		panic(err)
	}
	// 配置模板
	p.Server.SetHTMLTemplate(templates)
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	p.Server.StaticFS("/static/", http.FS(f))

	rootGroup := p.Server.Group("/")
	{
		rootGroup.GET("/", p.Controller.Index)
		rootGroup.GET("/index", p.Controller.Index)
		rootGroup.GET("/test/", p.Controller.Test)
		rootGroup.POST("/chatgpt", p.Controller.ChatGPT)

		//执行命令 swag init 初始化swagger
		//http://localhost:9090/swagger/index.html
		rootGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	logrus.Println("My ChatGPT 启动成功")
	p.Server.Run(":9090")
}
