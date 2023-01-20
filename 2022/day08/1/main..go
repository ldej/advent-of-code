package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	grid := tools.ReadIntGrid()

	var visible int
	for cell := range grid.Cells() {
		if isVisible(cell, grid) {
			visible++
		}
	}
	return visible
}

func isVisible(cell tools.IntCell, grid tools.IntGrid) bool {
	// right
	for i := cell.X + 1; i <= len(grid); i++ {
		if i == len(grid) {
			return true
		}
		if grid[cell.Y][i] >= cell.Value {
			break
		}
	}
	// left
	for i := cell.X - 1; i >= -1; i-- {
		if i == -1 {
			return true
		}
		if grid[cell.Y][i] >= cell.Value {
			break
		}
	}
	// top
	for i := cell.Y - 1; i >= -1; i-- {
		if i == -1 {
			return true
		}
		if grid[i][cell.X] >= cell.Value {
			break
		}
	}
	// bottom
	for i := cell.Y + 1; i <= len(grid); i++ {
		if i == len(grid) {
			return true
		}
		if grid[i][cell.X] >= cell.Value {
			break
		}
	}
	return false
}
