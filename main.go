package main

import (
	"log"
	"log-management-tool/handlers"
	"log-management-tool/storage"
	"net/http"
)

func main() {
	store := storage.NewInMemoryStorage()
	handler := &handlers.LogHandler{Storage: store}

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.AddLog(w, r)
		} else if r.Method == http.MethodGet {
			handler.GetLogs(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
