// +build wireinject

package boot

import (
	"dongo_game_server/src/config"
	"dongo_game_server/src/grpc"
	"dongo_game_server/src/support"
	"dongo_game_server/src/web"
	"dongo_game_server/src/web/controller"
	"dongo_game_server/src/web/service"

	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.DefaultConfig,
	config.DefaultEmailConfig,
	config.DefaultRpcConfig,
	config.DefaultUserServiceRpc,
	config.DefaultMemory,
	config.DefaultRedis,
)

var webSet = wire.NewSet(
	wire.Struct(new(controller.BaseHdl), "*"),
	wire.Struct(new(controller.CaptchaHdl), "*"),
	wire.Struct(new(controller.ManagerHdl), "*"),
	wire.Struct(new(controller.ProjectHdl), "*"),
	wire.Struct(new(controller.ResourceHdl), "*"),
	wire.Struct(new(controller.RpcHdl), "*"),
	wire.Struct(new(controller.SocketHdl), "*"),
	wire.Struct(new(controller.ToolHdl), "*"),
	wire.Struct(new(controller.TrackHdl), "*"),
	wire.Struct(new(controller.ManagerPathHdl), "*"),
	wire.Struct(new(controller.FakeHdl), "*"),
	wire.Struct(new(controller.ClientHdl), "*"),

	wire.Struct(new(service.ManagerService), "*"),
	wire.Struct(new(service.SocketService), "*"),
	wire.Struct(new(service.ProjectService), "*"),
	wire.Struct(new(service.ManagerPathService), "*"),
	wire.Struct(new(service.ClientService), "*"),
)

func InitWeb() (*web.WebApp, error) {
	panic(wire.Build(
		wire.Struct(new(web.WebApp), "*"),
		configSet,
		webSet,
		config.NewDatabaseWeb,
	))
}

func InitRpc() (*grpc.RpcApp, error) {
	panic(wire.Build(
		wire.Struct(new(grpc.RpcApp), "*"),
		configSet,
		config.NewDatabaseRpc,
	))
}

func InitSupport() (*support.SupportApp, error) {
	panic(wire.Build(
		wire.Struct(new(support.SupportApp), "*"),
		configSet,
		config.NewDatabaseWeb,
	))
}
