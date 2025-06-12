package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nicconuti/cognitive-api/utils"
)

// SequenceResponse wraps a list of sequence items with a timeout suggestion.
type SequenceResponse struct {
	Timeout int                  `json:"timeout"`
	Items   []utils.SequenceItem `json:"items"`
}

// SequenceHandler serves numeric sequence reasoning items.
// Example endpoint: /api/sequence?count=5
func SequenceHandler(w http.ResponseWriter, r *http.Request) {
	count := 5
	if c := r.URL.Query().Get("count"); c != "" {
		if v, err := strconv.Atoi(c); err == nil && v > 0 {
			count = v
		}
	}

	items := utils.GenerateSequenceSet(count)
	resp := SequenceResponse{Timeout: 10, Items: items}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
