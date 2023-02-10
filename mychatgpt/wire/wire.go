//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	gin "github.com/gin-gonic/gin"
	wire "github.com/google/wire"
	app "mychatgpt/app"
	"mychatgpt/app/controller"
	service "mychatgpt/app/service"
)

func InitApp() *app.App {
	panic(wire.Build(NewApp, NewEngine, NewController, NewService))
}

func NewApp(eng *gin.Engine, controller *controller.Controller) *app.App {
	return &app.App{
		Server:     eng,
		Controller: controller,
	}
}

func NewEngine() *gin.Engine {
	return gin.Default()
}

func NewController(openAI *service.OpenAIService) *controller.Controller {
	return &controller.Controller{
		OpenAI: openAI,
	}
}

func NewService() *service.OpenAIService {
	return &service.OpenAIService{
		Token: "",
	}
}
