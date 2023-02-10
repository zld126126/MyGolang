package main

import (
	"embed"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func main() {
	router := gin.Default()

	// asset加载html
	templates, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	// 配置模板
	router.SetHTMLTemplate(templates)
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	router.StaticFS("/static/", http.FS(f))

	rootGroup := router.Group("/")
	{
		rootGroup.GET("/", Index)
		rootGroup.GET("/index", Index)
		rootGroup.GET("/test/", Test)
		rootGroup.POST("/chatgpt", ChatGPT)
	}
	router.Run(":9090")
}

// 执行命令: go-assets-builder templates -o assets.go -p main
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		// 可以用.tmpl .html
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
