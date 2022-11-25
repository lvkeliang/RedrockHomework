package funcs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type information struct {
	name     string
	password string
}

// 定义注册函数
// username:注册名
// password：密码
// db:数据库
func Register(username string, password string, question string, answer string, db *(sql.DB)) {
	var flag int = 0
	if len(password) <= 6 {
		flag = 1
	} else {
		rows, err := db.Query("select id,name,password from users where name= ?", username)
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
			if user.name == username {
				flag = 2
			}
		}
	}

	//fmt.Printf("\n在注册函数里\n")
	if flag == 0 {
		_, err := db.Exec("insert into users (name, password, question1, answer1) values (?,?,?,?)", username, password, question, answer)
		//fmt.Printf("\n在注册函数里2\n")
		if err != nil {
			fmt.Printf("\nerr : %v\n", err)
			return
		}
	} else if flag == 1 {
		fmt.Println("The length of password must be greater than 6")
	} else if flag == 2 {
		fmt.Println("This user name has been registered!")
	}
}
