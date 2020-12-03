package tools

import (
	"fmt"
)

type IntSliceGrid [][][]int

func (g IntSliceGrid) Print() {
	for _, line := range g {
		fmt.Println(line)
	}
}

func (g IntSliceGrid) At(rowIndex, columnIndex int) []int {
	return g[rowIndex][columnIndex]
}

func (g IntSliceGrid) Set(rowIndex, columnIndex int, value []int) {
	g[rowIndex][columnIndex] = value
}

// GrowAll grows in all directions in one run
func (g IntSliceGrid) GrowAll(defaultValue []int) IntSliceGrid {
	var newGrid = make(IntSliceGrid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([][]int, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	newGrid = append(append(IntSliceGrid{emptyRow}, newGrid...), emptyRow)

	for i, row := range newGrid {
		newGrid[i] = append([][]int{defaultValue}, row...)
	}

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// GrowUp copies the grid and adds a row to the top
func (g IntSliceGrid) GrowUp(defaultValue []int) IntSliceGrid {
	var newGrid = make(IntSliceGrid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([][]int, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	newGrid = append([][][]int{emptyRow}, g...)
	return newGrid
}

// GrowDown copies the grid and adds a row to the bottom
func (g IntSliceGrid) GrowDown(defaultValue []int) IntSliceGrid {
	var newGrid = make(IntSliceGrid, len(g), len(g))
	copy(newGrid, g)

	width := len(g[0])

	emptyRow := make([][]int, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	newGrid = append(newGrid, emptyRow)
	return newGrid
}

// GrowLeft copies the grid and adds a column to the left
func (g IntSliceGrid) GrowLeft(defaultValue []int) IntSliceGrid {
	var newGrid = make(IntSliceGrid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append([][]int{defaultValue}, row...)
	}
	return newGrid
}

// GrowRight copies the grid and adds a column to the right
func (g IntSliceGrid) GrowRight(defaultValue []int) IntSliceGrid {
	var newGrid = make(IntSliceGrid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// Grow functions that use pointers

func (g *IntSliceGrid) PGrow(defaultValue []int) *IntSliceGrid {
	return g.PGrowUp(defaultValue).PGrowDown(defaultValue).PGrowLeft(defaultValue).PGrowRight(defaultValue)
}

func (g *IntSliceGrid) PGrowUp(defaultValue []int) *IntSliceGrid {
	width := len((*g)[0])

	emptyRow := make([][]int, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	*g = append([][][]int{emptyRow}, *g...)
	return g
}

func (g *IntSliceGrid) PGrowDown(defaultValue []int) *IntSliceGrid {
	width := len((*g)[0])

	emptyRow := make([][]int, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	*g = append(*g, emptyRow)
	return g
}

func (g *IntSliceGrid) PGrowLeft(defaultValue []int) *IntSliceGrid {
	for i, row := range *g {
		(*g)[i] = append([][]int{defaultValue}, row...)
	}
	return g
}

func (g *IntSliceGrid) PGrowRight(defaultValue []int) *IntSliceGrid {
	for i, row := range *g {
		(*g)[i] = append(row, defaultValue)
	}
	return g
}
