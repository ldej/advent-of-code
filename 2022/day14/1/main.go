package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadStrings(file...)

	source := tools.Point{X: 500, Y: 0}
	points := map[tools.Point]string{
		source: "+",
	}

	minX, maxX := source.X, source.Y
	minY, maxY := source.Y, source.Y

	for _, line := range input {
		var last *tools.Point
		parts := strings.Split(line, " -> ")
		for _, part := range parts {
			numbers := strings.Split(part, ",")
			next := tools.Point{
				X: tools.FindInt(numbers[0]),
				Y: tools.FindInt(numbers[1]),
			}

			minX = myints.Min(minX, next.X)
			maxX = myints.Max(maxX, next.X)
			minY = myints.Min(minY, next.Y)
			maxY = myints.Max(maxY, next.Y)

			if last == nil {
				last = &next
			} else {
				for x := myints.Min(last.X, next.X); x <= myints.Max(last.X, next.X); x++ {
					for y := myints.Min(last.Y, next.Y); y <= myints.Max(last.Y, next.Y); y++ {
						points[tools.Point{X: x, Y: y}] = "X"
					}
				}
				last = &next
			}
		}
	}

	c := cave{
		min: tools.Point{
			X: minX,
			Y: minY,
		},
		max: tools.Point{
			X: maxX,
			Y: maxY,
		},
		source: source,
		points: points,
	}

	result := c.AddSand()

	return result
}

type cave struct {
	min    tools.Point
	max    tools.Point
	source tools.Point
	points map[tools.Point]string
}

func (c *cave) AddSand() int {
	var added int

	for {
		a := c.getPointToAddSand(c.source)
		if a == nil {
			break
		}
		c.points[*a] = "o"
		added++
	}
	c.Print()
	return added
}

func (c *cave) isOutOfBounds(p tools.Point) bool {
	return p.Y > c.max.Y
}

func (c *cave) getPointToAddSand(p tools.Point) *tools.Point {
	oneDown := tools.Point{X: p.X, Y: p.Y + 1}
	downLeft := tools.Point{X: p.X - 1, Y: p.Y + 1}
	downRight := tools.Point{X: p.X + 1, Y: p.Y + 1}

	var next *tools.Point
	if _, found := c.points[oneDown]; !found {
		next = &oneDown
	} else if _, found = c.points[downLeft]; !found {
		next = &downLeft
	} else if _, found = c.points[downRight]; !found {
		next = &downRight
	}
	if next == nil {
		// all three options are blocked
		return &p
	}
	if c.isOutOfBounds(*next) {
		return nil
	}
	return c.getPointToAddSand(*next)
}

func (c *cave) Print() {
	for y := c.min.Y; y <= c.max.Y; y++ {
		for x := c.min.X; x <= c.max.X; x++ {
			if s, found := c.points[tools.Point{X: x, Y: y}]; found {
				fmt.Printf(s)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}
