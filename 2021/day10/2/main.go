package main

import (
	"fmt"
	"sort"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/stack"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadStrings("./2021/day10/input.txt")

	var scores []int
	for _, line := range input {
		score, completed := complete(line)
		if completed {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

var scoreMap = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var matchMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func complete(line string) (int, bool) {
	s := stack.RuneStack{}
	for _, current := range line {
		_, found := matchMap[current]
		if found {
			s.Push(current)
		} else {
			toMatch, isEmpty := s.Pop()
			if !isEmpty {
				continue
			}
			if current != matchMap[toMatch] {
				return -1, false
			}
		}
	}

	score := 0
	for !s.IsEmpty() {
		toMatch, isEmpty := s.Pop()
		if !isEmpty {
			continue
		}
		score = (score * 5) + scoreMap[matchMap[toMatch]]
	}
	return score, true
}
