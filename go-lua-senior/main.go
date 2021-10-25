package main

import (
	"github.com/Shopify/go-lua"
	"go-lua-senior/util"
	"log"
)

func main() {
	event()
	callback()
}

// event 测试:在luaGoFunction中调用golang的注册事件
func event() {
	m := util.NewManager()
	err := lua.DoFile(m.Lua, "test1.lua")
	if err != nil {
		log.Fatal(err)
	}
}

// callback 测试:在lua调用lua的方法作为luaGoFunc的参数,即回调测试
func callback() {
	m := util.NewManager()
	err := lua.DoFile(m.Lua, "test2.lua")
	if err != nil {
		log.Fatal(err)
	}
}
