package tools

import (
	"github.com/ldej/advent-of-code/tools/myints"
	"regexp"
	"strings"
)

const (
	ReHexColor = `^#[a-f0-9]{6}$`
	ReInt      = `-?\d+`
)

var (
	CompiledHexColorRegex = regexp.MustCompile(ReHexColor)
	CompiledIntRegex      = regexp.MustCompile(ReInt)
)

func IsHexColor(a string) bool {
	return CompiledHexColorRegex.MatchString(strings.ToLower(a))
}

func RegexNamedGroupsRepeat(a string, regex string) []map[string]string {
	compRegEx := regexp.MustCompile(regex)

	var results []map[string]string

	matches := compRegEx.FindAllStringSubmatch(a, -1)

	for _, match := range matches {
		resultsMap := make(map[string]string)
		for i, name := range compRegEx.SubexpNames() {
			if name == "" {
				continue
			}
			if i > 0 && i <= len(match) {
				resultsMap[name] = match[i]
			}
		}
		results = append(results, resultsMap)
	}
	return results
}

func RegexNamedGroups(a string, regex string) map[string]string {
	compRegEx := regexp.MustCompile(regex)

	match := compRegEx.FindStringSubmatch(a)

	resultsMap := make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if name == "" {
			continue
		}
		if i > 0 && i <= len(match) {
			resultsMap[name] = match[i]
		}
	}

	return resultsMap
}

// FindInt returns the first integer in the input
func FindInt(input string) int {
	return myints.ToInt(CompiledIntRegex.FindString(input))
}

func FindInts(input string) []int {
	matches := CompiledIntRegex.FindAllString(input, -1)
	var ints []int
	for _, i := range matches {
		ints = append(ints, myints.ToInt(i))
	}
	return ints
}
