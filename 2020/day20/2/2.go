package main

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myrunes"
	"github.com/ldej/advent-of-code/tools/mystrings"
	"github.com/ldej/advent-of-code/tools/runegrid"
	"math"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day20/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day20/input.txt")
	fmt.Println("Result:", result)
}

type Tile struct {
	ID   int
	Grid runegrid.RuneGrid
}

type Tiles []Tile

func run(input string) int {
	parts := tools.ReadStringsDoubleNewlines(input)

	var tiles = make(Tiles, len(parts))
	for i, part := range parts {
		lines := strings.Split(part, "\n")
		tiles[i] = Tile{ID: tools.FindInt(lines[0]), Grid: runegrid.FromStrings(lines[1:])}
	}

	image := CreateImage(tiles)

	var seaMonster = runegrid.RuneGrid{
		[]rune("                  # "),
		[]rune("#    ##    ##    ###"),
		[]rune(" #  #  #  #  #  #   "),
	}

	finalImage := DetectSeaMonsters(image, seaMonster)
	finalImage.Print()
	return finalImage.Count('#')
}

func CreateImage(tiles Tiles) runegrid.RuneGrid {
	var size = int(math.Sqrt(float64(len(tiles))))
	var countedEdges = tiles.CountEdges()

	var firstTile, _ = FindCorner(tiles, countedEdges).FindOrientation(nil, nil, countedEdges)
	var usedTiles = map[int]bool{
		firstTile.ID: true,
	}

	// edges to detect if the tile fits
	var left = firstTile.Grid.RightEdge()
	var top []rune

	var finalTiles = make([]Tiles, size)
	var row = []Tile{firstTile}

	// find a tile for each location
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for _, tile := range tiles {
				if usedTiles[tile.ID] {
					continue
				}
				if orientation, found := tile.FindOrientation(left, top, countedEdges); found {
					row = append(row, orientation)
					usedTiles[orientation.ID] = true

					left = orientation.Grid.RightEdge()
					if x > 0 && y < size-1 {
						top = finalTiles[x-1][y+1].Grid.BottomEdge() // Bottom of next tile of the row above
					}
					break
				}
			}
		}
		// end of the row
		left, top = nil, row[0].Grid.BottomEdge()
		finalTiles[x] = row
		row = Tiles{}
	}

	// Found all tiles in order with right orientation, now stitch them together

	var imageSize = (len(tiles[0].Grid) - 2) * size
	var grid = runegrid.NewRuneGrid(imageSize, imageSize)

	for i, row := range finalTiles {
		for j, tile := range row {
			t := tile.Grid.WithoutEdges()

			for ii, k := range t {
				for jj, l := range k {
					x := i*(len(tiles[0].Grid)-2) + ii
					y := j*(len(tiles[0].Grid)-2) + jj
					grid[x][y] = l
				}
			}
		}
	}
	return grid
}

func DetectSeaMonsters(image runegrid.RuneGrid, seaMonster runegrid.RuneGrid) runegrid.RuneGrid {
	var coordinates [][]int

	for _, orientation := range image.Orientations() {
		coordinates = FindSeaMonsters(orientation, seaMonster)
		if len(coordinates) > 0 {
			image = orientation
			break
		}
	}
	return ReplaceSeaMonsters(image, coordinates, seaMonster)
}

func FindSeaMonsters(image runegrid.RuneGrid, seaMonster runegrid.RuneGrid) [][]int {
	height := len(seaMonster)
	width := len(seaMonster[0])
	var coordinates [][]int

	for window := range image.Windows(height, width) {
		if HasSeaMonster(seaMonster, window) {
			coordinates = append(coordinates, []int{window.X, window.Y})
		}
	}
	return coordinates
}

func HasSeaMonster(toFind runegrid.RuneGrid, window runegrid.RuneWindow) bool {
	for x, row := range toFind {
		for y, cell := range row {
			if cell == '#' && window.Grid[x][y] != '#' {
				return false
			}
		}
	}
	return true
}

func ReplaceSeaMonsters(image runegrid.RuneGrid, coordinates [][]int, seaMonster runegrid.RuneGrid) runegrid.RuneGrid {
	for _, coordinate := range coordinates {
		x, y := coordinate[0], coordinate[1]
		for x1, row := range seaMonster {
			for y1, cell := range row {
				if cell == '#' {
					image[x+x1][y+y1] = 'O'
				}
			}
		}
	}
	return image
}

func FindCorner(tiles Tiles, countedEdges map[string]int) Tile {
	for _, tile := range tiles {
		var uniqueEdges int
		for _, edge := range tile.AllEdges() {
			if countedEdges[edge] == 1 {
				uniqueEdges++
			}
		}
		// Corner pieces have two unique edges, but we have also reversed them, so double
		if uniqueEdges == 4 {
			return tile
		}
	}
	panic("no corners")
}

func (tiles Tiles) AllEdges() []string {
	var edges []string
	for _, tile := range tiles {
		edges = append(edges, tile.AllEdges()...)
	}
	return edges
}

func (tiles Tiles) CountEdges() map[string]int {
	var countedEdges = make(map[string]int)
	for _, edge := range tiles.AllEdges() {
		countedEdges[edge]++
	}
	return countedEdges
}

func (tile Tile) AllEdges() []string {
	var edges []string
	for _, edge := range tile.Grid.Edges() {
		edges = append(edges, string(edge), mystrings.Reverse(string(edge)))
	}
	return edges
}

func (tile Tile) FindOrientation(left, top []rune, edges map[string]int) (Tile, bool) {
	for _, orientation := range tile.Grid.Orientations() {
		t, l := orientation.TopEdge(), orientation.LeftEdge()
		matchesLeft := (left == nil && edges[string(l)] == 1) || myrunes.Equal(left, l)
		matchesTop := (top == nil && edges[string(t)] == 1) || myrunes.Equal(top, t)
		if matchesLeft && matchesTop {
			tile.Grid = orientation
			return tile, true
		}
	}
	return tile, false
}
