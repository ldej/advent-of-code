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

	var best int
	for cell := range grid.Cells() {
		if count := countVisibleTrees(cell, grid); count > best {
			best = count
		}
	}
	return best
}

func countVisibleTrees(cell tools.IntCell, grid tools.IntGrid) int {
	var right, left, up, down int
	// right
	for i := cell.X + 1; i <= len(grid); i++ {
		if i == len(grid) {
			break
		}
		right++
		if grid[cell.Y][i] >= cell.Value {
			break
		}
	}
	// left
	for i := cell.X - 1; i >= -1; i-- {
		if i == -1 {
			break
		}
		left++
		if grid[cell.Y][i] >= cell.Value {
			break
		}
	}
	// up
	for i := cell.Y - 1; i >= -1; i-- {
		if i == -1 {
			break
		}
		up++
		if grid[i][cell.X] >= cell.Value {
			break
		}
	}
	// down
	for i := cell.Y + 1; i <= len(grid); i++ {
		if i == len(grid) {
			break
		}
		down++
		if grid[i][cell.X] >= cell.Value {
			break
		}
	}
	return right * left * up * down
}
