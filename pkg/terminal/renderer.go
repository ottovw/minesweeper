package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ottovw/minesweeper/pkg/minesweeper"
)


func RenderBoard(board *minesweeper.Board) string {
	var sb strings.Builder
	for i := 0; i < board.Height; i++ {
		if i == 0 {
			sb.WriteString("    ")
			for j := 0; j < board.Width; j++ {
				sb.WriteString(fmt.Sprintf("%d ", j))
			}
			sb.WriteString("\n  ")
			sb.WriteString(strings.Repeat("-", board.Width*2+1))
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%d | ", i))
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