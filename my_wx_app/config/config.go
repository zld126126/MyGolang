package config

import (
	"my_wx_app/database"

	"github.com/shenghui0779/gochat"
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/mp"
	"github.com/shenghui0779/gochat/oa"
	"github.com/zld126126/dongo_utils/dongo_utils"
)

// 核心配置
const (
	WX_APPID     = ""
	WX_MCHID     = ""
	WX_APIKEY    = ""
	WX_APPSECRET = ""

	WEB_ADDR = "9090"

	DB_WEB_USER     = "root"
	DB_WEB_PASSWORD = "123456"
	DB_WEB_ADDR     = "localhost"
	DB_WEB_DATABASE = "dongbao"
	DB_WEB_POOLSIZE = 20
	DB_WEB_SLOW     = 50
	DB_WEB_PORT     = 3306
)

type Config struct {
	WebAddr     string
	WxConfig    *WxConfig
	WebDatabase *database.Database
}

type WxConfig struct {
	WxMch *mch.Mch
	WxOa  *oa.OA
	WxMp  *mp.MP
}

func NewDatabaseWeb() *database.Database {
	return &database.Database{
		User:     DB_WEB_USER,
		Password: DB_WEB_PASSWORD,
		Addr:     DB_WEB_ADDR,
		Database: DB_WEB_DATABASE,
		PoolSize: DB_WEB_POOLSIZE,
		Slow:     DB_WEB_SLOW,
		Port:     DB_WEB_PORT,
	}
}

func DefaultConfig() *Config {
	wxConfig := &WxConfig{
		WxMch: gochat.NewMch(WX_APPID, WX_MCHID, WX_APIKEY),
		WxOa:  gochat.NewOA(WX_APPID, WX_APPSECRET),
		WxMp:  gochat.NewMP(WX_APPID, WX_APPSECRET),
	}
	WebDatabase := NewDatabaseWeb()

	return &Config{
		WxConfig:    wxConfig,
		WebAddr:     WEB_ADDR,
		WebDatabase: WebDatabase,
	}
}

func DefaultWebDatabase(config *Config) *database.DB {
	db := &database.DB{
		Gorm: database.NewGormDB_Mysql(config.WebDatabase),
	}
	err := db.InitModel_Web()
	if err != nil {
		dongo_utils.Chk(err)
	}
	return db
}
