package main

import "C"

//export Add
func Add(a, b int) int {
	return a + b
}

//编译生成golang-lib命令
//go build -buildmode=c-archive go_lib.go
//gcc go_lib.def go_lib.a -shared -lwinmm -lWs2_32 -o go_lib.dll -Wl,--out-implib,go_lib.lib

//编译时 放开main()
//func main() {}