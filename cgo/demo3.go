package main

//#include"lib/hello.h"
import "C"

func demo3(){
	C.SayHello(C.CString("Hello, World3\n"))
}