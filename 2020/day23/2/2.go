package main

import (
	"container/ring"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/myring"
	"time"
)

func main() {
	fmt.Println("Part 2")

	result := run(389125467, 10000000)
	fmt.Println("Example:", result)

	result = run(368195742, 10000000)
	fmt.Println("Result:", result)
}

func run(input int, steps int) int {
	var cupsCount = 1000000
	var numbers = myints.ToDigits(input)
	var cups = ring.New(cupsCount)

	// Use a map to easily find the position of a number in the ring
	var cupsMap = make(map[int]*ring.Ring)

	for _, cup := range numbers {
		cups.Value = cup
		cupsMap[cup] = cups
		cups = cups.Next()
	}
	for i := 10; i <= cupsCount; i++ {
		cups.Value = i
		cupsMap[i] = cups
		cups = cups.Next()
	}

	bar := pb.StartNew(steps)
	bar.SetRefreshRate(time.Second)

	for i := 0; i < steps; i++ {
		bar.Increment()

		nextThree := cups.Unlink(3)
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
		cupsMap[destination].Link(nextThree)
		cups = cups.Next()
	}
	bar.Finish()

	cup1 := cupsMap[1]
	return cup1.Move(1).Value.(int) * cup1.Move(2).Value.(int)
}
