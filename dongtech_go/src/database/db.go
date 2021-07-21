package database

import (
	"dongtech_go/config"
	"dongtech_go/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"time"
)

func GetDB() *gorm.DB {
	conf, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Println("get config err")
		util.Catch(err)
	}
	instance, err := gorm.Open("postgres", connStr(conf))
	if err != nil {
		logrus.WithError(err).WithField("conn str", connStr(conf)).Println("get database err")
		util.Catch(err)
	}
	instance.DB().SetConnMaxLifetime(time.Minute * 5)
	instance.DB().SetMaxIdleConns(10)
	instance.DB().SetMaxOpenConns(conf.Database.PoolSize)
	instance.LogMode(true)
	return instance
}

func connStr(config *config.Config) string {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		config.Database.Addr, config.Database.User, config.Database.Database, config.Database.Password, config.Database.Port)
	return connStr
}
