// hello.go
package main

//#include <lib/hello.h>
import "C"

import "fmt"

//export SayHello4
func SayHello4(s *C.char) {
	fmt.Print(C.GoString(s))
}

func demo4() {
	C.SayHello4(C.CString("Hello, World4\n"))
}
