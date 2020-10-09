package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/vpoletaev11/avitoParser/subscribe"
)

const (
	mySQLAddr = "user:@tcp(localhost:3306)"
)

func main() {
	db, err := sql.Open("mysql", mySQLAddr+"/avitoParser")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MySql database")

	http.HandleFunc("/", subscribe.Handler(db))

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
