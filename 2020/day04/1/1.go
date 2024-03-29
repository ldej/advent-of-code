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
	lines := tools.ReadStringsDoubleNewlines()

	valid := 0

	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, line := range lines {
		passport := strings.Replace(line, "\n", " ", -1)

		if !ContainsKeys(passport, requiredKeys) {
			continue
		}
		valid += 1
	}
	return valid
}

func ContainsKeys(line string, keys []string) bool {
	for _, key := range keys {
		if !strings.Contains(line, key+":") {
			return false
		}
	}
	return true
}
