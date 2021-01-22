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

	total := trees(grid, 1, 1)
	total *= trees(grid, 1, 3)
	total *= trees(grid, 1, 5)
	total *= trees(grid, 1, 7)
	total *= trees(grid, 2, 1)

	return total
}

func trees(grid runegrid.RuneGrid, rowSlope int, columnSlope int) int {
	row := 0
	column := 0
	trees := 0

	for {
		column = (column + columnSlope) % len(grid[0])
		row += rowSlope

		if row >= len(grid) {
			break
		}
		if grid.At(row, column) == '#' {
			trees += 1
		}
	}
	return trees
}
