package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day18/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day18/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStrings()

	var result int
	for _, line := range lines {
		v, _ := calculate(strings.Replace(line, " ", "", -1))
		result += v
	}

	return result
}

func calculate(a string) (int, string) {
	var result int
	var rest = a
	if a[0] == '(' {
		result, rest = calculate(a[1:])
	} else {
		result, rest = myints.ToInt(a[0:1]), a[1:]
	}

	for {
		if len(rest) == 0 {
			return result, ""
		} else if rest[0] == ')' {
			return result, rest[1:]
		}
		var op string
		op, rest = rest[0:1], rest[1:]

		var next string
		next, rest = rest[0:1], rest[1:]

		var right int
		if next == "(" {
			right, rest = calculate(rest)
		} else {
			right = myints.ToInt(next)
		}

		if op == "+" {
			result += right
		} else if op == "*" {
			result *= right
		}
	}
}
