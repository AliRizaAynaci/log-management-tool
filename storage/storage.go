package storage

import (
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Source    string    `json:"source"`
}

// InMemoryStorage, holds the log data in the memory
type InMemoryStorage struct {
	mu   sync.Mutex
	logs []LogEntry
}

func (s *InMemoryStorage) AddLog(entry LogEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.logs = append(s.logs, entry)
}

func (s *InMemoryStorage) GetLogs() []LogEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.logs
}
