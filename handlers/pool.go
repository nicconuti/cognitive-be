package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nicconuti/cognitive-api/utils"
)

// PoolHandler returns a randomized pool of tests for a session
// Level parameter influences which tests are included
// Example endpoint: /api/simulate?level=2
func PoolHandler(w http.ResponseWriter, r *http.Request) {
	level := 1
	if l := r.URL.Query().Get("level"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			level = v
		}
	}

	pool := utils.GeneratePool(level)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pool)
}
