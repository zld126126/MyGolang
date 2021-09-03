// +build wireinject

package boot

import (
	"dongo_game_server/src/config"
	"dongo_game_server/src/grpc"
	"dongo_game_server/src/support"
	"dongo_game_server/src/web"
	"dongo_game_server/src/web/controller"

	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.DefaultConfig,
	config.DefaultEmailConfig,
	config.DefaultGrpcConfig,
	config.Grpc_DefaultUserService,
)

var webSet = wire.NewSet(
	wire.Struct(new(controller.BaseHdl), "*"),
	wire.Struct(new(controller.CaptchaHdl), "*"),
	wire.Struct(new(controller.JWTHdl), "*"),
	wire.Struct(new(controller.ManagerHdl), "*"),
	wire.Struct(new(controller.ProjectHdl), "*"),
	wire.Struct(new(controller.ResourceHdl), "*"),
	wire.Struct(new(controller.RpcHdl), "*"),
	wire.Struct(new(controller.SocketHdl), "*"),
	wire.Struct(new(controller.ToolHdl), "*"),
	wire.Struct(new(controller.TrackHdl), "*"),
)

func InitWeb() (*web.WebApp, error) {
	panic(wire.Build(
		wire.Struct(new(web.WebApp), "*"),
		configSet,
		webSet,
		config.NewDatabase_Web,
	))
}

func InitGrpc() (*grpc.GrpcApp, error) {
	panic(wire.Build(
		wire.Struct(new(grpc.GrpcApp), "*"),
		configSet,
		config.NewDatabase_Grpc,
	))
}

func InitSupport() (*support.SupportApp, error) {
	panic(wire.Build(
		wire.Struct(new(support.SupportApp), "*"),
		configSet,
		config.NewDatabase_Web,
	))
}
