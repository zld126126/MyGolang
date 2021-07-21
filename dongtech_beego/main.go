package main

import (
	"dongtech/db"
	_ "dongtech/routers"
	"github.com/astaxie/beego"
)

func init() {
	db.DataBaseInit()
}

func main() {
	beego.Run()
}
