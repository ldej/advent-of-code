package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/myring"
)

func main() {
	fmt.Println("Part 1")

	result := run(389125467, 100)
	fmt.Println("Example:", result)

	result = run(368195742, 100)
	fmt.Println("Result:", result)
}

func run(input int, steps int) int {
	var numbers = myints.ToDigits(input)
	var cupsCount = len(numbers)
	var cups = ring.New(cupsCount)

	// Use a map to easily find the position of a number in the ring
	var cupsMap = make(map[int]*ring.Ring)

	for _, cup := range numbers {
		cups.Value = cup
		cupsMap[cup] = cups
		cups = cups.Next()
	}

	for i := 0; i < steps; i++ {
		fmt.Printf("-- move %d --\n", i)
		fmt.Printf("cups: %s\n", myring.String(cups))
		nextThree := cups.Unlink(3)
		fmt.Printf("pick up: %s\n", myring.String(nextThree))
		destination := cups.Value.(int) - 1

		if destination == 0 {
			destination = cupsCount
		}
		for myring.Contains(nextThree, destination) {
			destination = destination - 1
			if destination == 0 {
				destination = cupsCount
			}
		}
		fmt.Printf("destination: %d\n\n", destination)
		cupsMap[destination].Link(nextThree)
		cups = cups.Next()
	}

	cups = cupsMap[1].Next()

	var result strings.Builder
	for i := 0; i < cupsCount-1; i++ {
		result.WriteString(strconv.Itoa(cups.Value.(int)))
		cups = cups.Next()
	}
	return myints.ToInt(result.String())
}
