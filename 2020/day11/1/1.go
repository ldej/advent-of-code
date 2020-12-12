package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day11/example1.txt")
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
			window := grid.Window(3, 3, cell.X, cell.Y)

			count := CountOccupied(window)

			switch cell.Value {
			case empty:
				if count == 0 {
					newGrid[cell.X][cell.Y] = occupied
					changed = true
				}
			case occupied:
				if count >= 4 {
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

func CountOccupied(window tools.RuneWindow) int {
	count := 0

	for cell := range window.Grid.Cells() {
		if !(cell.X == window.CenterX && cell.Y == window.CenterY) && cell.Value == occupied {
			count += 1
		}
	}
	return count
}
