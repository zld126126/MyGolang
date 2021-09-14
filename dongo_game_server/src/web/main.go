package web

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/config"
	"dongo_game_server/src/web/controller"
	"fmt"
	"log"

	_ "dongo_game_server/src/web/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebApp struct {
	Config      *config.Config
	UserService inf.UserServiceClient

	Base    *controller.BaseHdl
	Captcha *controller.CaptchaHdl

	Manager     *controller.ManagerHdl
	Project     *controller.ProjectHdl
	Resource    *controller.ResourceHdl
	Tool        *controller.ToolHdl
	Track       *controller.TrackHdl
	ManagerPath *controller.ManagerPathHdl

	Client *controller.ClientHdl
	RPC    *controller.RpcHdl
	Socket *controller.SocketHdl

	Fake *controller.FakeHdl
}

func (p *WebApp) Start() {
	p.Socket.InitSocket()

	router := gin.New()

	router.Use(ServeRecover)
	//router.Use(gin.Recovery())
	// router.LoadHTMLGlob("./resources")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerGroup := router.Group("")
	p.Mount(routerGroup)

	err := router.Run(fmt.Sprintf(`:%s`, p.Config.Web.Addr))
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`WebApp Start error`)
		log.Fatal(err)
	}

	log.Println("web serve running")
}

func (p *WebApp) Mount(routerGroup *gin.RouterGroup) {
	// 无需登陆即可获取
	baseGroup := routerGroup.Group("/base")
	{
		baseGroup.GET("/version", p.Base.GetVersion())

		captcha := baseGroup.Group("/captcha")
		{
			captcha.GET("", p.Captcha.GetCaptcha)
			captcha.GET("/image/:captchaId", p.Captcha.GetCaptchaImg)
			captcha.POST("/verify/:captchaId/:value", p.Captcha.VerifyCaptcha)
		}

		manager := baseGroup.Group("/manager")
		{
			manager.POST("/login", p.Manager.Login)
			manager.POST("/logout ", p.Manager.Logout)
		}

		managerPath := baseGroup.Group("/manager_path")
		{
			managerPath.POST("/create", p.ManagerPath.Create)
			path := managerPath.Group("/:id", p.ManagerPath.Mid)
			{
				path.GET("", p.ManagerPath.Create)
				path.POST("/update", p.ManagerPath.Update)
				path.POST("/del", p.ManagerPath.Del)
			}
			managerPath.GET("/full_list", p.ManagerPath.FullList)
		}
	}

	// web后台对接api
	webGroup := routerGroup.Group("/web", p.Manager.MidLogin)
	{
		manager := webGroup.Group("/manager")
		{
			manager.POST("/create", p.Manager.IsSuperManager, p.Manager.Create)
			manager.POST("/login", p.Manager.Login)
			manager.POST("/logout ", p.Manager.Logout)
			manager.GET("/list", p.Manager.List)
			manager.GET("/refresh", p.Manager.Refresh)
			m := manager.Group("/:id", p.Manager.Mid)
			{
				m.GET("", p.Manager.Get)
				m.POST("/edit", p.Manager.IsSuperManager, p.Manager.Update)
				m.POST("/del", p.Manager.IsSuperManager, p.Manager.Del)
				m.POST("/path_list", p.Manager.IsSuperManager, p.Manager.PathList)
				m.POST("/path_bind", p.Manager.IsSuperManager, p.Manager.PathBind)
				m.POST("/path_unbind", p.Manager.IsSuperManager, p.Manager.PathUnBind)
			}
		}

		project := webGroup.Group("/project")
		{
			project.POST("/create", p.Project.Create)
		}

		webGroup.GET("/email", p.Tool.SendEmail)
	}

	// client后台对接api
	clientGroup := routerGroup.Group("/client")
	{
		socket := clientGroup.Group("/socket")
		{
			socket.POST("", p.Socket.Create)
		}

		http := clientGroup.Group("/http")
		{
			http.POST("", p.Track.Collect)
		}

		rpc := clientGroup.Group("/rpc")
		{
			rpc.GET("/user/:user_id", p.RPC.GetUser)
		}
	}

	// debug 命令
	debugGroup := routerGroup.Group("/debug", p.Fake.Mid)
	{
		debugGroup.POST("/manager/create", p.Manager.Create)
	}
}
