package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code/2018/common"
)

func main() {
	results := common.ReadAllLines("./day6/input.txt", `(?P<x>\d+), (?P<y>\d+)`)

	const gridSize = 350

	coordinates := [gridSize][gridSize][]int{}
	for i := 0; i < gridSize; i++ {
		coordinates[i] = [gridSize][]int{}
		for j := 0; j < gridSize; j++ {
			coordinates[i][j] = append(coordinates[i][j], 1000)
			coordinates[i][j] = append(coordinates[i][j], -1)
		}
	}

	input := [][]int{}

	for _, result := range results {
		x, _ := strconv.Atoi(result["x"])
		y, _ := strconv.Atoi(result["y"])
		input = append(input, []int{x, y})
	}

	for resultIndex, c := range input {
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				d := distance(c[0], c[1], i, j)
				if d < coordinates[i][j][0] {
					coordinates[i][j][0] = d
					coordinates[i][j][1] = resultIndex
				} else if d == coordinates[i][j][0] {
					coordinates[i][j][1] = -1
				}
			}
		}
	}

	counts := map[int]int{}
	edge := map[int]bool{}

	for rowIdx, row := range coordinates {
		for colIdx, res := range row {
			if rowIdx == 0 || rowIdx == gridSize-1 || colIdx == 0 || colIdx == gridSize-1 {
				edge[res[1]] = true
			}
		}
	}

	for _, row := range coordinates {
		for _, res := range row {
			if _, foundEdge := edge[res[1]]; !foundEdge {
				if _, found := counts[res[1]]; found {
					counts[res[1]] += 1
				} else {
					counts[res[1]] = 1
				}
			}
		}
	}

	max := 0
	for _, d := range counts {
		if d > max {
			max = d
		}
	}
	fmt.Println(max)
}

func distance(x1, y1, x2, y2 int) int {
	x := 0
	if x1 > x2 {
		x = x1 - x2
	} else {
		x = x2 - x1
	}
	y := 0
	if y1 > y2 {
		y = y1 - y2
	} else {
		y = y2 - y1
	}
	return x + y
}
