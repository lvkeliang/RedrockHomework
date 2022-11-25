package main

import (
	"Readrockhomework/Sixth/lv3/main/funcs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

//type Finduser struct {
//	Username string `json:"Username"`
//}

// 用于判断是否已经登录
// 若未登录则重定向到/login
func checklogin(c *gin.Context) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/redrocksixth")
	flag := true
	loggername, _ := c.Cookie("username")

	rows, err := db.Query("select * from active_users where cookie= ?", loggername)
	if err != nil {
		fmt.Printf("\nerr : %v\n", err)
		rows.Close()
		return
	}
	defer rows.Close()
	var user struct {
		id     int
		cookie string
	}
	for rows.Next() {
		err := rows.Scan(&user.id, &user.cookie)
		if err != nil {
			log.Println(err)
			return
		}
		if user.cookie == loggername {

			flag = false
		}
	}
	if flag {
		//fmt.Println("不通过")
		c.SetCookie("username", "", -1, "/", "localhost", false, false)
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
	}
}

// 用于访问/login和/register时判断是否已经登录
// 若已经登录则重定向到/homepage
func checker(c *gin.Context) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/redrocksixth")
	flag := true
	loggername, _ := c.Cookie("username")
	rows, err := db.Query("select id,cookie from active_users where cookie= ?", loggername)
	if err != nil {
		fmt.Printf("\nerr : %v\n", err)
		rows.Close()
		return
	}
	defer rows.Close()
	var user struct {
		id     int
		cookie string
	}
	for rows.Next() {
		err := rows.Scan(&user.id, &user.cookie)
		if err != nil {
			log.Println(err)
			return
		}
		if user.cookie == loggername {

			flag = false
		}
	}
	if !flag {
		c.Redirect(http.StatusMovedPermanently, "/homepage")
		c.Abort()
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/redrocksixth")
	if err != nil {
		return
	}

	//用于找回时临时储存用户名和密保问题
	Findername := ""
	Question := ""

	//userdata := make(map[string]string)
	r := gin.Default()
	r.LoadHTMLGlob("./Sixth/lv3/main/default/*")
	//访问根目录时若已经登录则重定向到/homepage，否则重定向到/login
	r.GET("/", checklogin, func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/homepage")
	})

	//登录成功后的页面
	r.GET("/homepage", checklogin, func(c *gin.Context) {
		loggername, _ := c.Cookie("username")
		fmt.Printf("\ncookie : %v\n", loggername)
		//fmt.Printf("\nmsg : %v\n", funcs.Messages(db))
		c.HTML(http.StatusOK, "default/homepage.html", gin.H{"msgs": funcs.Messages(db)})
	})

	r.POST("/homepage", func(c *gin.Context) {
		texts := c.PostForm("texts")
		name, _ := c.Cookie("username")
		if texts != "" && name != "" {
			_, err := db.Exec("insert into messages (name, message, time ) values (?,?,?)", name, texts, string(time.Now().Format("2006-01-02 15:04:05")))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		c.Redirect(http.StatusMovedPermanently, "/homepage")
	})

	//注销
	//将cookie清除并重定向到根目录
	r.GET("/announcement", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/announcement.html", gin.H{})
	})

	r.POST("/announcement", func(c *gin.Context) {
		username, _ := c.Cookie("username")
		c.SetCookie("username", "", -1, "/", "localhost", false, false)
		_, err := db.Exec("delete from  active_users where cookie = ?", username)
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			return
		}
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
	var user struct {
		id       int
		name     string
		password string
	}
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		flag := false

		rows, err := db.Query("select id,name,password from users where name= ?", username)
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			rows.Close()
			return
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&user.id, &user.name, &user.password)
			if err != nil {
				log.Println(err)
				return
			}
			if user.name == username && user.password == password {
				fmt.Println("成功")
				c.SetCookie("username", username, 0, "/", "localhost", false, false)
				_, err := db.Exec("insert into active_users (cookie) values (?)", username)
				if err != nil {
					fmt.Printf("\nerr : %v\n", err)
					return
				}
				flag = true
			}

			if flag {
				//fmt.Printf("flag = %v\n", flag)
				c.Redirect(http.StatusMovedPermanently, "/homepage")
			} else {
				c.Redirect(http.StatusMovedPermanently, "/login")
			}
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
		question := c.PostForm("question")
		answer := c.PostForm("answer")

		if len(password) <= 6 {
			flag = false
		}
		//判断用户名是否已经注册
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/redrocksixth")
		if err != nil {
			fmt.Println(err)
			return
		}

		rows, err := db.Query("select id,name,password from users where name = ?", username)
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			rows.Close()
			return
		}
		defer rows.Close()
		var user struct {
			id       int
			name     string
			password string
		}
		for rows.Next() {
			err := rows.Scan(&user.id, &user.name, &user.password)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(user)
		}
		if user.name == username {
			flag = false
		}
		if flag {
			//注册成功就重定向到login并登录
			funcs.Register(username, password, question, answer, db)
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			//注册不成功就重定向到login
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

	})

	//第一个找回页面
	r.GET("/findback/firststep", checker, func(c *gin.Context) {
		//fmt.Println("\n来到了找回界面\n\n")
		c.HTML(http.StatusOK, "default/findback1.html", gin.H{})
	})

	r.POST("/findback/firststep", func(c *gin.Context) {

		//fmt.Println("\n来到了找回界面2\n\n")
		flag := true
		username := c.PostForm("username")

		//判断用户名是否已经注册
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/redrocksixth")
		if err != nil {
			fmt.Println(err)
			return
		}

		rows, err := db.Query("select id,name,password from users where name = ?", username)
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			rows.Close()
			return
		}
		//fmt.Println("\n来到了找回界面3\n\n")
		defer rows.Close()
		var user struct {
			id       int
			name     string
			password string
		}

		//fmt.Println("\n来到了找回界面4\n\n")
		rows, err = db.Query("select id,name,password from users where name = ?", username)
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			rows.Close()
			return
		}
		for rows.Next() {
			err := rows.Scan(&user.id, &user.name, &user.password)
			if err != nil {
				log.Println(err)
				return
			}
			//fmt.Println(user)
		}
		if user.name == username {
			flag = false
		}
		fmt.Println("\n来到了找回界面5\n\n")
		if !flag {
			//如果已经注册就重定向到secondstep
			Findername = username

			rows, err := db.Query("select id,name,question1 from users where name = ?", Findername)
			if err != nil {
				fmt.Printf("\nerr : %v\n", err)
				rows.Close()
				return
			}

			defer rows.Close()
			var user struct {
				id       int
				name     string
				question string
			}
			for rows.Next() {
				err := rows.Scan(&user.id, &user.name, &user.question)
				if err != nil {
					log.Println(err)
					return
				}
				//fmt.Println(user)
			}
			Question = user.question

			c.Redirect(http.StatusMovedPermanently, "/findback/secondstep")
		} else {
			//如果未注册就重定向到login
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

	})

	r.GET("/findback/secondstep", checker, func(c *gin.Context) {
		//fmt.Println("\n来到了找回界面\n\n")
		c.HTML(http.StatusOK, "default/findback2.html", gin.H{"question": Question})
	})

	r.POST("/findback/secondstep", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/redrocksixth")
		flag := true
		answer := c.PostForm("answer")
		newpassword := c.PostForm("newpassword")
		rows, err := db.Query("select id,name,answer1 from users where answer1 = ?", answer)
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			rows.Close()
			return
		}
		defer rows.Close()
		var user struct {
			id     int
			name   string
			answer string
		}
		for rows.Next() {
			err := rows.Scan(&user.id, &user.name, &user.answer)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(user)
		}
		if user.name == Findername && user.answer == answer {
			flag = false
		}
		if !flag {
			db.Exec("update users set password=? where name=?", newpassword, Findername)
		}
		c.Redirect(http.StatusMovedPermanently, "/login")

	})
	r.Run()
}
