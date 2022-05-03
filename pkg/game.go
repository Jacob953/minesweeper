package pkg

import (
	. "Jacob953-minesweeper/internal/cell"
	grid "Jacob953-minesweeper/internal/grid"

	"fmt"
)

const (
	BOMB = "x"
	FLAG = "F"
	MASK = "#"
)

type Game struct {
	grid         grid.Grid
	total        int
	set          targetSet
	buf          []target
	rows, cols   int
	flags, lefts int
	opens        int
}

// NewGame init minesweeper games
func NewGame(rows, cols int, diffc float64) (*Game, int) {
	grid, flags := grid.NewGrid(rows, cols, diffc)

	return &Game{
		grid:  grid,
		total: len(grid) * len(grid[0]),
		set:   make(targetSet),
		buf:   make([]target, 0, 8),
		rows:  len(grid) - 1,
		cols:  len(grid[0]) - 1,
		flags: flags,
		lefts: flags,
		opens: 0,
	}, flags
}

// Flag sets or unsets the cell's flag
func (game *Game) Flag(x, y int, enable bool) int {
	game.lefts -= game.grid[x][y].Flag(enable)
	return game.lefts
}

// flagLeft returns left mines
func (game Game) Left() int {
	return game.total - game.opens - game.flags
}

// Open unfolds the mask
// Return left mines and game status
func (game *Game) Open(x, y int) (int, bool) {
	if game.isBomb(x, y) {
		game.grid[x][y].Open()
		return 0, false
	}
	game.unfold(x, y)
	return game.Left(), true
}

type target struct {
	x, y int
}

var tarnil = target{-1, -1}

type targetSet map[target]struct{}

func (set targetSet) add(tar target) {
	set[tar] = struct{}{}
}

func (set targetSet) drop() target {
	for tar := range set {
		delete(set, tar)
		return tar
	}
	return tarnil
}

// unfold
func (game *Game) unfold(x, y int) {
	tar := target{x, y}
	for ; tar != tarnil; tar = game.set.drop() {
		cell := &game.grid[tar.x][tar.y]
		if cell.In(Opened | Flagged) {
			continue
		}
		cell.Open()
		game.opens++
		bombs := game.suggestBombs(tar.x, tar.y)
		if bombs == 0 {
			// unfold cells around the target cell
			for _, cache := range game.buf {
				game.set.add(cache)
			}
		} else {
			cell.Suggest(bombs)
		}
		game.buf = game.buf[:0]
	}
}

// ShowGrid print grid of minesweeper
func (game Game) ShowGrid() {
	for _, row := range game.grid {
		for _, cell := range row {
			if cell.IsOpened() {
				if cell.IsBomb() {
					fmt.Print(BOMB)
				} else {
					fmt.Print(cell.Bombs())
				}
			} else if cell.IsFlagged() {
				fmt.Print(FLAG)
			} else {
				fmt.Print(MASK)
			}
			fmt.Print(" ")
		}
		fmt.Printf("\n")
	}
}

// isBomb checks whether the cell is bomb
func (game Game) isBomb(x, y int) bool {
	return game.grid[x][y].IsBomb()
}

// hasBomb checks whether the cell has a bomb
func (game *Game) hasBomb(x, y int) byte {
	if game.isBomb(x, y) {
		return 1
	}
	// push cells around target to buf
	game.buf = append(game.buf, target{x, y})
	return 0
}

// suggestBombs set suggested bombs of the cell
func (game *Game) suggestBombs(x, y int) (bombs byte) {
	if x > 0 {
		bombs += game.hasBomb(x-1, y)
		if y > 0 {
			bombs += game.hasBomb(x-1, y-1)
		}
		if y < game.cols {
			bombs += game.hasBomb(x-1, y+1)
		}
	}
	if x < game.rows {
		bombs += game.hasBomb(x+1, y)
		if y > 0 {
			bombs += game.hasBomb(x+1, y-1)
		}
		if y < game.cols {
			bombs += game.hasBomb(x+1, y+1)
		}
	}
	if y > 0 {
		bombs += game.hasBomb(x, y-1)
	}
	if y < game.cols {
		bombs += game.hasBomb(x, y+1)
	}
	return
}
