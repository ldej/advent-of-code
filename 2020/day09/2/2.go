package main

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myints"

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

	result := myints.Map(ints, func(i, v int) int {
		if res := findSliceSum(ints[i:], expected); res != nil {
			return myints.IntsSum(myints.MinAndMax(res))
		}
		return 0
	})

	return myints.IntsNonN(result, 0)
}

func findSliceSum(ints []int, expected int) []int {
	for j := 0; j < len(ints); j++ {
		sum := myints.Sum(ints[:j])

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
			return myints.IntsSum(myints.MinAndMax(ints[s:e]))
		} else if sum < expected {
			e += 1
			sum += ints[e]
		} else if sum > expected {
			sum -= ints[s]
			s += 1
		}
	}
}
