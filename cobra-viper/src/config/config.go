package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
)

const (
	config_file_address = `../resources/config.toml` //go build正式环境用
	//config_file_address = `resources/config.toml` goland本地启动用
	config_file = `config`
)

func init() {
	viper.SetDefault(config_file, config_file_address)
}

func GetConfig() (*Config, error) {
	configFilePath := viper.GetString(config_file)
	var conf Config
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		return nil, errors.New("decode config file err")
	}
	return &conf, nil
}

type Config struct {
	Base Base
}

type Base struct {
	Author string
	Age    int
}
