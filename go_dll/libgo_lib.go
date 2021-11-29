package main

import "C"

//export Add
func Add(a, b int) int {
	return a + b
}

//编译生成golang-lib命令
//go build -buildmode=c-archive libgo_lib.go
//gcc libgo_lib.def libgo_lib.a -shared -lwinmm -lWs2_32 -o libgo_lib.dll -Wl,--out-implib,libgo_lib.lib

//编译时 放开main()
func main() {}
