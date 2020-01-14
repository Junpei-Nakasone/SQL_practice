package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", "root:@/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Name, &user.Age)
		JSON, err := json.Marshal(user)
		if err != nil {
			continue
		}
		fmt.Println(string(JSON))
	}
}
