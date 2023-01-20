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
			score += 3
		case "BX":
			score += 1
		case "CX":
			score += 2
		case "AY":
			score += 4
		case "BY":
			score += 5
		case "CY":
			score += 6
		case "AZ":
			score += 8
		case "BZ":
			score += 9
		case "CZ":
			score += 7
		}
	}
	return score
}
