// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package boot

import (
	"github.com/google/wire"
	"my_wx_app/config"
	"my_wx_app/web"
)

// Injectors from wire.go:

func InitWeb() (*web.WebApp, error) {
	configConfig := config.DefaultConfig()
	db := config.DefaultWebDatabase(configConfig)
	webApp := &web.WebApp{
		Config: configConfig,
		DB:     db,
	}
	return webApp, nil
}

// wire.go:

var configSet = wire.NewSet(config.DefaultConfig)
