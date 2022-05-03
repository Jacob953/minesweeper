package pkg

import (
	. "Jacob953-minesweeper/internal/cell"
	"math"
	"math/rand"
)

type Grid [][]Cell

func initGrid(rows, cols int) (grid Grid) {
	grid = make(Grid, rows)
	for i := range grid {
		grid[i] = make([]Cell, cols)
	}
	return
}

func NewGrid(rows, cols int, diffc float64) (grid Grid, flags int) {
	cells := rows * cols
	bombs := int(math.Ceil(float64(cells) * diffc))
	flags = bombs
	grid = initGrid(rows, cols)
	var offset int
	for bombs > 0 {
		offset += rand.Intn(rand.Intn(cells-offset-bombs) + 1)
		i := offset / cols
		j := offset % cols
		grid[i][j] = Bomb
		bombs--
		offset++
	}
	return
}
