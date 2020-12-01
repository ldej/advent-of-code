package tools

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadBytes(location string) []byte {
	bytes, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func ReadLines(location string) *bufio.Scanner {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func ReadRegex(location string, regex string) []map[string]string {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
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

func ReadInts(location string) []int {
	var ints []int
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func ReadStrings(location string) []string {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), "\n")
}