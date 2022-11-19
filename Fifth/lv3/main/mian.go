package main

import (
	"Readrockhomework/Fifth/lv3/main/funcs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用于判断是否已经登录
// 若未登录则重定向到/login
func checklogin(c *gin.Context) {
	flag := true
	loggername, _ := c.Cookie("username")
	inf := funcs.ReadFile("./Fifth/lv3/main/users.data")
	for _, data := range inf {
		for _, value := range data {
			if value["name"] == loggername {
				flag = false
				break
			}
		}
		if flag {
			//fmt.Println("不通过")
			c.SetCookie("username", "", -1, "/", "localhost", false, false)
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}
	}
	funcs.Save(inf)
}

// 用于访问/login和/register时判断是否已经登录
// 若已经登录则重定向到/homepage
func checker(c *gin.Context) {
	flag := true
	loggername, _ := c.Cookie("username")
	inf := funcs.ReadFile("./Fifth/lv3/main/users.data")
	for _, data := range inf {
		for _, value := range data {
			if value["name"] == loggername {
				flag = false
				funcs.Save(inf)
				break
			}
		}
		if !flag {
			funcs.Save(inf)
			c.Redirect(http.StatusMovedPermanently, "/homepage")
			c.Abort()
		}
	}
}

func main() {
	//userdata := make(map[string]string)
	r := gin.Default()
	r.LoadHTMLGlob("./Fifth/lv3/main/default/*")

	//访问根目录时若已经登录则重定向到/homepage，否则重定向到/login
	r.GET("/", checklogin, func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/homepage")
	})

	//登录成功后的页面
	r.GET("/homepage", checklogin, func(c *gin.Context) {
		loggername, _ := c.Cookie("username")
		fmt.Printf("\ncookie : %v\n", loggername)
		c.HTML(http.StatusOK, "default/homepage.html", gin.H{})
	})

	//注销
	//将cookie清除并重定向到根目录
	r.POST("/homepage", func(c *gin.Context) {
		c.SetCookie("username", "", -1, "/", "localhost", false, false)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	//登录页面
	r.GET("/login", checker, func(c *gin.Context) {
		loggername, _ := c.Cookie("username")
		fmt.Printf("\ncookie : %v\n", loggername)
		c.HTML(http.StatusOK, "default/login.html", gin.H{})
	})

	//对提交的表单进行登录
	//若成功则重定向到/homepage
	//若不成功则重定向到/login
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		inf := funcs.ReadFile("./Fifth/lv3/main/users.data")
		flag := false
		data := inf["data"]
		for _, value := range data {
			//fmt.Printf("%v  %v\n", value["name"], value["password"])
			if value["name"] == username && value["password"] == password {
				fmt.Println("成功")
				c.SetCookie("username", username, 0, "/", "localhost", false, false)
				flag = true
				break
			}
		}
		funcs.Save(inf)
		if flag {
			//fmt.Printf("flag = %v\n", flag)
			c.Redirect(http.StatusMovedPermanently, "/homepage")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
	})

	//注册页面
	r.GET("/register", checker, func(c *gin.Context) {
		//fmt.Println("\n来到了注册界面\n\n")
		c.HTML(http.StatusOK, "default/register.html", gin.H{})
	})

	//对提交的表单进行注册
	//若注册成功就登录
	//若不成功就回到/login
	r.POST("/register", func(c *gin.Context) {
		flag := true
		username := c.PostForm("username")
		password := c.PostForm("password")
		inf := funcs.ReadFile("./Fifth/lv3/main/users.data")
		//判断用户名是否已经注册
		for _, data := range inf {
			for _, value := range data {
				if value["name"] == username {
					flag = false
					break
				}
			}
		}
		if flag {
			//注册成功就重定向到login并登录
			funcs.Register(username, password, inf)
			funcs.Save(inf)
			c.Request.URL.Path = "/login"
			r.HandleContext(c)
		} else {
			//注册不成功就重定向到login
			funcs.Save(inf)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

	})

	r.Run()
}
