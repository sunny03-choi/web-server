package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/sunny03-choi/web-server"
	
)

func main()  {
	db, err := sql.Open("web-server","root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err !=nil {
		log.Fatal(err)
	}
	defer db.Close()
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id =1").Scan(&name)
	if err !=nil {
		log.Fatal(err)
		}
		fmt.Println(name)
}
