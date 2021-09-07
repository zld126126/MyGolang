package config

import (
	"encoding/json"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"dongo_game_server/service/inf"
	"dongo_game_server/src/database"
	"dongo_game_server/src/global_const"

	"github.com/zld126126/dongo_utils/dongo_utils"
)

func DefaultConfig() *Config {
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func Grpc_DefaultUserService(config *Config) inf.UserServiceClient {
	conn, err := grpc.Dial(config.Grpc.UserServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return inf.NewUserServiceClient(conn)
}

func init() {
	configAddress := global_const.ConfigFileAddressDebug

	if gin.Mode() == gin.ReleaseMode {
		configAddress = global_const.ConfigFileAddressRelease
	}
	viper.SetDefault(global_const.ConfigFileKey, configAddress)

	err := configInit()
	if err != nil {
		dongo_utils.Chk(err)
	}
}

func configInit() error {
	configFilePath := viper.GetString(global_const.ConfigFileKey)
	var conf Config
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		return err
	}
	configStr, err := dongo_utils.ToJsonString(conf)
	if err != nil {
		logrus.WithError(err).Println("config init error")
		return err
	}
	viper.SetDefault(global_const.ConfigKey, configStr)
	return nil
}

func GetConfig() (*Config, error) {
	configStr := viper.GetString(global_const.ConfigKey)
	config, err := buildConfig(configStr)
	if err != nil {
		return nil, err
	}
	viper.SetDefault(global_const.ConfigVersionKey, config.Base.Version)
	return config, nil
}

func buildConfig(str string) (*Config, error) {
	var config Config
	err := json.Unmarshal([]byte(str), &config)
	return &config, errors.WithStack(err)
}

type Config struct {
	Base         *Base              `json:"base"`
	DatabaseWeb  *database.Database `json:"databaseWeb"`
	DatabaseGrpc *database.Database `json:"databaseGrpc"`
	Grpc         *GrpcConfig        `json:"grpc"`
	Web          *WebConfig         `json:"web"`
	Email        *EmailConfig       `json:"email"`
}

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewDatabase_Web(config *Config) *database.DB {
	db := &database.DB{
		Gorm: database.NewGormDB_Mysql(config.DatabaseWeb),
	}
	err := db.InitModel_Web()
	dongo_utils.Chk(err)
	return db
}

func NewDatabase_Grpc(config *Config) *database.DB {
	db := &database.DB{
		Gorm: database.NewGormDB_Mysql(config.DatabaseGrpc),
	}
	err := db.InitModel_Grpc()
	dongo_utils.Chk(err)
	return db
}

var Memory *dongo_utils.Memory

func DefaultMemory(config *Config) *dongo_utils.Memory {
	if Memory == nil {
		Memory = dongo_utils.NewMemory(config.Base.ProjectName)
	}
	return Memory
}

func DefaultGrpcConfig(config *Config) *GrpcConfig {
	return config.Grpc
}

func DefaultEmailConfig(config *Config) *EmailConfig {
	return config.Email
}

type Base struct {
	Author      string `json:"author"`
	Age         int    `json:"age"`
	Version     string `json:"version"`
	ProjectName string `json:"projectName"`
}

type GrpcConfig struct {
	UserServiceAddr string `json:"userServiceAddr"`
}

type WebConfig struct {
	Addr string `json:"addr"`
}
