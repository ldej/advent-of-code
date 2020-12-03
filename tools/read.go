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
		if len(line) == 0 {
			continue
		}
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
		if len(line) == 0 {
			continue
		}
		i, err := strconv.Atoi(line)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func ReadIntSlice(location string) []int {
	return ReadIntSlices(location)[0]
}

func ReadIntSlices(location string) [][]int {
	var result [][]int

	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		var ints []int
		for _, field := range fields {
			i, _ := strconv.Atoi(field)
			ints = append(ints, i)
		}
		result = append(result, ints)
	}
	return result
}

func ReadIntCsv(location string) [][]int {
	var result [][]int

	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// strings.Split instead of strings.Fields
		fields := strings.Split(line, ",")
		var ints []int
		for _, field := range fields {
			// Trim space
			i, _ := strconv.Atoi(strings.TrimSpace(field))
			ints = append(ints, i)
		}
		result = append(result, ints)
	}
	return result
}

func ReadIntCsvOneLine(location string) []int {
	return ReadIntCsv(location)[0]
}

func ReadString(location string) string {
	return ReadStrings(location)[0]
}

func ReadStrings(location string) []string {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
}

func ReadStringSlices(location string) [][]string {
	var result [][]string
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result = append(result, strings.Fields(line))
	}
	return result
}

func ReadRuneGrid(location string) RuneGrid {
	var result RuneGrid
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result = append(result, []rune(line))
	}
	return result
}
