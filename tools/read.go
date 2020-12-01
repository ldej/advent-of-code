package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ReadBytes(location string) []byte {
	bytes, err := ioutil.ReadFile(location)
	if err != nil {
		fmt.Print(err)
	}
	return bytes
}

func ReadLines(location string) *bufio.Scanner {
	file, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func ReadRegex(location string, regex string) []map[string]string {
	content, _ := ioutil.ReadFile(location)
	lines := strings.Split(string(content), "\n")

	compRegEx := regexp.MustCompile(regex)

	var results []map[string]string

	for _, line := range lines {
		match := compRegEx.FindStringSubmatch(line)
		resultsMap := make(map[string]string)
		for i, name := range compRegEx.SubexpNames() {
			if i > 0 && i <= len(match) {
				resultsMap[name] = match[i]
			}
		}
		results = append(results, resultsMap)
	}
	return results
}
