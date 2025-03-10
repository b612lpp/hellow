package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	slog.Info("Main page was called")
}
