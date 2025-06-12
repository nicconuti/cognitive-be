package utils

import (
	"math/rand"
	"time"
)

// TestType defines the type of test available in the system
// e.g., memory, stroop, arithmetic

const (
	TestMemory     = "memory"
	TestStroop     = "stroop"
	TestArithmetic = "arithmetic"
	TestSequence   = "sequence"
)

// PoolItem represents a single test to run with its type and count
// Count indicates number of items/questions for that test

type PoolItem struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

// GeneratePool creates a randomized list of tests based on difficulty level
// Level 1 includes only memory; higher levels introduce other tests
// The pool length is fixed but content changes with level

func GeneratePool(level int) []PoolItem {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	base := []PoolItem{{Type: TestMemory, Count: 5}}
	if level > 1 {
		base = append(base, PoolItem{Type: TestArithmetic, Count: 5})
	}
	if level > 2 {
		base = append(base, PoolItem{Type: TestStroop, Count: 5})
	}
	if level > 3 {
		base = append(base, PoolItem{Type: TestSequence, Count: 5})
	}

	r.Shuffle(len(base), func(i, j int) { base[i], base[j] = base[j], base[i] })
	return base
}
