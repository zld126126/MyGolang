package service

import "github.com/google/wire"

type Option struct {
	Message []string
	Ip      string
}

type ConfigService struct {
	Message []string
	Ip      string
}

func NewConfigService(option *Option) (*ConfigService, error) {
	return &ConfigService{
		Message: option.Message,
		Ip:      option.Ip,
	}, nil
}

func (p *ConfigService) Get() {
	println(p.Ip)
}

var ConfigSet = wire.NewSet(NewConfigService, wire.Struct(new(Option), "*"))
