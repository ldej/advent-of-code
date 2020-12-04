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
	lines := tools.ReadStringDoubleNewlines("./2020/day04/input.txt")

	valid := 0

	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, line := range lines {
		line = strings.Replace(line, "\n", " ", -1)

		if !ContainsKeys(line, requiredKeys) {
			continue
		}

		if !KeysAreCorrect(line) {
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

func KeysAreCorrect(line string) bool {
	keyValues := strings.Split(line, " ")

	for _, keyValue := range keyValues {
		split := strings.Split(keyValue, ":")

		if !KeyIsCorrect(split[0], split[1]) {
			return false
		}
	}
	return true
}

func KeyIsCorrect(key string, value string) bool {
	switch key {
	case "byr":
		return tools.InRange(tools.ToInt(value), 1920, 2002)
	case "iyr":
		return tools.InRange(tools.ToInt(value), 2010, 2020)
	case "eyr":
		return tools.InRange(tools.ToInt(value), 2020, 2030)
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			value = strings.TrimSuffix(value, "cm")
			return tools.InRange(tools.ToInt(value), 150, 193)
		} else if strings.HasSuffix(value, "in") {
			value = strings.TrimSuffix(value, "in")
			return tools.InRange(tools.ToInt(value), 59, 76)
		}
		return false
	case "hcl":
		return tools.IsHexColor(value)
	case "ecl":
		eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		return tools.StringSliceContains(eyeColors, value)
	case "pid":
		if len(value) != 9 {
			return false
		}
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return true
	case "cid":
		return true
	default:
		// unknown key
		return false
	}
}
