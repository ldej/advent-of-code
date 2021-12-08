package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadStrings("./2021/day08/input.txt")

	counter := map[int]int{}
	for _, line := range input {
		parts := strings.Split(line, "|")
		_, output := parts[0], parts[1] // signal
		for _, part := range strings.Fields(output) {
			counter[len(part)]++
		}
	}
	return counter[2] + counter[4] + counter[3] + counter[7]
}
