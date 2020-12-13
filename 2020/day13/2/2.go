package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("7,13,x,x,59,x,31,19")
	fmt.Println("Example:", result)

	result = run("17,x,13,19")
	fmt.Println("Example:", result)

	result = run("67,7,59,61")
	fmt.Println("Example:", result)

	result = run("29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,577,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,23,x,x,x,x,x,x,x,601,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37")
	fmt.Println("Result:", result)
}

func run(input string) int {
	inputs := strings.Split(input, ",")

	times := make([]int, 0)

	for _, in := range inputs {
		a, err := strconv.Atoi(in)
		if err == nil {
			times = append(times, a)
		} else {
			times = append(times, -1)
		}
	}

	earliest := times[0]
	departure := times[0]

	for i, t := range times {
		if t > 0 {
			for tools.GreatestCommonDivisor(earliest+i, t) != t {
				earliest += departure
			}
			departure = (departure * t) / tools.GreatestCommonDivisor(departure, t)
		}
	}

	return earliest
}
