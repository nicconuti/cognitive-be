package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// StroopItem represents a single stroop word with a display color
// Text is the word shown, Color is the color in which the word is displayed
// The user must name the color ignoring the text
// Example: Text="RED", Color="blue" -> correct answer is "blue"
type StroopItem struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

// StroopResponse contains a list of Stroop items for the test
// Timeout defines how long each item should be displayed in seconds
// Items contains the words with their display color
// Items are randomized each call
// Colors are chosen from a limited set

type StroopResponse struct {
	Timeout int          `json:"timeout"`
	Items   []StroopItem `json:"items"`
}

// StroopHandler serves a simple Stroop test
// Example endpoint: /api/stroop?count=5
func StroopHandler(w http.ResponseWriter, r *http.Request) {
	count := 5
	if c := r.URL.Query().Get("count"); c != "" {
		if v, err := strconv.Atoi(c); err == nil && v > 0 {
			count = v
		}
	}

	colors := []string{"red", "blue", "green", "yellow"}
	words := []string{"RED", "BLUE", "GREEN", "YELLOW"}

	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))

	items := make([]StroopItem, count)
	for i := 0; i < count; i++ {
		text := words[rgen.Intn(len(words))]
		color := colors[rgen.Intn(len(colors))]
		items[i] = StroopItem{Text: text, Color: color}
	}

	resp := StroopResponse{Timeout: 3, Items: items}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
