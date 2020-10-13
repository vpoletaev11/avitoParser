package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/vpoletaev11/avitoParser/scrapper"
	"github.com/vpoletaev11/avitoParser/subscribe"
)

func main() {
	dep := scrapper.NewDep()

	go scrapper.ComparePrices(dep)

	http.HandleFunc("/", subscribe.Handler(dep))

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
