package main

import (
	"fmt"
	"sync"
)

type Once struct {
	Count int
}

var st = &Once{Count: 0}
var one sync.Once

func OnceDemo() {
	add := func() *Once {
		initSt := func() {
			st.Count += 1
		}

		// sync.Once 保证只被调用一次
		one.Do(initSt)
		fmt.Println("OnceDemo:", st.Count)
		return st
	}

	for i := 0; i < 10; i++ {
		go add()
	}
}
