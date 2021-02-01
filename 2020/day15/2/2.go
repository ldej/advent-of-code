package main

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myints"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	fmt.Println("Part 2")

	result := run("0,3,6", 30000000)
	fmt.Println("Example:", result)

	result = run("6,19,0,5,7,13,1", 30000000)
	fmt.Println("Result:", result)
}

func run(input string, limit int) int {
	values := myints.ToInts(strings.Split(input, ","))

	var lastSpoken = map[int]int{}
	var lastSpokenBefore = map[int]int{}
	var spokenCount = map[int]int{}
	var lastNumber int

	bar := pb.StartNew(limit)
	bar.SetRefreshRate(time.Second)

	for i, value := range values {
		bar.Increment()

		spokenCount[value] = 1
		lastSpoken[value] = i + 1
		lastNumber = value
	}

	for i := len(values) + 1; i <= limit; i++ {
		bar.Increment()

		if spokenCount[lastNumber] == 1 {
			lastNumber = 0
		} else {
			lastNumber = lastSpoken[lastNumber] - lastSpokenBefore[lastNumber]
		}
		lastSpokenBefore[lastNumber] = lastSpoken[lastNumber]
		lastSpoken[lastNumber] = i
		spokenCount[lastNumber]++
	}
	bar.Finish()

	return lastNumber
}
