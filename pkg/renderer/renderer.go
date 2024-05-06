package renderer

import (
	"strconv"
	"strings"

	"github.com/ottovw/minesweeper/pkg/minesweeper"
)


func RenderBoard(board *minesweeper.Board) string {
	var sb strings.Builder
	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			cell := board.Cells[i][j]
			sb.WriteString(RenderCell(&cell)+" ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func RenderCell(cell *minesweeper.Cell) string {
	if cell.Revealed {
		if cell.Mine {
			return "X"
		} else {
			return strconv.Itoa(cell.Neighbors)
		}
	} else {
		return "."
	}
}

func PrintBoard(board *minesweeper.Board) {
	println(RenderBoard(board))
}