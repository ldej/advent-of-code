package tools

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/tools/myints"
)

func inputFilePath() string {
	for i := 0; i < 5; i++ {
		_, p, _, _ := runtime.Caller(i)
		if strings.Contains(p, "/tools/") && strings.Contains(p, "_test.go") {
			d, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			return d + "/testdata/input.txt"
		}
		if !strings.Contains(p, "/tools/") {
			return filepath.Dir(filepath.Dir(p)) + "/input.txt"
		}
	}
	log.Fatal("Can't find input.txt")
	return ""
}

func ReadBytes() []byte {
	bytes, err := ioutil.ReadFile(inputFilePath())
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func ReadRegex(regex string) []map[string]string {
	content, err := ioutil.ReadFile(inputFilePath())
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

func ReadInts() []int {
	var ints []int

	content, err := ioutil.ReadFile(inputFilePath())
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

func ReadIntSlice() []int {
	return ReadIntSlices()[0]
}

func ReadIntSlices() [][]int {
	var result [][]int

	content, err := ioutil.ReadFile(inputFilePath())
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

func ReadIntCsv() [][]int {
	var result [][]int

	content, err := ioutil.ReadFile(inputFilePath())
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

func ReadIntCsvOneLine() []int {
	return ReadIntCsv()[0]
}

func ReadIntDoubleNewlines() [][]int {
	var results [][]int

	content, err := ioutil.ReadFile(inputFilePath())
	if err != nil {
		log.Fatal(err)
	}
	blocks := strings.Split(string(content), "\n\n")
	for _, block := range blocks {
		if len(block) == 0 {
			continue
		}
		lines := strings.Split(block, "\n")
		var result []int
		for _, line := range lines {
			i, _ := strconv.Atoi(strings.TrimSpace(line))
			result = append(result, i)
		}
		results = append(results, result)
	}
	return results
}

func ReadString() string {
	return ReadStrings()[0]
}

func ReadStrings() []string {
	content, err := ioutil.ReadFile(inputFilePath())
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
}

func ReadStringSlices() [][]string {
	var result [][]string
	content, err := ioutil.ReadFile(inputFilePath())
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

func ReadStringsDoubleNewlines() []string {
	var result []string
	content, err := ioutil.ReadFile(inputFilePath())
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result = append(result, strings.Trim(line, "\n"))
	}
	return result
}

func ReadIntGrid() IntGrid {
	content, err := ioutil.ReadFile(inputFilePath())
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	grid := make(IntGrid, len(lines))
	for i, line := range lines {
		numbers := make([]int, len(line))
		for j, char := range line {
			numbers[j] = myints.ToInt(string(char))
		}
		grid[i] = numbers
	}
	return grid
}
