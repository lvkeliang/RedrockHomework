package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {

		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()
			if i == 9 {
				//等待所有协程完成后再向通道传入true
				wg.Wait()
				over <- true
			}
		}(i)

	}
	<-over
	fmt.Println("over!!!")
}

/*
func main() {
	over := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 9 {
				over <- true
			}

		}
	}()
	<-over
	fmt.Println("over!!!")
}
*/
