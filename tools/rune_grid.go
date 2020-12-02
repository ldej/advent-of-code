package tools

import (
	"fmt"
)

type RuneGrid [][]rune

func (g RuneGrid) Print() {
	for _, line := range g {
		fmt.Println(string(line))
	}
}

func (g RuneGrid) At(rowIndex, columnIndex int) rune {
	return g[rowIndex][columnIndex]
}

func (g RuneGrid) Set(rowIndex, columnIndex int, char rune) {
	g[rowIndex][columnIndex] = char
}

// GrowAll grows in all directions in one run
func (g RuneGrid) GrowAll(char rune) RuneGrid {
	var newGrid = make(RuneGrid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]rune, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = char
	}

	newGrid = append(append(RuneGrid{emptyRow}, newGrid...), emptyRow)

	for i, row := range newGrid {
		newGrid[i] = append([]rune{char}, row...)
	}

	for i, row := range newGrid {
		newGrid[i] = append(row, char)
	}
	return newGrid
}

// GrowUp copies the grid and adds a row to the top
func (g RuneGrid) GrowUp(char rune) RuneGrid {
	var newGrid = make(RuneGrid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]rune, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = char
	}
	newGrid = append([][]rune{emptyRow}, g...)
	return newGrid
}

// GrowDown copies the grid and adds a row to the bottom
func (g RuneGrid) GrowDown(char rune) RuneGrid {
	var newGrid = make(RuneGrid, len(g), len(g))
	copy(newGrid, g)

	width := len(g[0])

	emptyRow := make([]rune, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = char
	}
	newGrid = append(newGrid, emptyRow)
	return newGrid
}

// GrowLeft copies the grid and adds a column to the left
func (g RuneGrid) GrowLeft(char rune) RuneGrid {
	var newGrid = make(RuneGrid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append([]rune{char}, row...)
	}
	return newGrid
}

// GrowRight copies the grid and adds a column to the right
func (g RuneGrid) GrowRight(char rune) RuneGrid {
	var newGrid = make(RuneGrid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append(row, char)
	}
	return newGrid
}

// Grow functions that use pointers

func (g *RuneGrid) PGrow(char rune) *RuneGrid {
	return g.PGrowUp(char).PGrowDown(char).PGrowLeft(char).PGrowRight(char)
}

func (g *RuneGrid) PGrowUp(char rune) *RuneGrid {
	width := len((*g)[0])

	emptyRow := make([]rune, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = char
	}
	*g = append([][]rune{emptyRow}, *g...)
	return g
}

func (g *RuneGrid) PGrowDown(char rune) *RuneGrid {
	width := len((*g)[0])

	emptyRow := make([]rune, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = char
	}
	*g = append(*g, emptyRow)
	return g
}

func (g *RuneGrid) PGrowLeft(char rune) *RuneGrid {
	for i, row := range *g {
		(*g)[i] = append([]rune{char}, row...)
	}
	return g
}

func (g *RuneGrid) PGrowRight(char rune) *RuneGrid {
	for i, row := range *g {
		(*g)[i] = append(row, char)
	}
	return g
}