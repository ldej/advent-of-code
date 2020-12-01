package common

import (
	"fmt"
)

func Print2D(grid [][]int, print func(int) string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf(print(cell))
		}
		fmt.Printf("\n")
	}

	fmt.Println("\n\nPress ENTER to continue\n")
	var input string
	fmt.Scanln(&input)
}
