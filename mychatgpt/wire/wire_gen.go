// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"mychatgpt/app"
	"mychatgpt/app/controller"
	"mychatgpt/app/service"
)

// Injectors from wire.go:

func InitApp() *app.App {
	engine := NewEngine()
	openAIService := NewService()
	controller := NewController(openAIService)
	appApp := NewApp(engine, controller)
	return appApp
}

// wire.go:

func NewApp(eng *gin.Engine, controller2 *controller.Controller) *app.App {
	return &app.App{
		Server:     eng,
		Controller: controller2,
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
