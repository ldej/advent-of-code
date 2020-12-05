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

	rowMap := map[int][]int{}

	for _, line := range lines {
		row, col, seatID := Seat(line, nrows, ncols)
		rowMap[row] = append(rowMap[row], col)

		seats = append(seats, seatID)
	}

	firstRow := nrows
	lastRow := 0

	for rowNumber, _ := range rowMap {
		if rowNumber < firstRow {
			firstRow = rowNumber
		}
		if rowNumber > lastRow {
			lastRow = rowNumber
		}
	}

	mySeat := 0

	for rowNumber, seatNumbers := range rowMap {
		if len(seatNumbers) != 8 && rowNumber > firstRow && rowNumber < lastRow {
			for i := 0; i < 8; i++ {
				if !tools.IntSliceContains(seatNumbers, i) {
					mySeat = rowNumber*8 + i
				}
			}
		}
	}

	return mySeat
}

func Seat(line string, nrows int, ncols int) (int, int, int) {
	rowLower := 0
	rowUpper := nrows

	colLower := 0
	colUpper := ncols

	for _, c := range line {
		rowDifference := ((rowUpper - rowLower) + 1) / 2
		colDifference := ((colUpper - colLower) + 1) / 2

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

	seatID := rowLower*8 + colLower
	return rowLower, colLower, seatID
}
