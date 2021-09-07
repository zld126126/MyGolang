package web

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/config"
	"dongo_game_server/src/util"
	"dongo_game_server/src/web/controller"
	"fmt"
	"log"

	_ "dongo_game_server/src/web/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebApp struct {
	Config      *config.Config
	UserService inf.UserServiceClient

	Base     *controller.BaseHdl
	Captcha  *controller.CaptchaHdl
	JWT      *controller.JWTHdl
	Manager  *controller.ManagerHdl
	Project  *controller.ProjectHdl
	Resource *controller.ResourceHdl
	RPC      *controller.RpcHdl
	Socket   *controller.SocketHdl
	Tool     *controller.ToolHdl
	Track    *controller.TrackHdl
}

func (p *WebApp) Start() {
	p.Socket.InitSocket()

	router := gin.New()
	router.Use(ServeRecover)
	// router.LoadHTMLGlob("./resources")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerGroup := router.Group("")
	p.Mount(routerGroup)

	err := router.Run(fmt.Sprintf(`:%s`, p.Config.Web.Addr))
	if err != nil {
		log.Fatalln(err)
		util.Chk(err)
	}

	log.Println("web serve running")
}

func (p *WebApp) Mount(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/version", p.Base.GetVersion())
	routerGroup.GET("/grpc/user/:user_id", p.RPC.GetGrpcUser)

	manager := routerGroup.Group("/manager")
	{
		manager.POST("/create", p.Manager.Create)
		manager.POST("/login", p.Manager.Login)
		manager.POST("/logout ", p.Manager.Logout)
		manager.GET("/list", p.Manager.List)
		m := manager.Group("/:id", p.Manager.Mid)
		{
			m.GET("", p.Manager.Get)
			m.POST("/edit", p.Manager.Update)
			m.POST("/del", p.Manager.Del)
		}
	}

	captcha := routerGroup.Group("/captcha")
	{
		captcha.GET("", p.Captcha.GetCaptcha)
		captcha.GET("/image/:captchaId", p.Captcha.GetCaptchaImg)
		captcha.POST("/verify/:captchaId/:value", p.Captcha.VerifyCaptcha)
	}

	socket := routerGroup.Group("/socket")
	{
		socket.POST("", p.Socket.Create)
	}

	project := routerGroup.Group("/project")
	{
		project.POST("/create", p.Project.Create)
	}

	routerGroup.GET("/email", p.Tool.SendEmail)
}
