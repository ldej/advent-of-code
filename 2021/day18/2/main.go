package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/mystrings"
	"github.com/ldej/advent-of-code/tools/stack"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadStrings()

	var snailSums [][]string
	for _, line := range input {
		snailSums = append(snailSums, strings.Split(line, ""))
	}

	max := 0
	for i := 0; i < len(snailSums); i++ {
		for j := 0; j < len(snailSums); j++ {
			mag := magnitude(reduce(addSnailSums(snailSums[i], snailSums[j])))
			if mag > max {
				max = mag
			}
		}
	}

	return max
}

func addSnailSums(one []string, two []string) []string {
	return mystrings.AppendAll([]string{"["}, one, []string{","}, two, []string{"]"})
}

func reduce(snailSum []string) []string {
	change := true
	for change {
		snailSum, change = explode(snailSum)
		if change {
			continue
		}
		snailSum, change = split(snailSum)
	}
	return snailSum
}

func explode(snailSum []string) ([]string, bool) {
	s := stack.StringStack{}
	depth := 0
	for i, part := range snailSum {
		switch part {
		case "[":
			s.Push(part)
			depth += 1
		case "]":
			s.Pop()
			depth -= 1
		case ",":
		default:
			value := myints.ToInt(part)
			top, _ := s.Peek()
			topValue := myints.ToIntOr(top, -1)
			if depth > 4 && topValue >= 0 {
				// explode to the left and right
				snailSum = add(snailSum, topValue, i-2, true)
				snailSum = add(snailSum, value, i, false)
				// replace the pair with a 0
				return mystrings.AppendAll(snailSum[:i-3], []string{"0"}, snailSum[i+2:]), true
			} else {
				s.Push(part)
			}
		}
	}
	return snailSum, false
}

func add(snailSum []string, toAdd int, fromIndex int, left bool) []string {
	if left {
		for i := fromIndex - 1; i >= 0; i-- {
			value := myints.ToIntOr(snailSum[i], -1)
			if value >= 0 {
				snailSum[i] = strconv.Itoa(value + toAdd)
				return snailSum
			}
		}
	} else {
		for i := fromIndex + 1; i < len(snailSum); i++ {
			value := myints.ToIntOr(snailSum[i], -1)
			if value >= 0 {
				snailSum[i] = strconv.Itoa(value + toAdd)
				return snailSum
			}
		}
	}
	return snailSum
}

func split(snailSum []string) ([]string, bool) {
	for i, part := range snailSum {
		if value := myints.ToIntOr(part, -1); value >= 10 {
			left := value / 2
			right := value - left
			toInsert := []string{"[", strconv.Itoa(left), ",", strconv.Itoa(right), "]"}
			return mystrings.AppendAll(snailSum[:i], toInsert, snailSum[i+1:]), true
		}
	}
	return snailSum, false
}

func magnitude(snailSum []string) int {
	for len(snailSum) > 1 {
		snailSum = calculateOneMagnitude(snailSum)
	}
	return myints.ToInt(snailSum[0])
}

func calculateOneMagnitude(snailSum []string) []string {
	s := stack.StringStack{}
	for i, part := range snailSum {
		switch part {
		case "[":
			s.Push(part)
		case "]":
			s.Pop()
		case ",":
		default:
			value := myints.ToInt(part)
			top, _ := s.Peek()
			topValue := myints.ToIntOr(top, -1)
			if topValue >= 0 {
				total := 3*topValue + 2*value
				return mystrings.AppendAll(snailSum[:i-3], []string{strconv.Itoa(total)}, snailSum[i+2:])
			} else {
				s.Push(part)
			}
		}
	}
	return snailSum
}
