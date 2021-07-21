package routers

import (
	"dongtech/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//1.访问localhost:9090/user
	beego.Router("/info", &controllers.InfoController{})
	//2.访问localhost:9090/sayHello?name='dongtech'
	beego.Router("/sayHello", &controllers.SayHelloController{})
	//3.访问localhost:9090/getUser?id=1
	beego.Router("/user", &controllers.UserController{})

	beego.ErrorHandler("/Error", controllers.Error)
	beego.ErrorController(&controllers.ErrorController{})
}
