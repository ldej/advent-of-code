package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day24/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day24/input.txt")
	fmt.Println("Result:", result)
}

type Point struct {
	x int
	y int
}

func run(input string) int {
	lines := tools.ReadStrings()

	var tiles = make(map[Point]bool)
	tiles[Point{0, 0}] = true

	for _, line := range lines {
		var p Point
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'e':
				p.x++
			case 'w':
				p.x--
			case 'n':
				p.y--
				p.x += adjust(line[i+1], p.y)
				i++
			case 's':
				p.y++
				p.x += adjust(line[i+1], p.y)
				i++
			}
		}
		if isWhite, found := tiles[p]; found {
			tiles[p] = !isWhite
		} else {
			tiles[p] = false
		}
	}

	// 100 days
	for i := 0; i < 100; i++ {
		var newTiles = make(map[Point]bool)
		var minX, maxX, minY, maxY = findSearchSpace(tiles)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				newTiles[Point{x, y}] = decideColor(tiles, x, y)
			}
		}
		tiles = newTiles
	}

	var black = 0
	for _, isWhite := range tiles {
		if !isWhite {
			black++
		}
	}

	return black
}

// https://www.redblobgames.com/grids/hexagons/
// using the Offset coordinates - "even-r" horizontal layout shoves even rows right
// In hindsight, doubled coordinates would have been easier: https://github.com/mnml/aoc/blob/master/2020/24/2.go
func adjust(dir uint8, y int) int {
	if dir == 'e' && myints.IsOdd(y) {
		return 1
	} else if dir == 'w' && myints.IsEven(y) {
		return -1
	}
	return 0
}

func findSearchSpace(tiles map[Point]bool) (int, int, int, int) {
	var minX, maxX, minY, maxY int
	for point, isWhite := range tiles {
		if isWhite {
			continue
		}
		if point.x < minX {
			minX = point.x
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.y > maxY {
			maxY = point.y
		}
	}
	return minX - 1, maxX + 1, minY - 1, maxY + 1
}

func decideColor(tiles map[Point]bool, x, y int) bool {
	var adjacent []Point
	if myints.IsEven(y) {
		adjacent = []Point{
			{x, y - 1},     // nw
			{x + 1, y - 1}, // ne
			{x + 1, y},     // e
			{x + 1, y + 1}, // se
			{x, y + 1},     // sw
			{x - 1, y},     // w
		}
	} else {
		adjacent = []Point{
			{x - 1, y - 1}, // nw
			{x, y - 1},     // ne
			{x + 1, y},     // e
			{x, y + 1},     // se
			{x - 1, y + 1}, // sw
			{x - 1, y},     // w
		}
	}

	var black int
	for _, point := range adjacent {
		if isWhite, found := tiles[point]; found {
			if !isWhite {
				black++
			}
		}
	}

	isWhite, found := tiles[Point{x, y}]
	if (!found || isWhite) && black == 2 {
		return false
	} else if found && !isWhite && (black == 0 || black > 2) {
		return true
	} else if found {
		return isWhite
	}
	return true
}
