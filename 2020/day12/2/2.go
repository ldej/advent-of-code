package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
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
		letter, number := line[0], myints.ToInt(line[1:])

		switch letter {
		case 'F':
			positionX += waypointX * number
			positionY += waypointY * number
		case 'E':
			waypointX += number
		case 'S':
			waypointY += number
		case 'W':
			waypointX -= number
		case 'N':
			waypointY -= number
		case 'L':
			waypointX, waypointY = turn(waypointX, waypointY, -1, number)
		case 'R':
			waypointX, waypointY = turn(waypointX, waypointY, 1, number)
		}
	}

	return myints.ManhattanDistance(0, 0, positionX, positionY)
}

func turn(x int, y int, direction int, degrees int) (int, int) {
	switch degrees {
	case 90:
		x, y = -y*direction, x*direction
	case 180:
		x, y = -x, -y
	case 270:
		x, y = y*direction, -x*direction
	}
	return x, y
}
