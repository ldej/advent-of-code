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
	field := tools.ReadIntGrid("./2021/day09/input.txt")

	totalRiskLevel := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			lowest := true
			v := field[i][j]

			if i-1 >= 0 {
				lowest = lowest && v < field[i-1][j]
			}
			if i+1 < len(field) {
				lowest = lowest && v < field[i+1][j]
			}
			if j-1 >= 0 {
				lowest = lowest && v < field[i][j-1]
			}
			if j+1 < len(field[i]) {
				lowest = lowest && v < field[i][j+1]
			}

			if lowest {
				totalRiskLevel += v + 1
			}
		}
	}
	return totalRiskLevel
}
