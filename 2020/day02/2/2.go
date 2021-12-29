package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	result := tools.ReadRegex(`(?P<min>\d+)-(?P<max>\d+) (?P<letter>\w): (?P<password>.*)`)

	valid := 0
	for _, line := range result {
		min := myints.ToInt(line["min"])
		max := myints.ToInt(line["max"])
		letter := line["letter"]
		password := line["password"]

		a := mystrings.Index(password, min-1)
		b := mystrings.Index(password, max-1)

		if mystrings.CompareXOR(a, b, letter) {
			valid += 1
		}
	}
	return valid
}
