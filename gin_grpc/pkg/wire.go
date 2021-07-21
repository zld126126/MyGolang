// +build wireinject

package pkg

import (
	"github.com/google/wire"
)

var configSet = wire.NewSet(DefaultConfig)

func InitWeb() (*WebHandler, error) {
	panic(wire.Build(
		wire.Struct(new(WebHandler), "*"),
		configSet,
		DefaultUserService,
		NewDatabase,
	))
}

func InitUserService() (*UserServiceHandler, error) {
	panic(wire.Build(
		wire.Struct(new(UserServiceHandler), "*"),
		configSet,
		NewDatabase,
	))
}
