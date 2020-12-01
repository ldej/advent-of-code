package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/2019/common"
)

func main() {
	scanner := common.ReadLines("./day03/input.txt")

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//line1 := strings.Split("R8,U5,L5,D3", ",")
	//line2 := strings.Split("U7,R6,D4,L4", ",")
	//line1 := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	//line2 := strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	//line1 := strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ",")
	//line2 := strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ",")
	line1 := strings.Split(lines[0], ",")
	line2 := strings.Split(lines[1], ",")

	grid := make([][]int, 1)
	grid[0] = make([]int, 1)
	grid[0][0] = 1

	originX := 0
	originY := 0

	grid, originX, originY = drawLine(grid, line1, originX, originY,11)
	grid, originX, originY = drawLine(grid, line2, originX, originY, 12)

	//common.Print2D(grid, charForInt)

	shortestDistance := 99999999
	for y, row := range grid {
		for x, cell := range row {
			if cell == 5 {
				distance := common.Abs(originX - x) + common.Abs(originY - y)
				if distance < shortestDistance {
					shortestDistance = distance
				}
			}
		}
	}
	fmt.Println(shortestDistance)
}

func charForInt(i int) string {
	switch i {
	case 0:
		return " "
	case 1:
		return "o"
	//case 2:
	//	return "-"
	//case 3:
	//	return "|"
	//case 4:
	//	return "+"
	case 5:
		return "X"
	//case 6:
	//	return "^"

	case 11:
		return "A"
	case 12:
		return "B"
	default:
		log.Fatal(i)
		return "?"
	}
}

func drawLine(grid [][]int, line []string, startX int, startY int, lineNumber int) ([][]int, int, int) {

	originX := startX
	originY := startY

	currentX := startX
	currentY := startY
	for _, item := range line {
		direction, steps := getInstruction(item)

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				currentX += 1
				if currentX >= len(grid[0]) {
					grid = extendGridRight(grid, steps - i)
				}

			case "L":
				currentX -= 1
				if currentX < 0 {
					grid = extendGridLeft(grid, steps - i)
					currentX += steps - i
					originX += steps - i
				}

			case "D":
				currentY += 1
				if currentY >= len(grid) {
					grid = extendGridDown(grid, steps - i)
				}

			case "U":
				currentY -= 1
				if currentY < 0 {
					grid = extendGridUp(grid, steps - i)
					currentY += steps - i
					originY += steps - i
				}
			}

			// it's a line and not itself
			if grid[currentY][currentX] > 1 && grid[currentY][currentX] != lineNumber {
				grid[currentY][currentX] = 5
			} else {
				grid[currentY][currentX] = lineNumber
			}
		}
	}
	return grid, originX, originY
}

func getInstruction(thing string) (string, int)  {
	direction, count := string([]rune(thing)[0]), thing[1:]
	a, _ := strconv.Atoi(count)
	return direction, a
}

func extendGridRight(grid [][]int, count int) [][]int {
	for i, row := range grid {
		grid[i] = append(row, make([]int, count)...)
	}
	return grid
}

func extendGridLeft(grid [][]int, count int) [][]int {
	for i, row := range grid {
		grid[i] = append(make([]int, count), row...)
	}
	return grid
}

func extendGridDown(grid [][]int, count int) [][]int {
	for i := 0; i < count; i++ {
		grid = append(grid, make([]int, len(grid[0])))
	}
	return grid
}

func extendGridUp(grid [][]int, count int) [][]int {
	for i := 0; i < count; i++ {
		grid = append([][]int{make([]int, len(grid[0]))}, grid...)
	}
	return grid
}
