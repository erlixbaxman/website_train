package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()
	
	router.HandleFunc("GET /{$}", func (w http.ResponseWriter, r *http.Request){
		w.Header().Add("Conten-Type", "text/html; charser utf-8")
		w.Write([]byte("<h1>Hello, World</h1>"))
	})
	
	srv := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	logger.Info("Server listening port :8080")

	srv.ListenAndServe()
}