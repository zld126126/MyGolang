package main

import "fmt"

func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}

func ChanDemo() {
	var ch chan int
	ch = make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	Parse(ch)
}


