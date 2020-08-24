package main

import (
	"database/sql"
	"fmt"
	"log"
)

func mani()  {
	db, err := sql.Open("web-server", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack")
	if err !=nil {
		log.Fatal(err)
	}
	n, err := result.RowsAffected()
	if n == 1  {
		fmt.Println("1 row inserted.")

	}
}