package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run(2020)
	fmt.Println(result)
}

func run(expected int) int {
	ints := tools.ReadInts("./2020/day01/input.txt")

	for ia, a := range ints {
		for ib, b := range ints {
			if ib == ia {
				continue
			}
			if a + b == expected {
				fmt.Printf("%d + %d = %d\n", a, b, expected)
				fmt.Printf("%d * %d= %d\n", a, b, a*b)
				return a * b
			}
		}
	}
	return -1
}
