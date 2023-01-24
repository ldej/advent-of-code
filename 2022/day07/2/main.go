package main

import (
	"fmt"
	"math"
	"path"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadStrings()

	files := make(map[string]int)

	var currentDir string
	for _, line := range input {
		if strings.HasPrefix(line, "$") {
			line = strings.TrimPrefix(line, "$ ")
			if strings.HasPrefix(line, "ls") {
				// ignore
			} else if strings.HasPrefix(line, "cd") {
				currentDir = path.Join(currentDir, strings.TrimPrefix(line, "cd "))
			}
		} else if strings.HasPrefix(line, "dir") {
			// ignore
		} else {
			parts := strings.Split(line, " ")
			files[path.Join(currentDir, parts[1])] = tools.FindInt(parts[0])
		}
	}

	dirs := make(map[string]int)

	for p, size := range files {
		for {
			p = path.Dir(p)
			dirs[p] += size
			if p == "/" {
				break
			}
		}
	}

	diskSize := 70000000
	usedSpace := dirs["/"]
	availableSpace := diskSize - usedSpace
	requiredForUpdate := 30000000
	sizeToDelete := requiredForUpdate - availableSpace

	smallest := math.MaxInt

	for _, size := range dirs {
		if size >= sizeToDelete && size < smallest {
			smallest = size
		}
	}
	return smallest
}
