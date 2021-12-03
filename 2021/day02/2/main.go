package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	strs := tools.ReadStrings("./2021/day02/input.txt")

	position := 0
	depth := 0
	aim := 0

	for _, str := range strs {
		parts := strings.Split(str, " ")
		command, x := parts[0], myints.ToInt(parts[1])

		switch command {
		case "forward":
			position += x
			depth += aim * x
		case "up":
			aim -= x
		case "down":
			aim += x
		}
	}
	return position * depth
}
