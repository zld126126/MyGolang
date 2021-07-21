// +build wireinject

package boot

import "github.com/google/wire"

func InitHandle() (*Handle, func(), error) {
	panic(wire.Build(wire.Struct(
		new(Handle), "*"),
		baseSet, configSet, databaseSet,
	))
}

var configSet = wire.NewSet(DefaultConfig)
var databaseSet = wire.NewSet(NewDatabase)
var baseSet = wire.NewSet(initMicroServiceClient, initAssistService, initMicroService, initMicroWebService)
