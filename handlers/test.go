package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nicconuti/cognitive-api/utils"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Legge il parametro "stage" dalla query string
	stageStr := r.URL.Query().Get("stage")
	stage, err := strconv.Atoi(stageStr)
	if err != nil || stage < 1 || stage > 5 {
		http.Error(w, "Invalid stage", http.StatusBadRequest)
		return
	}

	// Genera la griglia
	grid := utils.GenerateTest(stage)

	// Imposta il timeout (esempio semplice: 5 secondi fissi)
	resp := TestResponse{
		Grid:    grid,
		Timeout: 5,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
