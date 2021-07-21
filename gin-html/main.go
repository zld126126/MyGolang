package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// asset 加载 htmls
	templates, err := loadTemplate()
	if err != nil {
		panic(err)
	}

	// 配置模板
	// 	router.LoadHTMLGlob("resources/templates/*")
	router.SetHTMLTemplate(templates)

	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径

	router.StaticFS("/static", Dir)
	// 请求
	group := router.Group("/")
	{
		group.GET("/", Index)
		group.GET("/index", Index)
		group.GET("/hello", Hello)
		group.GET("/sayHello/:name", SayHello)
	}
	router.Run(":9090")
}

//http://localhost:9090/sayHello/dong
func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, "hello,"+name)
}

//http://localhost:9090/hello
func Hello(c *gin.Context) {
	c.HTML(http.StatusOK, "/templates/hello.html", gin.H{
		"Hello": fmt.Sprintf("%v	%v", "HelloWorld!", time.Now().Local()),
	})
}

//http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "/templates/index.html", gin.H{
		`WEBSITE_TITLE`: `東の博客`,
	})
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
