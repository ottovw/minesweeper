package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func ReadCoordinates(maxX, maxY int) (int, int, error) {
	fmt.Print("Next move: x y (?) ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		fmt.Println("")
		return 0, 0, fmt.Errorf("invalid input")
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("")
		return 0, 0, fmt.Errorf("invalid input")
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("")
		return 0, 0, fmt.Errorf("invalid input")
	}
	fmt.Println("")
	return x, y, nil
}