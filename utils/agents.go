package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Agent represents a simulated comparison subject with a normalized score.
type Agent struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

// GenerateAgents creates n synthetic agents with scores following
// a truncated normal distribution centered around typical test performance.
// Scores are clipped to [0,1] and agent names are mostly unique.
func GenerateAgents(n int) []Agent {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	names := []string{"Alex", "Jordan", "Casey", "Riley", "Morgan", "Taylor", "Chris", "Sam", "Jamie", "Drew", "Cameron", "Quinn", "Peyton", "Skyler", "Avery", "Hayden", "Charlie", "Dakota", "Emerson", "Harper"}
	r.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	agents := make([]Agent, n)
	for i := 0; i < n; i++ {
		var name string
		if i < len(names) {
			name = names[i]
		} else {
			name = fmt.Sprintf("Agent%02d", i-len(names)+1)
		}

		score := r.NormFloat64()*0.12 + 0.65
		if score < 0 {
			score = 0
		}
		if score > 1 {
			score = 1
		}

		agents[i] = Agent{Name: name, Score: score}
	}
	return agents
}

// EstimateIQ returns an approximate IQ score for a given normalized test score
// by comparing it with a population of agents. The IQ scale has mean 100 and
// standard deviation 15.
func EstimateIQ(userScore float64, agents []Agent) int {
	if len(agents) == 0 {
		return 100
	}
	if userScore < 0 {
		userScore = 0
	}
	if userScore > 1 {
		userScore = 1
	}
	// compute mean and standard deviation of agent scores
	mean := 0.0
	for _, a := range agents {
		mean += a.Score
	}
	mean /= float64(len(agents))
	variance := 0.0
	for _, a := range agents {
		diff := a.Score - mean
		variance += diff * diff
	}
	sd := math.Sqrt(variance / float64(len(agents)))
	if sd == 0 {
		sd = 1
	}
	z := (userScore - mean) / sd
	iq := int(math.Round(100 + z*15))
	if iq < 40 {
		iq = 40
	}
	if iq > 160 {
		iq = 160
	}
	return iq
}
