package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

	result := run("0,3,6", 2020)
	fmt.Println("Example:", result)

	result = run("6,19,0,5,7,13,1", 2020)
	fmt.Println("Result:", result)
}

func run(input string, limit int) int {
	values := myints.ToInts(strings.Split(input, ","))

	var lastSpoken = map[int]int{}
	var lastSpokenBefore = map[int]int{}
	var spokenCount = map[int]int{}

	var lastNumber = 0
	for i, value := range values {
		spokenCount[value] = 1
		lastSpoken[value] = i + 1
		lastNumber = value
	}

	for i := len(values) + 1; i <= limit; i++ {
		if spokenCount[lastNumber] == 1 {
			lastNumber = 0
		} else {
			lastNumber = lastSpoken[lastNumber] - lastSpokenBefore[lastNumber]
		}
		lastSpokenBefore[lastNumber] = lastSpoken[lastNumber]
		lastSpoken[lastNumber] = i
		spokenCount[lastNumber]++
	}

	return lastNumber
}
