package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadStrings()

	var score int
	for i := 0; i < len(input); i += 3 {
		letter := mystrings.Intersection(input[i : i+3])[0]
		l := int(letter) - 96 // lowercase
		if l < 0 {
			l += 58 // uppercase
		}
		score += l
	}
	return score
}
