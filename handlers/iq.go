package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nicconuti/cognitive-api/utils"
)

// IQRequest represents the incoming data for IQ estimation.
type IQRequest struct {
	Score int `json:"score"`
	Total int `json:"total"`
}

// IQResponse returns the estimated IQ value and the simulated agents used.
type IQResponse struct {
	IQ     int           `json:"iq"`
	Agents []utils.Agent `json:"agents"`
}

// IQHandler estimates an IQ-like value from a raw score.
func IQHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo non consentito", http.StatusMethodNotAllowed)
		return
	}

	var req IQRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dati non validi", http.StatusBadRequest)
		return
	}

	if req.Total <= 0 || req.Score < 0 || req.Score > req.Total {
		http.Error(w, "Dati non validi", http.StatusBadRequest)
		return
	}

	agents := utils.GenerateAgents(30)
	normalized := float64(req.Score) / float64(req.Total)
	iq := utils.EstimateIQ(normalized, agents)

	resp := IQResponse{IQ: iq, Agents: agents}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
