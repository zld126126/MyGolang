package database

import (
	"fmt"
	"my_wx_app/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zld126126/dongo_utils/dongo_utils"
)

type DB struct {
	Gorm *gorm.DB
}

func (p *DB) GetGorm() *gorm.DB {
	return p.Gorm
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

func NewGormDB_Mysql(db *Database) *gorm.DB {
	instance, err := gorm.Open("mysql", connStr_Mysql(db))
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Println(db.Database + "connect error")
		dongo_utils.Chk(err)
	}
	instance.DB().SetConnMaxLifetime(time.Minute * 5)
	instance.DB().SetMaxIdleConns(10)
	instance.DB().SetMaxOpenConns(db.PoolSize)
	instance.LogMode(true)
	return instance
}

func connStr_Mysql(db *Database) string {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		db.User, db.Password, db.Addr, db.Port, db.Database)
	return connStr
}

func NewGormDB_Postgres(db *Database) *gorm.DB {
	instance, err := gorm.Open("postgres", connStr_Postgres(db))
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Println(db.Database + "connect error")
		dongo_utils.Chk(err)
	}
	instance.DB().SetConnMaxLifetime(time.Minute * 5)
	instance.DB().SetMaxIdleConns(10)
	instance.DB().SetMaxOpenConns(db.PoolSize)
	instance.LogMode(true)
	return instance
}

func connStr_Postgres(db *Database) string {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		db.Addr, db.User, db.Database, db.Password, db.Port)
	return connStr
}

func (p *DB) InitModel_Web() error {
	logrus.Println(`initModel_Web start`)
	models := []interface{}{
		model.Test{},
	}
	err := p.Gorm.Debug().AutoMigrate(models...).Error
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`initModel_Web err`)
		return errors.WithStack(err)
	}
	logrus.Println(`initModel_Web end`)
	return nil
}

func (p *DB) InitModel_Grpc() error {
	logrus.Println(`initModel_Grpc start`)
	models := []interface{}{
		//model.User{},
	}
	err := p.Gorm.Debug().AutoMigrate(models...).Error
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`initModel_Grpc err`)
		return errors.WithStack(err)
	}
	logrus.Println(`initModel_Grpc end`)
	return nil
}

func (p *DB) IsGormNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
