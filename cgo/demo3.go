package main

//#include "lib/hello.c"
import "C"

func demo3() {
	C.SayHello3(C.CString("Hello, World3\n"))
}
