package funcs

import (
	"database/sql"
	"fmt"
)

type Message struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Message string `json:"Message"`
	Time    string `json:"Time"`
}

func Messages(db *(sql.DB)) []Message {
	rows, err := db.Query("select id,name,message,time from messages")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var msgs []Message
	var m Message
	for rows.Next() {
		//fmt.Println("\n检查报错地点\n")
		err = rows.Scan(&m.Id, &m.Name, &m.Message, &m.Time)
		//fmt.Println("\n检查报错地点\n")
		if err != nil {
			fmt.Println(err)
		}
		msgs = append(msgs, m)
	}
	return msgs
}
