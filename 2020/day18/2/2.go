package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day18/example2.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day18/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStrings(input)

	var result int
	for _, line := range lines {
		v, _ := evaluate(strings.Replace(line, " ", "", -1), "")
		result += v
	}

	return result
}

func evaluate(input string, lastOp string) (int, string) {
	var result int
	var rest = input

	if input[0] == '(' {
		result, rest = evaluate(input[1:], "")
		rest = rest[1:]
	} else {
		result, rest = tools.ToInt(input[0:1]), input[1:]
	}

	for {
		if len(rest) == 0 || rest[0] == ')' {
			return result, rest
		}

		var op string
		op = rest[0:1]

		if lastOp == "+" && op == "*" {
			return result, rest
		}

		var right int
		right, rest = evaluate(rest[1:], op)

		if op == "+" {
			result += right
		} else if op == "*" {
			result *= right
		}
	}
}
