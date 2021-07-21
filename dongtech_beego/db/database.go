package db

import (
	"dongtech/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

func DataBaseInit() {
	// set default database
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable", 30)

	// register model
	orm.RegisterModel(
		new(models.User),
		new(models.Post),
	)

	// create table
	orm.RunSyncdb("default", false, true)

	//根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC

	//打印sql
	orm.Debug = true
}
