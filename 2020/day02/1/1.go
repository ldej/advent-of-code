package main

import (
	"fmt"
	"strconv"
	"strings"

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

		count := strings.Count(password, letter)

		if count >= min && count <= max {
			valid += 1
		}
	}
	return valid
}