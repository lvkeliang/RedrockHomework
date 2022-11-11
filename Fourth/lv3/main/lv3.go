package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type information struct {
	name     string
	password string
}

// 定义读取文件函数
// 该函数将json格式的数据文件读取为map[string][]map[string]string格式
// filewd:读取的文件的路径
func readFile(filewd string) map[string][]map[string]string {
	var inf map[string][]map[string]string
	str, err := ioutil.ReadFile(filewd)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(str, &inf)
	return inf
}

// 定义注册函数
// username:注册名
// passwor：密码
// inf：inf （作用的数据对象）
func register(username string, password string, inf map[string][]map[string]string) {
	var flag int = 0
	if len(password) <= 6 {
		flag = 1
	} else {
		for _, value := range inf["data"] {
			//fmt.Println(value)
			if value["name"] == username {
				flag = 2
				break
			}
		}
	}
	if flag == 0 {
		inf["data"] = append(inf["data"], map[string]string{
			"name":     username,
			"password": password,
		})
	} else if flag == 1 {
		fmt.Println("The length of password must be greater than 6")
	} else if flag == 2 {
		fmt.Println("This user name has been registered!")
	}
}

func save(inf map[string][]map[string]string) {
	file, err := os.OpenFile("users.data", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	a, _ := json.Marshal(inf)
	file.Write(a)
	file.Close()
}

func main() {
	var username, password string
	os.Chdir("./Fourth/lv3/main")
	file, _ := os.OpenFile("users.data", os.O_CREATE|os.O_RDONLY, 0644)
	file.Close()
	inf := readFile("users.data")
	fmt.Println("Please input your username:")
	fmt.Scanf("%v", &username)
	fmt.Scanln()
	fmt.Println("Please input your password:")
	fmt.Scanf("%v", &password)
	register(username, password, inf)
	//str, _ := ioutil.ReadFile("users.data")
	//json.Unmarshal(str, &users)
	fmt.Println(inf)
	save(inf)
	os.Exit(0)

}
