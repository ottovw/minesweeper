package minesweeper

import (
	"math/rand"
)

type Board struct {
	Height    int
	Width     int
	MineCount int
	Cells     [][]Cell
}

type Cell struct {
	Mine      bool
	Neighbors int
	Revealed  bool
}

func NewBoard(height, width, mineCount int) *Board {
	board := &Board{
		Height:    height,
		Width:     width,
		MineCount: mineCount,
	}
	board.initCells()
	board.seedRandomMines()
	board.calculateNeighbors()
	return board
}

func (b *Board) GetCell(x, y int) *Cell {
	return &b.Cells[x][y]
}

func (b *Board) RevealCell(x, y int) bool {
	cell := b.GetCell(x, y)
	cell.Revealed = true
	if cell.Mine {
		return false
	}
	if cell.Neighbors == 0 {
		b.revealNeighbors(x, y)
	}
	return true
}

func (b *Board) RevealRandomCell() bool {
	x, y := b.randomCoords()
	return b.RevealCell(x, y)
}

func (b *Board) initCells() {
	b.Cells = make([][]Cell, b.Height)
	for i := range b.Cells {
		b.Cells[i] = make([]Cell, b.Width)
	}
}

func (b *Board) revealNeighbors(x, y int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			var dx = x + i
			var dy = y + j
			if dx < 0 || dx >= b.Height || dy < 0 || dy >= b.Width {
				continue
			}
			cell := b.GetCell(dx, dy)
			if !cell.Mine && !cell.Revealed {
				cell.Revealed = true
				if cell.Neighbors == 0 {
					b.revealNeighbors(dx, dy)
				}
			}
		}
	}
}

func (b *Board) seedRandomMines() {
	for i := 0; i < b.MineCount; i++ {
		x, y := b.randomCoords()
		cell := b.GetCell(x, y)
		if cell.Mine {
			i--
		} else {
			cell.Mine = true
		}
	}
}

func (b *Board) calculateNeighbors() {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			cell := b.GetCell(i, j)
			if cell.Mine {
				continue
			}
			cell.Neighbors = b.countNeighbors(i, j)
		}
	}
}

func (b *Board) countNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var dx = x + i
			var dy = y + j
			if dx < 0 || dx >= b.Height || dy < 0 || dy >= b.Width {
				continue
			}
			if b.GetCell(dx, dy).Mine {
				count++
			}
		}
	}
	return count
}

func (b *Board) randomCoords() (int, int) {
	return rand.Intn(b.Height), rand.Intn(b.Width)
}
