package main

import (
	"fmt"
)

func main() {

	grid := [300][300]int{}
	gridSerialNumber := 1718

	for j := 0; j < 300; j++ {
		for i := 0; i < 300; i++ {
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

	x, y, _ := getLargestPower(grid, 3)

	fmt.Println(x, y)
}

func getLargestPower(grid [300][300]int, squareSize int) (int, int, int) {
	maxValue := 0
	x := 0
	y := 0

	for j := 0; j < 300-squareSize; j++ {
		for i := 0; i < 300-squareSize; i++ {
			cellSum := 0
			for l := 0; l < squareSize; l++ {
				for k := 0; k < squareSize; k++ {
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
	return x, y, maxValue
}
