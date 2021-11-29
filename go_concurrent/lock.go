package main

import (
	"fmt"
	"sync"
)

var (
	mu  sync.Mutex
	rmu sync.RWMutex
)

func LockDemo_Mutex() {
	count := 0
	add := func() {
		mu.Lock()
		defer mu.Unlock()
		count += 1
		fmt.Println("LockDemo_Mutex:", count)
	}

	for i := 0; i <= 5; i++ {
		go add()
	}
}

// 多读单写锁
func LockDemo_RWMutex() {
	count := 0
	add := func() {
		rmu.Lock()
		defer rmu.Unlock()
		count += 1
		fmt.Println("LockDemo_RWMutex add:",count)
	}

	get := func() int {
		rmu.RLock()
		defer rmu.RUnlock()
		fmt.Println("LockDemo_RWMutex get:",count)
		return count
	}

	for i := 0;i < 6; i++{
		if i % 2 == 0{
			go add()
		}else{
			go get()
		}
	}
}
