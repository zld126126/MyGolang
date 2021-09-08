// +build wireinject

package boot

import (
	"my_wx_app/config"
	"my_wx_app/web"

	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.DefaultConfig,
)

func InitWeb() (*web.WebApp, error) {
	panic(wire.Build(
		wire.Struct(new(web.WebApp), "*"),
		configSet,
		config.DefaultWebDatabase,
	))
}
