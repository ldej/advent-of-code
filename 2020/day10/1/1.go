package main

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myints"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day10/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day10/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	ints := tools.ReadInts(input)
	ints = myints.Prepend(append(ints, myints.MaxSlice(ints)+3), 0)

	diff1 := 0
	diff3 := 0

	for _, i := range ints {
		if myints.SliceContains(ints, i+1) {
			diff1 += 1
		} else if myints.SliceContains(ints, i+3) {
			diff3 += 1
		}
	}

	return diff1 * diff3
}
