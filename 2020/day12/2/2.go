package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day12/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day12/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStrings(input)

	positionX := 0
	positionY := 0

	waypointX := 10
	waypointY := -1

	for _, line := range lines {
		letter, number := line[0], tools.ToInt(line[1:])

		switch letter {
		case 'F':
			positionX += waypointX * number
			positionY += waypointY * number
		case 'N':
			waypointY -= number
		case 'S':
			waypointY += number
		case 'E':
			waypointX += number
		case 'W':
			waypointX -= number
		case 'L':
			switch number {
			case 90:
				waypointX, waypointY = waypointY, -waypointX
			case 180:
				waypointX, waypointY = -waypointX, -waypointY
			case 270:
				waypointX, waypointY = -waypointY, waypointX
			}
		case 'R':
			switch number {
			case 90:
				waypointX, waypointY = -waypointY, waypointX
			case 180:
				waypointX, waypointY = -waypointX, -waypointY
			case 270:
				waypointX, waypointY = waypointY, -waypointX
			}
		}
	}

	return tools.ManhattanDistance(0, 0, positionX, positionY)
}
