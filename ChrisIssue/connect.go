package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB { // username:password@tcp(IP_HOST)/NAMADATABASE
	db, error := sql.Open("mysql", "root:root@tcp(127.0.0.1)/sample")

	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println("Connection Success")
	return db
}
