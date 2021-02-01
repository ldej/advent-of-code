package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

	result := run("7,13,x,x,59,x,31,19", 939)
	fmt.Println("Example:", result)

	result = run("29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,577,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,23,x,x,x,x,x,x,x,601,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37", 1000001)
	fmt.Println("Result:", result)
}

func run(input string, start int) int {
	inputs := strings.Split(input, ",")

	buses := []int{}
	for _, i := range inputs {
		a, err := strconv.Atoi(i)
		if err == nil {
			buses = append(buses, a)
		}
	}

	lowest := []int{}

	for _, bus := range buses {
		for i := 0; ; i++ {
			if bus*i > start {
				lowest = append(lowest, bus*i)
				break
			}
		}
	}
	fmt.Println(lowest)
	wait := myints.MinList(lowest) - start

	busID := buses[myints.IndexOf(lowest, myints.MinList(lowest))]
	return wait * busID
}
