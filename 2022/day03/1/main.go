package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myrunes"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadStrings()

	var score int

	for _, line := range input {
		letter := mystrings.Intersection([]string{line[:len(line)/2], line[len(line)/2:]})[0]
		score += myrunes.ToInt(letter)
	}
	return score
}
