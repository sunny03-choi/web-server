package main
import (
	"database/sql"
	"fmt"
	_"github.com/sunny03-choi/web-server"
	"log"
)

func main()  {
	db, err :=sql.Open("web-server","root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err !=nil{
		log.Fatal(err)
	}
	defer db.Close()
	var id int
	var name string
	rows, err := db.Query("SELECT id, name FROM test1 where id >= ?", 1)
	if err != nil {
		log.Fatal(err)

	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(id,name)
	}

}
