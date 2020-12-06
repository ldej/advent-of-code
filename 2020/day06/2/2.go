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
		total += CountCommonAnswers(people)
	}

	return total
}

func CountCommonAnswers(people []string) int {
	person := people[0]

	count := 0
	for _, question := range person {
		if AllPeopleAnswered(people, question) {
			count += 1
		}
	}
	return count
}

func AllPeopleAnswered(people []string, question rune) bool {
	for _, person := range people {
		if !strings.Contains(person, string(question)) {
			return false
		}
	}
	return true
}
