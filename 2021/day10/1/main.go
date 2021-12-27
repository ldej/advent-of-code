package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/stack"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadStrings("./2021/day10/input.txt")

	score := 0
	for _, line := range input {
		score += validate(line)
	}
	return score
}

var scoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var matchMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func validate(line string) int {
	s := stack.RuneStack{}
	for _, current := range line {
		_, found := matchMap[current]
		if found {
			s.Push(current)
		} else {
			a, b := s.Pop()
			if !b {
				return 0
			}
			if current != matchMap[a] {
				return scoreMap[current]
			}
		}
	}
	return 0
}
