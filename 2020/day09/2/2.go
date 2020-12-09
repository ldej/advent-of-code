package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")
	result := run("./2020/day09/example1.txt", 127)
	fmt.Println("Example:", result)

	result = run("./2020/day09/input.txt", 731031916)
	fmt.Println("Result:", result)
}

func run(input string, expected int) int {
	ints := tools.ReadInts(input)

	result := tools.MapInts(ints, func(i, v int) int {
		if res, contains := containsSum(ints[i:], expected); contains {
			return tools.IntsSumVar(tools.MinAndMax(res))
		}
		return -1
	})

	return tools.IntsNonN(result, 0)
}

func containsSum(ints []int, expected int) ([]int, bool) {
	sum := 0

	for j := 0; sum <= expected; j++ {
		sum = tools.IntsSum(ints[:j])

		if sum == expected {
			return ints[:j], true
		}
	}

	return nil, false
}
