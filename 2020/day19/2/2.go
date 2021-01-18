package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day19/example2.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day19/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStringsDoubleNewlines(input)
	ruleLines := strings.Split(lines[0], "\n")
	messages := strings.Split(lines[1], "\n")

	rulesMap := map[int]string{}
	for _, ruleLine := range ruleLines {
		parts := strings.Split(ruleLine, ": ")
		rulesMap[tools.ToInt(parts[0])] = parts[1]
	}

	rulesMap[8] = "42 | 42 8"
	rulesMap[11] = "42 31 | 42 11 31"

	rule42 := createRegex(rulesMap[42], rulesMap, make(map[string]string, 0))
	rule31 := createRegex(rulesMap[31], rulesMap, make(map[string]string, 0))

	combined := fmt.Sprintf("^(?P<42>(%s)+)(?P<31>(%s)+)$", rule42, rule31)

	var correct int
	for _, message := range messages {
		result := tools.RegexNamedGroups(message, combined)

		if len(result) > 0 {
			matches42 := regexp.MustCompile(rule42).FindAllStringIndex(result["42"], -1)
			matches31 := regexp.MustCompile(rule31).FindAllStringIndex(result["31"], -1)

			if len(matches42) > len(matches31) {
				correct += 1
			}
		}
	}
	return correct
}

func createRegex(value string, rules map[int]string, cache map[string]string) string {
	value = strings.Trim(value, "\"")

	if val, found := cache[value]; found {
		return val
	}

	if value == "a" || value == "b" {
		cache[value] = value
		return value
	} else if strings.Contains(value, " | ") {
		options := strings.Split(value, " | ")
		var regexes []string
		for _, option := range options {
			regexes = append(regexes, createRegex(option, rules, cache))
		}
		result := fmt.Sprintf("(%s)", strings.Join(regexes, "|"))
		cache[value] = result
		return result
	} else {
		numbers := strings.Split(value, " ")
		var results []string
		for _, number := range numbers {
			results = append(results, createRegex(rules[tools.ToInt(number)], rules, cache))
		}
		result := strings.Join(results, "")
		cache[value] = result
		return result
	}
}
