package main

import (
	"log/slog"
	"net/http"

	"b612lpp.hellow/handlers"
	"b612lpp.hellow/infra"
)

func main() {

	logger := infra.NewLogger(infra.NewLogConnector())
	slog.SetDefault(logger)
	q := http.NewServeMux()
	q.HandleFunc("/", handlers.MainPage)
	q.HandleFunc("/calc/", handlers.Calc)
	port := ":8080"
	logger.Info("Server started", "server port", port)
	logger.Error("Error start", "start failed", http.ListenAndServe(port, q))
}
