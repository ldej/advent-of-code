package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day11/example2.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day11/input.txt")
	fmt.Println("Result:", result)
}

const (
	empty    = 'L'
	occupied = '#'
	floor    = '.'
)

func run(input string) int {
	grid := tools.ReadRuneGrid(input)

	iterate(grid)

	return grid.Count(occupied)
}

func iterate(grid tools.RuneGrid) {
	changed := true

	for changed {
		changed = false

		newGrid := grid.Copy()

		for cell := range grid.Cells() {

			count := CountOccupied(grid, cell.X, cell.Y)

			switch cell.Value {
			case empty:
				if count == 0 {
					newGrid[cell.X][cell.Y] = occupied
					changed = true
				}
			case occupied:
				if count >= 5 {
					newGrid[cell.X][cell.Y] = empty
					changed = true
				}
			case floor:
				// nothing
			}
		}
		copy(grid, newGrid)
	}

	return
}

func CountOccupied(grid tools.RuneGrid, x, y int) int {

	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if findOccupied(grid, x, y, i, j) {
				count++
			}
		}
	}

	return count
}

func findOccupied(grid tools.RuneGrid, x, y, xdir, ydir int) bool {
	if xdir == 0 && ydir == 0 {
		return false
	}

	for i := 1; ; i++ {
		checkX, checkY := x+i*xdir, y+i*ydir

		if grid.OutOfBounds(checkX, checkY) {
			return false
		}
		v := grid[checkX][checkY]

		if v == empty {
			return false
		} else if v == occupied {
			return true
		}
	}
}
