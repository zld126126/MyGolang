package web

import (
	"fmt"
	"log"
	"net/http"

	"my_wx_app/config"
	"my_wx_app/database"

	"github.com/gin-gonic/gin"
	"github.com/zld126126/dongo_utils/dongo_utils"
)

type WebApp struct {
	Config *config.Config
	DB     *database.DB
}

func (p *WebApp) Start() {
	router := gin.New()
	// 替换用法 router.Use(gin.Recovery())
	router.Use(ServeRecover)
	routerGroup := router.Group("")
	p.Mount(routerGroup)

	err := router.Run(fmt.Sprintf(`:%s`, p.Config.WebAddr))
	if err != nil {
		log.Fatalln(err)
		dongo_utils.Chk(err)
	}

	log.Println("web serve running")
}

func hi(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (p *WebApp) Mount(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/hi", hi)
}
