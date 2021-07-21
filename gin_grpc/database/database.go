package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	Gorm *gorm.DB
}

func (db *DB) GetGorm() *gorm.DB {
	return db.Gorm
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

func NewGormDB(db *Database) *gorm.DB {
	instance, err := gorm.Open("postgres", connStr(db))
	if err != nil {
		return nil
	}
	instance.DB().SetConnMaxLifetime(time.Minute * 5)
	instance.DB().SetMaxIdleConns(10)
	instance.DB().SetMaxOpenConns(db.PoolSize)
	instance.LogMode(true)
	return instance
}

func connStr(db *Database) string {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		db.Addr, db.User, db.Database, db.Password, db.Port)
	return connStr
}
