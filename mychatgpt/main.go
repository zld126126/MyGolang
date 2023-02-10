package main

import "mychatgpt/wire"

// main 主程序入口
func main() {
	app := wire.InitApp()
	app.Start()
}
