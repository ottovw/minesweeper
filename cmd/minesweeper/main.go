package main

import (
	"fmt"

	"github.com/ottovw/minesweeper/pkg/minesweeper"
	"github.com/ottovw/minesweeper/pkg/renderer"
)

func main() {
	fmt.Println("Welcome to Minesweeper! \n---")
	board := minesweeper.NewBoard(5, 10, 10)
	renderer.PrintBoard(board)
	for i := 0; i < 5; i++ {
		ok := board.RevealRandomCell()
		if !ok {
			fmt.Println("Game over!")
			break
		}
		renderer.PrintBoard(board)
	}
}