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

	r := slidingWindow(ints, expected)
	fmt.Println(r)

	result := tools.MapInts(ints, func(i, v int) int {
		if res := findSliceSum(ints[i:], expected); res != nil {
			return tools.IntsSumVar(tools.MinAndMax(res))
		}
		return 0
	})

	return tools.IntsNonN(result, 0)
}

func findSliceSum(ints []int, expected int) []int {
	for j := 0; j < len(ints); j++ {
		sum := tools.IntsSum(ints[:j])

		if sum == expected {
			return ints[:j]
		} else if sum > expected {
			return nil
		}
	}
	return nil
}

func slidingWindow(ints []int, expected int) int {
	s, e := 0, 1
	sum := ints[s] + ints[e]

	for {
		if sum == expected {
			return tools.IntsSumVar(tools.MinAndMax(ints[s:e]))
		} else if sum < expected {
			e += 1
			sum += ints[e]
		} else if sum > expected {
			sum -= ints[s]
			s += 1
		}
	}
}
