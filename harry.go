package main

import (
	"database/sql"
	"log"
)

func main()  {
	db, err :=sql.Open("web-server","root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
