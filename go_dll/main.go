package main

import "fmt"

func main() {
	fmt.Println("go call cppDLL")
	GoCallDllTest()
	
	fmt.Println("go call goDLL")
	GoCallDll(3, 5)

	Pause()
}
