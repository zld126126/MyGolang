// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"awesomeProject/service"
	"github.com/google/wire"
)

func NewUserService(id int, name string) (*service.UserService, error) {
	wire.Build(service.NewUserService)
	return &service.UserService{}, nil
}

func InitLogService(name string) (*service.LogService, error) {
	wire.Build(service.LogSet)
	return &service.LogService{}, nil
}

func InitConfigService(msg []string, ip string) (*service.ConfigService, error) {
	wire.Build(service.ConfigSet)
	return &service.ConfigService{}, nil
}
