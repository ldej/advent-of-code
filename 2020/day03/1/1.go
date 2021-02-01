package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools/runegrid"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	grid := runegrid.Read("./2020/day03/input.txt")

	row := 0
	column := 0
	trees := 0

	for {
		column = (column + 3) % len(grid[0])
		row += 1

		if row >= len(grid) {
			break
		}
		if grid.At(row, column) == '#' {
			trees += 1
		}
	}
	return trees
}
