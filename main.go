package main

import (
	"log"
	"net/http"

	"b612lpp.hellow/handlers"
)

func main() {
	q := http.NewServeMux()
	q.HandleFunc("/", handlers.MainPage)
	q.HandleFunc("/calc/", handlers.Calc)
	log.Fatal(http.ListenAndServe(":8080", q))
}
