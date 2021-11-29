package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//lock是全局互斥锁,synchornized
	lock sync.Mutex
)

func GoRoutineFailedDemo1() {
	cal := func(n int){
		res := 1
		for i := 1; i <= n; i++ {
			res *= i
		}
		myMap[n] = res
	}

	for i := 1; i <= 15; i++ {
		go cal(i)
	}
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

func GoRoutineFailedDemo2() {
	cal := func(n int){
		res := 1
		for i := 1; i <= n; i++ {
			res *= i
		}
		lock.Lock()
		myMap[n] = res
		lock.Unlock()
	}

	for i := 1; i <= 15; i++ {
		go cal(i)
	}

	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

func GoRoutineSuccessDemo(){
	cal := func(n int){
		res := 1
		for i := 1; i <= n; i++ {
			res *= i
		}
		lock.Lock()
		myMap[n] = res
		lock.Unlock()
	}

	for i := 1; i <= 15; i++ {
		go cal(i)
	}

	time.Sleep(time.Second * 4)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}

