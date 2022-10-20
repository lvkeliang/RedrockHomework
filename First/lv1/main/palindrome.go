package main

import (
	"bytes"
	"fmt"
)

// Get the right characters.
func Index(s string, index uint) string {
	runes := bytes.Runes([]byte(s))
	for i, rune := range runes {
		if i == int(index) {
			return string(rune)
		}
	}
	return ""
}

// Get the right length of a string.
func Strlen(s string) uint {
	var a uint = 0
	for i := " "; i != ""; a++ {
		i = Index(s, a)
	}
	return a - 1
}

func main() {
	var strin, strout string
	var cond bool = true
	fmt.Scan(&strin)
	var a uint = 0
	b := Strlen(strin)
	b--

	for a < b {
		if Index(strin, a) == Index(strin, b) {
			strout = strout + Index(strin, a)
			a++
			b--
		} else {
			cond = false
			break
		}
	}
	if cond == true {
		fmt.Print(strout)
	} else {
		fmt.Print("false")
	}

}
