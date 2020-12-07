package tools

import (
	"regexp"
	"strings"
)

const (
	ReHexColor = `^#[a-f0-9]{6}$`
)

var (
	CompiledHexColorRegex = regexp.MustCompile(ReHexColor)
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
