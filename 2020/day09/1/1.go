package main

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myints"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day09/example1.txt", 5)
	fmt.Println("Example:", result)

	result = run("./2020/day09/input.txt", 25)
	fmt.Println("Result:", result)
}

func run(input string, lookBack int) int {
	ints := tools.ReadInts(input)

	results := myints.Map(ints, func(i, v int) int {
		if i < lookBack {
			return 1
		}
		for combination := range tools.CombinationsInt(ints[i-lookBack:i], 2) {
			if myints.Sum(combination) == v {
				return 1
			}
		}
		return -1
	})

	return ints[myints.IndexOf(results, -1)]
}
