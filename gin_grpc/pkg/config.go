package pkg

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"gin_grpc/database"
	"gin_grpc/pkg/model"
	"gin_grpc/service/inf"
	"gin_grpc/util"
)

func DefaultConfig() *Config {
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func DefaultUserService(config *Config) inf.UserServiceClient {
	conn, err := grpc.Dial(config.UserService.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return inf.NewUserServiceClient(conn)
}

const (
	ConfigFileAddressRelease = `../resources/config.toml` // go build正式环境用
	ConfigFileAddressDebug   = `resources/config.toml`    // goland本地启动用
	ConfigFileKey            = `configFile`
	ConfigKey                = `config`
)

func init() {
	configAddress := ConfigFileAddressDebug
	if gin.Mode() == gin.ReleaseMode {
		configAddress = ConfigFileAddressRelease
	}
	viper.SetDefault(ConfigFileKey, configAddress)

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
	config, err := buildConfig(configStr)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func buildConfig(str string) (*Config, error) {
	var config Config
	err := json.Unmarshal([]byte(str), &config)
	return &config, errors.WithStack(err)
}

type Config struct {
	Base        *Base              `json:"base"`
	Database    *database.Database `json:"database"`
	UserService *UserService       `json:"userService"`
	Web         *Web               `json:"web"`
	Email       *util.EmailConfig  `json:"email"`
}

func NewDatabase(config *Config) *database.DB {
	db := &database.DB{
		Gorm: database.NewGormDB(config.Database),
	}
	err := initModel(db.Gorm)
	util.Catch(err)
	return db
}

func initModel(db *gorm.DB) error {
	logrus.Println(`init model start`)
	models := []interface{}{
		model.User{},
	}
	err := db.Debug().AutoMigrate(models...).Error
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`init model err`)
		return errors.WithStack(err)
	}
	logrus.Println(`init model end`)
	return nil
}

type Base struct {
	Author string `json:"author"`
	Age    int    `json:"age"`
}

type UserService struct {
	Addr string `json:"addr"`
}

type Web struct {
	Addr string `json:"addr"`
}
