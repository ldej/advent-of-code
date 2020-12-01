package main

import (
	"github.com/ldej/advent-of-code/2018/common"
	"fmt"
	"strings"
)

func main() {
	results := common.ReadAllLines("./day5/input.txt", `(?P<input>.*)`)

	linestring := results[0]["input"]

	for i := 'a'; i <= 'z'; i++ {
		fmt.Println(string(byte(i)), calculateLength(linestring, i))
	}
}

func calculateLength(linestring string, skip int32) int {
	linestring = strings.Replace(linestring, string(skip), "", -1)
	linestring = strings.Replace(linestring, string(skip-32), "", -1)
	line := []byte(linestring)

	somethingChanged := true
	for somethingChanged {
		somethingChanged = false

		newline := line[:0]
		for idx := 0; idx < len(line); idx += 1 {

			// Don't forget to copy the last letter
			if idx == len(line) - 1 {
				newline = append(newline, line[idx])
				continue
			}

			// Find difference between letters
			diff := uint8(0)
			if line[idx] > line[idx+1] {
				diff = line[idx] - line[idx+1]
			} else {
				diff = line[idx+1] - line[idx]
			}

			// If it's the same letters
			if diff == 32 {
				// skip one
				idx += 1
				somethingChanged = true
			} else {
				newline = append(newline, line[idx])
			}
		}
		line = newline
	}
	return len(line)
}
