package main

import (
	"fmt"

	"github.com/ottovw/minesweeper/pkg/minesweeper"
)

func main() {
	fmt.Println("Hello, World!")
	board := minesweeper.NewBoard(5, 5)
	fmt.Println(board.ToString())
}