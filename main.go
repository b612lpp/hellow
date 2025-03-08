package main

import (
	"fmt"
	"net/http"
)

func main() {
	q := http.NewServeMux()
	q.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", q)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
