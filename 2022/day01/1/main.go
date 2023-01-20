package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadIntDoubleNewlines()

	var elves []int
	for _, elf := range input {
		elves = append(elves, myints.Sum(elf))
	}

	return myints.Max(elves...)
}
