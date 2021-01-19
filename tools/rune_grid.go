package tools

import (
	"fmt"
	"log"
)

type RuneGrid [][]rune

type RuneWindow struct {
	Grid    RuneGrid
	CenterX int
	CenterY int
}

type RuneCell struct {
	Value rune
	X     int
	Y     int
}

func (g RuneGrid) Print() {
	for _, line := range g {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func (g RuneGrid) At(rowIndex, columnIndex int) rune {
	return g[rowIndex][columnIndex]
}

func (g RuneGrid) Set(rowIndex, columnIndex int, char rune) {
	g[rowIndex][columnIndex] = char
}

func (g RuneGrid) Count(char rune) int {
	count := 0
	for cell := range g.Cells() {
		if cell.Value == char {
			count += 1
		}
	}
	return count
}

func (g RuneGrid) OutOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= len(g) || y >= len(g[0])
}

func (g RuneGrid) Copy() RuneGrid {
	cp := make(RuneGrid, len(g))
	for i := range g {
		cp[i] = make([]rune, len(g[i]))
		copy(cp[i], g[i])
	}
	return cp
}

func (g RuneGrid) Cells() chan RuneCell {
	ch := make(chan RuneCell)

	go func() {
		for i := 0; i < len(g); i++ {
			for j := 0; j < len(g[0]); j++ {
				ch <- RuneCell{
					Value: g[i][j],
					X:     i,
					Y:     j,
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g RuneGrid) Window(windowHeight int, windowWidth int, x, y int) RuneWindow {
	if windowHeight == 1 && windowWidth == 1 {
		log.Fatal("use Cells for 1x1 windows")
	}
	oddSquare := windowHeight == windowWidth && windowHeight&1 == 1 && windowWidth&1 == 1

	window := make(RuneGrid, 0)

	subtract := 0
	if oddSquare {
		subtract = (windowWidth - 1) / 2
	}

	for i := Max(x-subtract, 0); i < Min(x+windowHeight-subtract, len(g)); i++ {
		min, max := Max(y-subtract, 0), Min(y+windowWidth-subtract, len(g[0]))
		window = append(window, g[i][min:max])
	}

	var centerX, centerY int

	if oddSquare {
		// TODO clipped origin and template
		centerX = (windowHeight - 1) / 2
		if x == 0 {
			centerX = 0
		}
		centerY = (windowWidth - 1) / 2
		if y == 0 {
			centerY = 0
		}
	}
	return RuneWindow{
		Grid:    window,
		CenterX: centerX,
		CenterY: centerY,
	}
}

func (g RuneGrid) Windows(windowHeight int, windowWidth int) chan RuneWindow {
	ch := make(chan RuneWindow)

	go func() {
		for i := 0; i <= len(g)-windowHeight; i++ {
			for j := 0; j <= len(g[0])-windowWidth; j++ {
				ch <- g.Window(windowHeight, windowWidth, i, j)
			}
		}
		close(ch)
	}()
	return ch
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

func (g RuneGrid) Transpose() RuneGrid {
	var newGrid = make(RuneGrid, 0)

	var height = len(g)
	var width = len(g[0])

	for i := 0; i < width; i++ {
		var row []rune
		for j := 0; j < height; j++ {
			row = append(row, g[j][i])
		}
		newGrid = append(newGrid, row)
	}

	return newGrid
}

// Rotate in a clockwise direction
func (g RuneGrid) Rotate(degrees int) RuneGrid {
	switch degrees {
	case 90, -270:
		return g.Transpose().FlipHorizontal()
	case -90, 270:
		return g.Transpose().FlipVertical()
	case 180, -180:
		return g.FlipHorizontal().FlipVertical()
	case 360:
		return g
	default:
		log.Fatal("Unsupported degrees")
		return nil
	}
}

func (g RuneGrid) FlipVertical() RuneGrid {
	var newGrid = make(RuneGrid, 0)

	for i := len(g) - 1; i >= 0; i-- {
		newGrid = append(newGrid, g[i])
	}

	return newGrid
}

func (g RuneGrid) FlipHorizontal() RuneGrid {
	var newGrid = make(RuneGrid, 0)

	for i := 0; i < len(g); i++ {
		var row []rune
		for j := len(g[0]) - 1; j >= 0; j-- {
			row = append(row, g[i][j])
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}
