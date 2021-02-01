package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

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
	lines := tools.ReadStrings(input)

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

	var black int
	for _, isWhite := range tiles {
		if !isWhite {
			black++
		}
	}

	return black
}

// https://www.redblobgames.com/grids/hexagons/
// using the Offset coordinates - "even-r" horizontal layout shoves even rows right
func adjust(dir uint8, y int) int {
	if dir == 'e' && myints.IsOdd(y) {
		return 1
	} else if dir == 'w' && myints.IsEven(y) {
		return -1
	}
	return 0
}
