package tools

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/tools/myints"
)

func inputFilePath(file ...string) string {
	for i := 0; i < 5; i++ {
		_, p, _, _ := runtime.Caller(i)
		if !strings.Contains(p, "/tools/") {
			dir := filepath.Dir(filepath.Dir(p))
			if len(file) == 1 {
				return path.Join(dir, file[0])
			}
			return dir + "/input.txt"
		}
	}
	log.Fatal("Can't find input.txt")
	return ""
}

func ReadBytes(file ...string) []byte {
	bytes, err := os.ReadFile(inputFilePath(file...))
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func ReadRegex(regex string, file ...string) []map[string]string {
	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadInts(file ...string) []int {
	var ints []int

	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadIntSlice(file ...string) []int {
	return ReadIntSlices(file...)[0]
}

func ReadIntSlices(file ...string) [][]int {
	var result [][]int

	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadIntCsv(file ...string) [][]int {
	var result [][]int

	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadIntCsvOneLine(file ...string) []int {
	return ReadIntCsv(file...)[0]
}

func ReadIntDoubleNewlines(file ...string) [][]int {
	var results [][]int

	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadString(file ...string) string {
	return ReadStrings(file...)[0]
}

func ReadStrings(file ...string) []string {
	content, err := os.ReadFile(inputFilePath(file...))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
}

func ReadStringSlices(file ...string) [][]string {
	var result [][]string
	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadStringsDoubleNewlines(file ...string) []string {
	var result []string
	content, err := os.ReadFile(inputFilePath(file...))
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

func ReadIntGrid(file ...string) IntGrid {
	content, err := os.ReadFile(inputFilePath(file...))
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
