package tools

import (
	"fmt"
)

type Int8Grid [][]int8

func (g Int8Grid) Print8() {
	for _, line := range g {
		fmt.Println(line)
	}
}

func (g Int8Grid) At(rowIndex, columnIndex int8) int8 {
	return g[rowIndex][columnIndex]
}

func (g Int8Grid) Set(rowIndex, columnIndex int8, value int8) {
	g[rowIndex][columnIndex] = value
}

// GrowAll grows in all directions in one run
func (g Int8Grid) GrowAll(defaultValue int8) Int8Grid {
	var newGrid = make(Int8Grid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]int8, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}

	newGrid = append(append(Int8Grid{emptyRow}, newGrid...), emptyRow)

	for i, row := range newGrid {
		newGrid[i] = append([]int8{defaultValue}, row...)
	}

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// GrowUp copies the grid and adds a row to the top
func (g Int8Grid) GrowUp(defaultValue int8) Int8Grid {
	var newGrid = make(Int8Grid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]int8, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	newGrid = append([][]int8{emptyRow}, g...)
	return newGrid
}

// GrowDown copies the grid and adds a row to the bottom
func (g Int8Grid) GrowDown(defaultValue int8) Int8Grid {
	var newGrid = make(Int8Grid, len(g), len(g))
	copy(newGrid, g)

	width := len(g[0])

	emptyRow := make([]int8, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	newGrid = append(newGrid, emptyRow)
	return newGrid
}

// GrowLeft copies the grid and adds a column to the left
func (g Int8Grid) GrowLeft(defaultValue int8) Int8Grid {
	var newGrid = make(Int8Grid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append([]int8{defaultValue}, row...)
	}
	return newGrid
}

// GrowRight copies the grid and adds a column to the right
func (g Int8Grid) GrowRight(defaultValue int8) Int8Grid {
	var newGrid = make(Int8Grid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// Grow functions that use point8ers

func (g *Int8Grid) PGrow(defaultValue int8) *Int8Grid {
	return g.PGrowUp(defaultValue).PGrowDown(defaultValue).PGrowLeft(defaultValue).PGrowRight(defaultValue)
}

func (g *Int8Grid) PGrowUp(defaultValue int8) *Int8Grid {
	width := len((*g)[0])

	emptyRow := make([]int8, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	*g = append([][]int8{emptyRow}, *g...)
	return g
}

func (g *Int8Grid) PGrowDown(defaultValue int8) *Int8Grid {
	width := len((*g)[0])

	emptyRow := make([]int8, width, width)
	if defaultValue != 0 {
		for i, _ := range emptyRow {
			emptyRow[i] = defaultValue
		}
	}
	*g = append(*g, emptyRow)
	return g
}

func (g *Int8Grid) PGrowLeft(defaultValue int8) *Int8Grid {
	for i, row := range *g {
		(*g)[i] = append([]int8{defaultValue}, row...)
	}
	return g
}

func (g *Int8Grid) PGrowRight(defaultValue int8) *Int8Grid {
	for i, row := range *g {
		(*g)[i] = append(row, defaultValue)
	}
	return g
}
