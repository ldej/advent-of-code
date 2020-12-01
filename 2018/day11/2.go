package main

import (
	"fmt"
	"strings"
)

type result struct {
	X int
	Y int
	Value int
	Size int
}

const (
	gridSerialNumber = 1718
	maxSquareSize = 150
	gridSize = 300
)


func main() {
	channel := make(chan result)

	for size := 1; size <= maxSquareSize; size++ {
		go getLargestPower(channel, size)
	}

	results := []result{}

	for len(results) < maxSquareSize {
		result, _ := <-channel
		results = append(results, result)

		dots := []string{}
		for i := 0; i < maxSquareSize; i++ {
			if i < len(results) {
				dots = append(dots, "#")
			} else {
				dots = append(dots, ".")
			}
		}
		fmt.Printf("\r" + strings.Join(dots, ""))
	}
	close(channel)

	best := result{}
	for _, result := range results {
		if result.Value > best.Value {
			best = result
		}
	}

	fmt.Println()
	fmt.Println(best.X, best.Y, best.Size)
}

func getLargestPower(channel chan result, squareSize int) {
	grid := makeGrid()

	maxValue := 0
	x := 0
	y := 0

	for j := 0; j < gridSize - squareSize; j++ {
		for i := 0; i < gridSize - squareSize; i++ {
			cellSum := 0
			for l := 0; l < squareSize; l ++ {
				for k := 0; k < squareSize; k ++ {
					cellSum += grid[j+l][i+k]
				}
			}
			if cellSum > maxValue {
				maxValue = cellSum
				x = i + 1
				y = j + 1
			}
		}
	}
	channel <- result{
		X: x,
		Y: y,
		Value: maxValue,
		Size: squareSize,
	}
}

func makeGrid() [gridSize][gridSize]int {
	grid := [gridSize][gridSize]int{}

	for j := 0; j < gridSize; j++ {
		for i := 0; i < gridSize; i++ {
			x := i + 1
			y := j + 1

			rackID := x + 10
			powerLevelStart := ((rackID * y) + gridSerialNumber) * rackID
			powerLevel := 0
			if powerLevelStart > 100 {
				powerLevel = (powerLevelStart / 100) % 10
			}
			grid[j][i] = powerLevel - 5
		}
	}
	return grid
}
