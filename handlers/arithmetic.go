package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// ArithmeticItem represents a single math question
// Question is the string representation (e.g. "3 + 4")
// Answer is the result for validation on frontend
// Only simple additions/subtractions are generated

type ArithmeticItem struct {
	Question string `json:"question"`
	Answer   int    `json:"answer"`
}

// ArithmeticResponse returns a list of arithmetic questions for the user
// Timeout is suggested time per question (seconds)

type ArithmeticResponse struct {
	Timeout int              `json:"timeout"`
	Items   []ArithmeticItem `json:"items"`
}

// ArithmeticHandler exposes random arithmetic questions
// Example endpoint: /api/math?count=5
func ArithmeticHandler(w http.ResponseWriter, r *http.Request) {
	count := 5
	if c := r.URL.Query().Get("count"); c != "" {
		if v, err := strconv.Atoi(c); err == nil && v > 0 {
			count = v
		}
	}

	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))

	items := make([]ArithmeticItem, count)
	for i := 0; i < count; i++ {
		a := rgen.Intn(10)
		b := rgen.Intn(10)
		if rgen.Intn(2) == 0 {
			items[i] = ArithmeticItem{
				Question: strconv.Itoa(a) + " + " + strconv.Itoa(b),
				Answer:   a + b,
			}
		} else {
			// Ensure non-negative results for subtraction
			if a < b {
				a, b = b, a
			}
			items[i] = ArithmeticItem{
				Question: strconv.Itoa(a) + " - " + strconv.Itoa(b),
				Answer:   a - b,
			}
		}
	}

	resp := ArithmeticResponse{Timeout: 4, Items: items}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
