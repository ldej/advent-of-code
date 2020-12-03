package main

import "fmt"

type Grid [][]int

func main() {
	const input = 361527

	grid, value := findLargerValue(input)
	fmt.Println("Answer:", value)
	grid.print()

}

func findLargerValue(input int) (Grid, int) {
	width := 5
	row := 2
	col := 2
	grid := Grid{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	for {
		grid = upgradeGrid(grid)

		// Right
		for col < width-2 {
			col++
			val := calculateValue(grid, row, col)
			grid[row][col] = val
			if val > input {
				return grid, val
			}
		}
		// Top
		for row > 1 {
			row--
			val := calculateValue(grid, row, col)
			grid[row][col] = val
			if val > input {
				return grid, val
			}
		}
		// Left
		for col > 1 {
			col--
			val := calculateValue(grid, row, col)
			grid[row][col] = val
			if val > input {
				return grid, val
			}
		}
		// Down
		for row < width-2 {
			row++
			val := calculateValue(grid, row, col)
			grid[row][col] = val
			if val > input {
				return grid, val
			}
		}

		width += 2
		row++
		col++
	}
}

func upgradeGrid(grid Grid) Grid {
	width := len(grid)
	newWidth := width + 2
	newGrid := make(Grid, newWidth)
	for i := 0; i < newWidth; i++ {
		newGrid[i] = make([]int, newWidth)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			newGrid[i+1][j+1] = grid[i][j]
		}
	}
	return newGrid
}

func calculateValue(grid Grid, row int, col int) int {
	value := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			value += grid[i][j]
		}
	}
	return value
}

func (grid Grid) print() {
	for i := range grid {
		for j := range grid {
			fmt.Printf("%8d|", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println("\n")
}
