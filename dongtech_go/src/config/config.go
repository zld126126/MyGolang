package config

import (
	"dongtech_go/util"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigFileAddress = `../resources/config.toml` //go build正式环境用
	//ConfigFileAddress = `resources/config.toml` //goland本地启动用
	ConfigFileKey = `configFile`
	ConfigKey     = `config`
)

func init() {
	viper.SetDefault(ConfigFileKey, ConfigFileAddress)

	err := configInit()
	if err != nil {
		util.Catch(err)
	}
}

func configInit() error {
	configFilePath := viper.GetString(ConfigFileKey)
	var conf Config
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		return err
	}
	configStr, err := util.ToJsonString(conf)
	if err != nil {
		logrus.WithError(err).Println("config init error")
		return err
	}
	viper.SetDefault(ConfigKey, configStr)
	return nil
}

func GetConfig() (*Config, error) {
	configStr := viper.GetString(ConfigKey)
	config, err := toConfig(configStr)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func toConfig(str string) (*Config, error) {
	var config Config
	err := json.Unmarshal([]byte(str), &config)
	return &config, errors.WithStack(err)
}

type Config struct {
	Base     *Base     `json:"base"`
	Database *Database `json:"database"`
	Grpc     *Grpc     `json:"grpc"`
	Web      *Web      `json:"web"`
	Email    *Email    `json:"email"`
}

type Base struct {
	Author string `json:"author"`
	Age    int    `json:"age"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
	Database string `json:"database"`
	PoolSize int    `json:"pool_size"`
	Slow     int    `json:"slow"`
	Port     int    `json:"port"`
}

type Grpc struct {
	Addr string `json:"addr"`
}

type Web struct {
	Addr string `json:"addr"`
}

type Email struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
