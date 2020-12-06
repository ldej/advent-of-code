package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	lines := tools.ReadStringDoubleNewlines("./2020/day06/input.txt")

	total := 0

	for _, line := range lines {
		people := strings.Split(line, "\n")
		total += CountUniqueAnswers(people)
	}

	return total
}

func CountUniqueAnswers(people []string) int {
	answers := map[rune]bool{}

	for _, person := range people {
		for _, answer := range person {
			answers[answer] = true
		}
	}

	return len(answers)
}
