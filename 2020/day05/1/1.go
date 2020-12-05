package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	lines := tools.ReadStrings("./2020/day05/input.txt")
	nrows := 127
	ncols := 8

	seats := []int{}

	for _, line := range lines {
		seats = append(seats, Seat(line, nrows, ncols))
	}

	return tools.MaxList(seats)
}

func Seat(line string, nrows int, ncols int) int {
	rowLower := 1
	rowUpper := nrows

	colLower := 1
	colUpper := ncols

	for _, c := range line {
		rowDifference := (rowUpper - rowLower) / 2
		colDifference := (colUpper - colLower) / 2

		if c == 'F' {
			rowUpper -= rowDifference
		} else if c == 'B' {
			rowLower += rowDifference
		} else if c == 'L' {
			colUpper -= colDifference
		} else if c == 'R' {
			colLower += colDifference
		}
	}

	return rowLower*8 + colLower
}
