package tools

import (
	"fmt"
)

type IntGrid [][]int

func (g IntGrid) Print() {
	for _, line := range g {
		fmt.Println(line)
	}
}

func (g IntGrid) At(rowIndex, columnIndex int) int {
	return g[rowIndex][columnIndex]
}

func (g IntGrid) Set(rowIndex, columnIndex int, value int) {
	g[rowIndex][columnIndex] = value
}

// GrowAll grows in all directions in one run
func (g IntGrid) GrowAll(defaultValue int) IntGrid {
	var newGrid = make(IntGrid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]int, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}

	newGrid = append(append(IntGrid{emptyRow}, newGrid...), emptyRow)

	for i, row := range newGrid {
		newGrid[i] = append([]int{defaultValue}, row...)
	}

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// GrowUp copies the grid and adds a row to the top
func (g IntGrid) GrowUp(defaultValue int) IntGrid {
	var newGrid = make(IntGrid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]int, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	newGrid = append([][]int{emptyRow}, g...)
	return newGrid
}

// GrowDown copies the grid and adds a row to the bottom
func (g IntGrid) GrowDown(defaultValue int) IntGrid {
	var newGrid = make(IntGrid, len(g), len(g))
	copy(newGrid, g)

	width := len(g[0])

	emptyRow := make([]int, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	newGrid = append(newGrid, emptyRow)
	return newGrid
}

// GrowLeft copies the grid and adds a column to the left
func (g IntGrid) GrowLeft(defaultValue int) IntGrid {
	var newGrid = make(IntGrid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append([]int{defaultValue}, row...)
	}
	return newGrid
}

// GrowRight copies the grid and adds a column to the right
func (g IntGrid) GrowRight(defaultValue int) IntGrid {
	var newGrid = make(IntGrid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// Grow functions that use pointers

func (g *IntGrid) PGrow(defaultValue int) *IntGrid {
	return g.PGrowUp(defaultValue).PGrowDown(defaultValue).PGrowLeft(defaultValue).PGrowRight(defaultValue)
}

func (g *IntGrid) PGrowUp(defaultValue int) *IntGrid {
	width := len((*g)[0])

	emptyRow := make([]int, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	*g = append([][]int{emptyRow}, *g...)
	return g
}

func (g *IntGrid) PGrowDown(defaultValue int) *IntGrid {
	width := len((*g)[0])

	emptyRow := make([]int, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	*g = append(*g, emptyRow)
	return g
}

func (g *IntGrid) PGrowLeft(defaultValue int) *IntGrid {
	for i, row := range *g {
		(*g)[i] = append([]int{defaultValue}, row...)
	}
	return g
}

func (g *IntGrid) PGrowRight(defaultValue int) *IntGrid {
	for i, row := range *g {
		(*g)[i] = append(row, defaultValue)
	}
	return g
}