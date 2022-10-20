package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Ascending bubbling sort
func bubb1(arr []int) []int {
	m := 0
	mid := 0
	count := 1
	for count != 0 {
		count = 0
		//fmt.Printf("count = %d\n", count)
		for m+1 < len(arr) {
			//fmt.Printf("n = %d,m= %d\n", n, m)
			if arr[m] > arr[m+1] {
				mid = arr[m]
				//fmt.Printf("m = %d,arr[m] = %d\nn = %d,arr[n] = %d\n", m, arr[m], n, arr[n])
				arr[m] = arr[m+1]
				arr[m+1] = mid
				count++
				//fmt.Printf("count = %d\n", count)
			}
			m++
		}
		m = 0

	}
	return arr
}

// Descending bubbling sort
func bubb2(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for n := 0; n < length-1-i; n++ {
			if arr[n] < arr[n+1] {
				arr[n], arr[n+1] = arr[n+1], arr[n]
			}
		}
	}
	return arr
}

// Selection sort
func sele(arr []int) []int {
	length := len(arr)
	var min int
	for i := 0; i < length-1; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// Insertion sort
func inse(arr []int) []int {
	for i := range arr {
		pre := i - 1
		current := arr[i]
		//fmt.Printf("i = %d,", i)
		for pre >= 0 && arr[pre] > current {
			arr[pre+1] = arr[pre]
			pre--
		}
		arr[pre+1] = current
	}
	return arr
}

func main() {
	var arr = make([]int, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		//fmt.Printf("i = %d,", i)
		arr[i] = rand.Intn(1000)
		//fmt.Printf("arr[i] = %d\n", arr[i])
	}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Printf("len = %d\n", len(arr))
	arr = inse(arr)
	fmt.Print(arr)
}
