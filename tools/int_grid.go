package tools

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myints"
	"log"
)

type IntGrid [][]int

type IntWindow struct {
	Grid    IntGrid
	CenterX int
	CenterY int
}

type IntCell struct {
	Value int
	X     int
	Y     int
}

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

func (g IntGrid) Count(value int) int {
	count := 0
	for cell := range g.Cells() {
		if cell.Value == value {
			count += 1
		}
	}
	return count
}

func (g IntGrid) OutOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= len(g) || y >= len(g[0])
}

func (g IntGrid) Copy() IntGrid {
	cp := make(IntGrid, len(g))
	for i := range g {
		cp[i] = make([]int, len(g[i]))
		copy(cp[i], g[i])
	}
	return cp
}

func (g IntGrid) Cells() chan IntCell {
	ch := make(chan IntCell)

	go func() {
		for i := 0; i < len(g); i++ {
			for j := 0; j < len(g[0]); j++ {
				ch <- IntCell{
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

// MaxWindowSum calculates the sum of all cells for a moving window
// it returns the x and y index of the top left corner of the
// window with the highest sum and the sum value
func (g IntGrid) MaxWindowSum(windowHeight int, windowWidth int) (int, int, int) {
	maxSum := 0
	x := 0
	y := 0

	for cell := range g.Cells() {
		if sum := g.WindowSum(windowHeight, windowWidth, cell.X, cell.Y); sum > maxSum {
			maxSum = sum
			x = cell.X
			y = cell.Y
		}
	}
	return x, y, maxSum
}

// WindowSum returns the sum of the integers within the specified window
func (g IntGrid) WindowSum(windowHeight int, windowWidth int, x, y int) int {
	sum := 0

	window := g.Window(windowHeight, windowWidth, x, y)
	height := len(window.Grid)
	width := len(window.Grid[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sum += window.Grid[i][j]
		}
	}
	return sum
}

// Window returns a window of the given height and width
// the window will be clipped by the borders (meaning they can be smaller than the requested window)
// the origin of the window is at x, y
// a square window with odd sides has the origin in the center of the window
func (g IntGrid) Window(windowHeight int, windowWidth int, x, y int) IntWindow {
	if windowHeight == 1 && windowWidth == 1 {
		log.Fatal("use Cells for 1x1 windows")
	}
	oddSquare := windowHeight == windowWidth && windowHeight&1 == 1 && windowWidth&1 == 1

	window := make(IntGrid, 0)

	subtract := 0
	if oddSquare {
		subtract = (windowWidth - 1) / 2
	}

	for i := myints.Max(x-subtract, 0); i < myints.Min(x+windowHeight-subtract, len(g)); i++ {
		min, max := myints.Max(y-subtract, 0), myints.Min(y+windowWidth-subtract, len(g[0]))
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
	return IntWindow{
		Grid:    window,
		CenterX: centerX,
		CenterY: centerY,
	}
}

func (g IntGrid) Windows(windowHeight int, windowWidth int) chan IntWindow {
	ch := make(chan IntWindow)

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
