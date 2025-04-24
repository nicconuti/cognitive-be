package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// TestSubmission: struttura inviata dal frontend
type TestSubmission struct {
	Stage       int        `json:"stage"`
	Score       int        `json:"score"`
	Total       int        `json:"total"`
	UserGrid    [][]string `json:"userGrid"`
	CorrectGrid [][]string `json:"correctGrid"`
}

// in-memory store per tutte le submissions
var (
	submissions []TestSubmission
	mu          sync.Mutex
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo non consentito", http.StatusMethodNotAllowed)
		return
	}

	var submission TestSubmission
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		http.Error(w, "Dati non validi", http.StatusBadRequest)
		return
	}

	mu.Lock()
	submissions = append(submissions, submission)
	mu.Unlock()

	log.Printf("✅ Risultato salvato – Stage %d: %d/%d corrette", submission.Stage, submission.Score, submission.Total)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submissions)
}
