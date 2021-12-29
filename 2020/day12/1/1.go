package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day12/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day12/input.txt")
	fmt.Println("Result:", result)
}

const (
	East  = 0
	South = 90
	West  = 180
	North = 270
)

func run(input string) int {
	lines := tools.ReadStrings()

	direction := East

	positionX := 0
	positionY := 0

	for _, line := range lines {
		letter, number := line[0], myints.ToInt(line[1:])

		switch letter {
		case 'F':
			switch direction {
			case East:
				positionX += number
			case South:
				positionY += number
			case West:
				positionX -= number
			case North:
				positionY -= number
			}
		case 'E':
			positionX += number
		case 'S':
			positionY += number
		case 'W':
			positionX -= number
		case 'N':
			positionY -= number
		case 'L':
			direction -= number
			if direction < 0 {
				direction += 360
			}
		case 'R':
			direction += number
			direction %= 360
		}
	}

	return myints.ManhattanDistance(0, 0, positionX, positionY)
}
