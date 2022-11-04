package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	//filepath.Join(wd, "/Third/lv3/main")
	os.Chdir(wd + "/Third/lv3/main")
	fmt.Println(wd)
	var err error
	content := make([]byte, 1024)
	file, err := os.Create("plan.txt")
	//fmt.Println("create")
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte("I’m not afraid of difficulties and insist on learning programming"))
	//fmt.Println("write")
	if err != nil {
		panic(err)
	}
	file.Close()
	file, err = os.Open("plan.txt")
	file.Read(content)
	//fmt.Println("read")
	fmt.Println(string(content))
	file.Close()
}
