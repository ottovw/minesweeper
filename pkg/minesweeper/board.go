package minesweeper


type Board struct {
	height int
	width int
	Cells [][]Cell
}

type Cell struct {
	Mine bool
	Neighbors int
	Revealed bool
}

func NewBoard(height, width int) *Board {
	board := &Board{
		height: height,
		width: width,
	}
	board.initCells()
	return board
}

func (b *Board) initCells() {
	b.Cells = make([][]Cell, b.height)
	for i := range b.Cells {
		b.Cells[i] = make([]Cell, b.width)
	}
}

func (b *Board) ToString() string {
	return "My Board"
}