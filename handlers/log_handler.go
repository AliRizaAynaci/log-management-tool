package handlers

import (
	"encoding/json"
	"log-management-tool/storage"
	"net/http"
	"time"
)

type LogHandler struct {
	Storage *storage.InMemoryStorage
}

func (h *LogHandler) AddLog(w http.ResponseWriter, r *http.Request) {
	var entry storage.LogEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	entry.TimeStamp = time.Now()
	h.Storage.AddLog(entry)
	w.WriteHeader(http.StatusCreated)
}

func (h *LogHandler) GetLogs(w http.ResponseWriter, r *http.Request) {
	logs := h.Storage.GetLogs()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
