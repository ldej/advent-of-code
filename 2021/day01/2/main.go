package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run1()
	fmt.Println(result)

	result = run2()
	fmt.Println(result)
}

// run1 calculates the sum of the window by adding and subtracting from
// the current sum.
func run1() int {
	ints := tools.ReadInts("./2021/day01/input.txt")
	increased := 0
	windowSize := 3
	start := 0
	current := 0

	for end := range ints {
		last := current
		current += ints[end]
		if end <= windowSize-1 {
			continue
		}

		current -= ints[start]
		start++

		if current > last {
			increased++
		}
	}
	return increased
}

// run2 calculates the sum of the window by summing each window individually
func run2() int {
	ints := tools.ReadInts("./2021/day01/input.txt")
	increased := 0
	windowSize := 3

	last := myints.Sum(ints[0:windowSize])

	for window := range myints.SlidingWindow(ints, windowSize) {
		current := myints.Sum(window)
		if current > last {
			increased++
		}
		last = current
	}
	return increased
}
