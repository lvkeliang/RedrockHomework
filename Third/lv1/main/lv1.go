package main

import (
	"fmt"
	"sync"
)

//原代码
/*
var x int64
var wg sync.WaitGroup
var mu sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		mu.Lock()
		x = x + 1
		mu.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
*/

// 用Channel实现
var wg sync.WaitGroup
var x int64 = 0
var ch1 = make(chan int64)
var ch2 = make(chan int64)

func add1() {
	for i := 0; i < 50000; i++ {
		ch2 <- 1
		x += <-ch1
	}
	wg.Done()
}

func add2() {
	for i := 0; i < 50000; i++ {
		x += <-ch2
		ch1 <- 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add1()
	go add2()
	wg.Wait()

	fmt.Println(x)
}
