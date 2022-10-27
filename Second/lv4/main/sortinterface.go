package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Str []string

// 定义返回长度的方法
func (str Str) Len() int {
	return len(str)
}

// 定义比较大小的方法
func (str Str) Less(i, j int) bool {
	return str[i] < str[j]
}

// 定义交换元素的方法
func (str Str) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}
func main() {
	stringlist := Str{"1", "3", "999", "0", "101", "520"}
	sort.Sort(stringlist)
	fmt.Printf("%v\n", stringlist)

}
