package main

import "fmt"

func main() {
	fmt.Println("go call cppDLL")
	GoCallDllTest()

	//生成golang-dll命令
	//go build -o go.dll -buildmode=c-shared main.go
	fmt.Println("go call goDLL")
	GoCallDll(3, 5)

	Pause()
}
