package main

import "fmt"

func main() {
	fmt.Println("go call cppDLL")
	GoCallDllTest()

	fmt.Println("go call go-lib DLL")
	GoCallDll(3, 5)

	Pause()
}
