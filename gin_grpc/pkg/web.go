package pkg

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gin_grpc/database"
	"gin_grpc/pkg/controller"
	"gin_grpc/service/inf"
	"gin_grpc/util"
)

type WebHandler struct {
	Config      *Config
	UserService inf.UserServiceClient
	DB          *database.DB
}

func (p *WebHandler) Start() {
	router := gin.New()
	router.Use(ServeRecover)
	routerGroup := router.Group("")
	p.Mount(routerGroup)
	err := router.Run(fmt.Sprintf(`:%s`, p.Config.Web.Addr))
	if err != nil {
		log.Fatalln(err)
		util.Catch(err)
	}
	log.Println("web serve running")
}

func (p *WebHandler) Mount(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/sayHello", controller.SayHello)
	routerGroup.GET("/grpc/user/:user_id", controller.GetGrpcUser(p.UserService))
	routerGroup.GET("/user/:id", controller.GetUser(p.DB.Gorm))

	captcha := routerGroup.Group("/captcha")
	{
		captcha.GET("", controller.GetCaptcha)
		captcha.GET("/img", controller.GetCaptchaImg)
		captcha.POST("/verify", controller.VerifyCaptcha)
	}

	routerGroup.GET("/email", controller.SendEmail(p.Config.Email))

}
