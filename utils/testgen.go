package utils

import (
	"math/rand"
	"time"
)

// GenerateTest crea una griglia visiva in base allo stage cognitivo specificato.
func GenerateTest(stage int) [][]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var pool []string

	switch stage {
	case 1:
		pool = Symbols[:4]
		rows, cols := 2, 2

		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[r.Intn(len(pool))]
			}
		}
		return grid

	case 2:
		pool = Symbols[:3]
		rows, cols := 3, 3

		values := []string{}
		for _, s := range pool {
			for range 3 {
				values = append(values, s)
			}
		}
		r.Shuffle(len(values), func(i, j int) {
			values[i], values[j] = values[j], values[i]
		})

		grid := make([][]string, rows)
		index := 0
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = values[index]
				index++
			}
		}
		return grid

	case 3:
		pool = Symbols[:5]
		rows, cols := 3, 3

		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[r.Intn(len(pool))]
			}
		}
		return grid

	case 4:
		pool = Symbols[:4]
		rows, cols := 4, 4

		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				if i == j {
					grid[i][j] = pool[0] // diagonale principale
				} else if i+j == cols-1 {
					grid[i][j] = pool[1] // diagonale opposta
				} else {
					grid[i][j] = pool[2+r.Intn(len(pool)-2)]
				}
			}
		}
		return grid

	case 5:
		pool = Symbols[2:7]
		rows, cols := 5, 5

		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[(i+j)%len(pool)]
			}
		}
		return grid

	default:
		pool = Symbols
		rows, cols := 3, 3

		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, cols)
			for j := range grid[i] {
				grid[i][j] = pool[r.Intn(len(pool))]
			}
		}
		return grid
	}
}
