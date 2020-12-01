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

		if isIncreasing(a, b, c, d, e, f) && hasTwoAdjacent(a, b, c, d, e, f) {
			validCount += 1
		}
	}

	fmt.Println(validCount)
}

func isIncreasing(numbers ...int) bool {
	for i := 0; i < len(numbers) - 1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func hasTwoAdjacent(numbers ...int) bool {
	for i := 0; i < len(numbers) - 1; i++ {
		if numbers[i] == numbers[i+1] {
			return true
		}
	}
	return false
}
