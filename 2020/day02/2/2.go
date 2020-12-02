package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	result := tools.ReadRegex("./2020/day02/input.txt", `(?P<min>\d+)-(?P<max>\d+) (?P<letter>\w): (?P<password>.*)`)

	valid := 0
	for _, line := range result {
		min, _ := strconv.Atoi(line["min"])
		max, _ := strconv.Atoi(line["max"])
		letter := line["letter"]
		password := line["password"]

		a := tools.StringIndex(password, min-1)
		b := tools.StringIndex(password, max-1)

		if tools.StringCompareXOR(a, b, letter) {
			valid += 1
		}
	}
	return valid
}