// hello.go
package main

//#include <stdio.h>
import "C"

func demo1() {
	C.puts(C.CString("Hello, World1\n"))
}
