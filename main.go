package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/vpoletaev11/avitoParser/scrapper"
	"github.com/vpoletaev11/avitoParser/subscribe"
)

const (
	mySQLAddr = "root:@tcp(mysql:3306)"
	dbName    = "avitoParser"
)

func main() {
	db, err := sql.Open("mysql", mySQLAddr+"/"+dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MySql database")

	go scrapper.ComparePrices(db)

	http.HandleFunc("/", subscribe.Handler(db))

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
