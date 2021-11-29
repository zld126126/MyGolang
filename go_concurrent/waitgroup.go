package main

import (
	"fmt"
	"sync"
)

func WaitGroupDemo_chan(){
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("WaitGroupDemo_chan:",i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

func WaitGroupDemo_waitGroup(){
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("WaitGroupDemo_waitGroup:",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}