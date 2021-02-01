package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code/2018/common"
)

func main() {
	results := common.ReadAllLines("./day6/input.txt", `(?P<x>\d+), (?P<y>\d+)`)

	const gridSize = 351

	coordinates := [gridSize][gridSize]int{}
	for i := 0; i < gridSize; i++ {
		coordinates[i] = [gridSize]int{}
		for j := 0; j < gridSize; j++ {
			coordinates[i][j] = 0
		}
	}

	input := [][]int{}
	for _, result := range results {
		x, _ := strconv.Atoi(result["x"])
		y, _ := strconv.Atoi(result["y"])
		if x > 0 && y > 0 {
			input = append(input, []int{x, y})
		}
	}

	for _, c := range input {
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				d := distance(c[0], c[1], i, j)
				coordinates[i][j] += d
			}
		}
	}

	count := 0
	for _, row := range coordinates {
		for _, res := range row {
			//if res < 10000 {
			//	fmt.Print(".")
			//} else {
			//	fmt.Print(" ")
			//}
			if res < 10000 {
				count += 1
			}
		}
		//fmt.Println("")
	}

	fmt.Println(count)
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
