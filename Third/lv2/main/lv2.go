package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64 = 0
var ch1 = make(chan int64)
var ch2 = make(chan int64)

func add2() {
	for i := 0; i < 50; i++ {
		ch2 <- 1
		x += <-ch1
		fmt.Printf("add1: %v\n", x)
	}
	wg.Done()
}

func add1() {
	for i := 0; i < 50; i++ {
		x += <-ch2
		ch1 <- 1
		fmt.Printf("add2: %v\n", x)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add2()
	go add1()
	wg.Wait()
}
