package utils

import (
	"math/rand"
	"time"
)

// SequenceItem defines a numeric series and the expected next value.
type SequenceItem struct {
	Series []int `json:"series"`
	Answer int   `json:"answer"`
}

// GenerateSequence produces a single sequence item using common patterns.
// Patterns include arithmetic progression, geometric progression and Fibonacci.
func GenerateSequence() SequenceItem {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pattern := r.Intn(3)

	switch pattern {
	case 0: // arithmetic progression
		start := r.Intn(10)
		diff := r.Intn(5) + 1
		series := make([]int, 4)
		for i := 0; i < 4; i++ {
			series[i] = start + i*diff
		}
		answer := start + 4*diff
		return SequenceItem{Series: series, Answer: answer}
	case 1: // geometric progression
		start := r.Intn(5) + 1
		ratio := r.Intn(3) + 2
		series := make([]int, 4)
		value := start
		for i := 0; i < 4; i++ {
			series[i] = value
			value *= ratio
		}
		answer := value
		return SequenceItem{Series: series, Answer: answer}
	default: // Fibonacci-like
		a := r.Intn(5) + 1
		b := r.Intn(5) + 1
		series := []int{a, b}
		for i := 2; i < 4; i++ {
			series = append(series, series[i-1]+series[i-2])
		}
		answer := series[2] + series[3]
		return SequenceItem{Series: series, Answer: answer}
	}
}

// GenerateSequenceSet creates a list of sequence items.
func GenerateSequenceSet(count int) []SequenceItem {
	items := make([]SequenceItem, count)
	for i := 0; i < count; i++ {
		items[i] = GenerateSequence()
	}
	return items
}
