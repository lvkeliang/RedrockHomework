package main

import (
	"fmt"
)

type Redrock struct {
	title     string
	code      string
	agreement string
}

type Cqupt struct {
	title     string
	code      string
	agreement string
}

type operation interface {
	open() error
	close() error
}

func (redrock Redrock) open() error {
	fmt.Printf("Redrock opened!\n")
	fmt.Printf("%v %v %v\n", redrock.title, redrock.agreement, redrock.code)
	return nil
}

func (redrock Redrock) close() error {
	fmt.Printf("Redrock closed!\n")
	return nil
}

func (cqupt Cqupt) open() error {
	fmt.Printf("cqupt opened!\n")
	fmt.Printf("%v %v %v\n", cqupt.title, cqupt.agreement, cqupt.code)
	return nil
}

func (cqupt Cqupt) close() error {
	fmt.Printf("cqupt closed!\n")
	return nil
}

func operate(web operation) {
	web.open()
	web.close()
}

func main() {
	var redrock Redrock
	redrock.title = "redrock"
	redrock.code = "utf-8"
	redrock.agreement = "https"

	var cqupt Cqupt
	cqupt.title = "cqupt"
	cqupt.code = "utf-16"
	cqupt.agreement = "http"

	/*cqupt := Cqupt{
		title:     "cqupt",
		code:      "utf-16",
		agreement: "http",
	}*/

	//var oper operation = new(Redrock)
	operate(redrock)
	//oper = new(Cqupt)
	operate(cqupt)
}
