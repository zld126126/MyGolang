// hello.go
package main

/*
#include <stdio.h>

static void SayHello2(const char* s) {
    puts(s);
}
*/
import "C"

func demo2() {
	C.SayHello2(C.CString("Hello, World2\n"))
}