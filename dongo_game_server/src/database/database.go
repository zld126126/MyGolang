package database

import (
	"dongo_game_server/src/model"
	"fmt"
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

func NewMysqlGormDB(db *Database) *gorm.DB {
	instance, err := gorm.Open("mysql", connMysql(db))
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

func connMysql(db *Database) string {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		db.User, db.Password, db.Addr, db.Port, db.Database)
	return connStr
}

func NewPostgresGormDB(db *Database) *gorm.DB {
	instance, err := gorm.Open("postgres", connPostgres(db))
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

func connPostgres(db *Database) string {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		db.Addr, db.User, db.Database, db.Password, db.Port)
	return connStr
}

func (p *DB) InitWebModel() error {
	logrus.Println(`InitWebModel start`)
	models := []interface{}{
		model.SocketConfig{},

		model.Manager{},
		model.ManagerPath{},
		model.ManagerPathRelation{},

		model.Track{},
		model.TrackItem{},

		model.Consumer{},
		model.ConsumerItem{},

		model.Project{},
		model.ProjectConsumerRelation{},
	}
	err := p.Gorm.Debug().AutoMigrate(models...).Error
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`InitWebModel err`)
		return errors.WithStack(err)
	}
	logrus.Println(`InitWebModel end`)
	return nil
}

func (p *DB) InitRpcModel() error {
	logrus.Println(`InitRpcModel start`)
	models := []interface{}{
		model.User{},
	}
	err := p.Gorm.Debug().AutoMigrate(models...).Error
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`InitRpcModel error`)
		return errors.WithStack(err)
	}
	logrus.Println(`InitRpcModel end`)
	return nil
}

func (p *DB) IsGormNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
