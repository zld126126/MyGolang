package main

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lgo_lib
#include "libgo_lib.h"
*/
import "C"
import "fmt"

func main() {
	testDLL()
	testCgo()

	Pause()
}

func testDLL() {
	//WIN32
	//fmt.Println("go call cppDLL")
	//GoCallDllTest()

	fmt.Println("go call goDLL")
	GoCallDll(3, 5)
}

func testCgo() {
	a := C.Add(C.longlong(3), C.longlong(5))
	fmt.Println(int(a))
}
