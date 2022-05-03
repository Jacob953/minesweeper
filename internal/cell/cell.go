package pkg

const (
	Bomb    = Cell(9)
	Opened  = Cell(1 << 4)
	Flagged = Cell(1 << 5)
)

type Cell byte

// same checks the cell with status
func (cell Cell) same(status Cell) bool {
	return cell&status == status
}

func (cell *Cell) open() {
	cell.Flag(false)
	*cell |= Opened
}

func (cell Cell) in(status Cell) bool {
	return cell&status > 0
}

func (cell *Cell) suggest(bombs byte) {
	*cell |= Cell(bombs)
}

// Bombs return bombs the cell contains
func (cell Cell) Bombs() byte {
	return byte(cell &^ (Flagged | Opened))
}

// IsBomb checks whether the cell contains a bomb
func (cell Cell) IsBomb() bool {
	return cell.same(Bomb)
}

// IsUnfolded checks whether the cell is unfolded
func (cell Cell) IsOpened() bool {
	return cell.same(Opened)
}

// IsFlagged checks whether the cell is flagged
func (cell Cell) IsFlagged() bool {
	return cell.same(Flagged)
}

// Flag sets or unsets a flag on the cell
func (cell *Cell) Flag(enabled bool) int {
	if cell.IsOpened() {
		return 0
	}
	if enabled {
		*cell |= Flagged
		return 1
	} else {
		*cell &^= Flagged
		return -1
	}
}

// Open unflags the cell
func (cell *Cell) Open() {
	cell.open()
}

// In checks the cell's status
func (cell *Cell) In(status Cell) bool {
	return cell.in(status)
}

// Suggest set the cell suggested bombs
func (cell *Cell) Suggest(bombs byte) {
	cell.suggest(bombs)
}
