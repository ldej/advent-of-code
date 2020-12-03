package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Type string
	Name string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "The subtype used for the grid being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the grid being generated. This should start with a capital letter so that it is exported.")
	flag.Parse()

	t := template.Must(template.New("grid").Parse(gridTemplate))
	t.Execute(os.Stdout, d)
}

var gridTemplate = `
package tools

import (
	"fmt"
)

type {{.Name}}Grid [][]{{.Type}}

func (g {{.Name}}Grid) Print() {
	for _, line := range g {
		fmt.Println(line)
	}
}

func (g {{.Name}}Grid) At(rowIndex, columnIndex int) {{.Type}} {
	return g[rowIndex][columnIndex]
}

func (g {{.Name}}Grid) Set(rowIndex, columnIndex int, value {{.Type}}) {
	g[rowIndex][columnIndex] = value
}

// GrowAll grows in all directions in one run
func (g {{.Name}}Grid) GrowAll(defaultValue {{.Type}}) {{.Name}}Grid {
	var newGrid = make({{.Name}}Grid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]{{.Type}}, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	newGrid = append(append({{.Name}}Grid{emptyRow}, newGrid...), emptyRow)

	for i, row := range newGrid {
		newGrid[i] = append([]{{.Type}}{defaultValue}, row...)
	}

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// GrowUp copies the grid and adds a row to the top
func (g {{.Name}}Grid) GrowUp(defaultValue {{.Type}}) {{.Name}}Grid {
	var newGrid = make({{.Name}}Grid, len(g), len(g))
	copy(newGrid, g)
	width := len(g[0])

	emptyRow := make([]{{.Type}}, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	newGrid = append([][]{{.Type}}{emptyRow}, g...)
	return newGrid
}

// GrowDown copies the grid and adds a row to the bottom
func (g {{.Name}}Grid) GrowDown(defaultValue {{.Type}}) {{.Name}}Grid {
	var newGrid = make({{.Name}}Grid, len(g), len(g))
	copy(newGrid, g)

	width := len(g[0])

	emptyRow := make([]{{.Type}}, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	newGrid = append(newGrid, emptyRow)
	return newGrid
}

// GrowLeft copies the grid and adds a column to the left
func (g {{.Name}}Grid) GrowLeft(defaultValue {{.Type}}) {{.Name}}Grid {
	var newGrid = make({{.Name}}Grid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append([]{{.Type}}{defaultValue}, row...)
	}
	return newGrid
}

// GrowRight copies the grid and adds a column to the right
func (g {{.Name}}Grid) GrowRight(defaultValue {{.Type}}) {{.Name}}Grid {
	var newGrid = make({{.Name}}Grid, len(g), len(g))
	copy(newGrid, g)

	for i, row := range newGrid {
		newGrid[i] = append(row, defaultValue)
	}
	return newGrid
}

// Grow functions that use pointers

func (g *{{.Name}}Grid) PGrow(defaultValue {{.Type}}) *{{.Name}}Grid {
	return g.PGrowUp(defaultValue).PGrowDown(defaultValue).PGrowLeft(defaultValue).PGrowRight(defaultValue)
}

func (g *{{.Name}}Grid) PGrowUp(defaultValue {{.Type}}) *{{.Name}}Grid {
	width := len((*g)[0])

	emptyRow := make([]{{.Type}}, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	*g = append([][]{{.Type}}{emptyRow}, *g...)
	return g
}

func (g *{{.Name}}Grid) PGrowDown(defaultValue {{.Type}}) *{{.Name}}Grid {
	width := len((*g)[0])

	emptyRow := make([]{{.Type}}, width, width)
	for i, _ := range emptyRow {
		emptyRow[i] = defaultValue
	}

	*g = append(*g, emptyRow)
	return g
}

func (g *{{.Name}}Grid) PGrowLeft(defaultValue {{.Type}}) *{{.Name}}Grid {
	for i, row := range *g {
		(*g)[i] = append([]{{.Type}}{defaultValue}, row...)
	}
	return g
}

func (g *{{.Name}}Grid) PGrowRight(defaultValue {{.Type}}) *{{.Name}}Grid {
	for i, row := range *g {
		(*g)[i] = append(row, defaultValue)
	}
	return g
}
`
