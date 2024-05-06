package main

import (
	"fmt"
	"os"

	"github.com/ottovw/minesweeper/pkg/minesweeper"
	"github.com/ottovw/minesweeper/pkg/terminal"
)

func main() {
	fmt.Println("Welcome to Minesweeper! \n---")
	board := minesweeper.NewBoard(5, 10, 10)
	terminal.PrintBoard(board)
	for {
		x, y, err := terminal.ReadCoordinates(board.Height, board.Width)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ok := board.RevealCell(x,y)
		if !ok {
			terminal.PrintBoard(board)
			fmt.Println("Game over!")
			os.Exit(1)
		}
		terminal.PrintBoard(board)
	}
}