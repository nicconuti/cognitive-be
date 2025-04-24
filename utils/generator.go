package utils

import (
	"math/rand"
	"time"
)

// GenerateGrid produce una griglia 2D con logica cognitiva crescente a seconda dello stage
func GenerateGrid(stage int) [][]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var rows, cols int
	var pool []string
	var grid [][]string

	switch stage {
	case 1:
		// Stage 1: memoria visiva base → griglia 2x2 con simboli unici
		rows, cols = 2, 2
		pool = Symbols[:4]

		grid = make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[r.Intn(len(pool))]
			}
		}

	case 2:
		// Stage 2: 3 simboli ripetuti → griglia 3x3 bilanciata
		rows, cols = 3, 3
		pool = Symbols[:3]

		values := []string{}
		for _, s := range pool {
			for i := 0; i < 3; i++ {
				values = append(values, s)
			}
		}
		r.Shuffle(len(values), func(i, j int) {
			values[i], values[j] = values[j], values[i]
		})

		grid = make([][]string, rows)
		index := 0
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = values[index]
				index++
			}
		}

	case 3:
		// Stage 3: simboli simili → aumenta confusione visiva
		rows, cols = 3, 3
		pool = Symbols[1:6] // B–F (più simili visivamente)

		grid = make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[r.Intn(len(pool))]
			}
		}

	case 4:
		// Stage 4: introduzione pattern → diagonale fissa, resto random
		rows, cols = 4, 4
		pool = Symbols[1:6]

		grid = make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				if i == j {
					grid[i][j] = pool[0] // Pattern diagonale
				} else {
					grid[i][j] = pool[r.Intn(len(pool))]
				}
			}
		}

	case 5:
		// Stage 5: sequenza mnemonica (crescente in righe)
		rows, cols = 5, 5
		pool = Symbols[2:7]

		grid = make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[(i+j)%len(pool)]
			}
		}

	default:
		// Fallback → 3x3 casuale con simboli completi
		rows, cols = 3, 3
		pool = Symbols

		grid = make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[r.Intn(len(pool))]
			}
		}
	}

	return grid
}
