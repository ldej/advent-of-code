package main

import (
	"fmt"
	"strconv"
)

func main() {

	start := 168630
	end := 718098

	validCount := 0

	for i := start; i <= end; i++ {
		numberString := strconv.Itoa(i)
		var a, b, c, d, e, f int
		_, _ = fmt.Sscanf(numberString, "%1d%1d%1d%1d%1d%1d", &a, &b, &c, &d, &e, &f)

		if isIncreasing(a, b, c, d, e, f) && hasExactlyTwoAdjacent(a, b, c, d, e, f) {
			validCount += 1
		}
	}

	fmt.Println(validCount)
}

func isIncreasing(numbers ...int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func hasExactlyTwoAdjacent(numbers ...int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		previous := -1
		if i > 0 {
			previous = numbers[i-1]
		}
		next := -1
		if i+2 < len(numbers) {
			next = numbers[i+2]
		}
		if previous != numbers[i] && numbers[i] == numbers[i+1] && numbers[i] != next {
			return true
		}
	}
	return false
}
