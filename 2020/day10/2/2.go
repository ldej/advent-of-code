package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day10/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day10/example2.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day10/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	ints := tools.ReadInts()

	possibleValues := make(map[int][]int)
	possibleValues[0] = []int{1, 2, 3}

	for _, value := range ints {
		possibleValues[value] = []int{value + 1, value + 2, value + 3}
	}

	result := calculate(possibleValues, make(map[int]int), myints.MaxSlice(ints)+3, 0)

	return result
}

// Climbing stairwell problem

func calculate(possibleValues map[int][]int, result map[int]int, target int, index int) int {
	if value, found := result[index]; found {
		return value
	}

	counter := 0

	for _, current := range possibleValues[index] {
		if current != target {
			counter += calculate(possibleValues, result, target, current)
			continue
		}
		counter += 1
	}

	result[index] = counter
	return counter
}
