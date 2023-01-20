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
	rounds := tools.ReadStringSlices()
	var score int

	for _, round := range rounds {
		switch round[0] + round[1] {
		case "AX":
			score += 4
		case "BX":
			score += 1
		case "CX":
			score += 7
		case "AY":
			score += 8
		case "BY":
			score += 5
		case "CY":
			score += 2
		case "AZ":
			score += 3
		case "BZ":
			score += 9
		case "CZ":
			score += 6
		}
	}
	return score
}
