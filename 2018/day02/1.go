package main

import (
	"fmt"
		"github.com/ldej/advent-of-code-2018/common"
)

func main() {
	scanner := common.ReadLines("./day2/input.txt")

	two := 0
	three := 0
	for scanner.Scan() {
		line := scanner.Text()
		foundTwo, foundThree := containsTwoThree(countLetters(line))
		if foundTwo {
			two += 1
		}
		if foundThree {
			three += 1
		}
	}
	fmt.Println(two * three)
}

func countLetters(a string) map[string]int {
	counts := map[string]int{}
	for i := 0; i < len(a); i++ {
		letter := string(a[i])
		if _, found := counts[letter]; found {
			counts[letter] += 1
		} else {
			counts[letter] = 1
		}
	}
	return counts
}

func containsTwoThree(a map[string]int) (bool, bool) {
	hasTwo := false
	hasThree := false

	for _, count := range a {
		if count == 2 {
			hasTwo = true
		} else if count == 3 {
			hasThree = true
		}
	}
	return hasTwo, hasThree
}
