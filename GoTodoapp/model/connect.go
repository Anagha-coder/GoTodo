package model

import (
	"database/sql"
	"fmt"
)

var con *sql.DB

// 8sql.DB is a return type
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB!")
	con = db

	return db

}
