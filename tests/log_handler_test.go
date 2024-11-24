package tests

import (
	"bytes"
	"encoding/json"
	"log-management-tool/handlers"
	"log-management-tool/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddLog(t *testing.T) {
	store := storage.NewInMemoryStorage()
	handler := &handlers.LogHandler{Storage: store}

	// Yeni bir log verisi hazırlayalım
	logData := storage.LogEntry{
		Level:   "INFO",
		Message: "User logged in",
		Source:  "auth-service",
	}

	// Log ekleme isteği
	body, _ := json.Marshal(logData)
	req := httptest.NewRequest("POST", "/logs", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.AddLog(w, req)

	// Durum kodu kontrolü
	if w.Code != http.StatusCreated {
		t.Errorf("Beklenen durum kodu %d, ama %d geldi", http.StatusCreated, w.Code)
	}
}

func TestGetLogs(t *testing.T) {
	store := storage.NewInMemoryStorage()
	handler := &handlers.LogHandler{Storage: store}

	// Test için log ekleyelim
	logData := storage.LogEntry{
		Level:   "INFO",
		Message: "User logged in",
		Source:  "auth-service",
	}
	store.AddLog(logData)

	// Logları alma isteği
	req := httptest.NewRequest("GET", "/logs", nil)
	w := httptest.NewRecorder()
	handler.GetLogs(w, req)

	// Durum kodu kontrolü
	if w.Code != http.StatusOK {
		t.Errorf("Beklenen durum kodu %d, ama %d geldi", http.StatusOK, w.Code)
	}

	// Log verilerini kontrol et
	var logs []storage.LogEntry
	if err := json.NewDecoder(w.Body).Decode(&logs); err != nil {
		t.Errorf("Log verileri doğru şekilde decode edilemedi: %v", err)
	}

	if len(logs) == 0 {
		t.Error("Log verisi bekleniyor, ama gelmedi")
	}
}
