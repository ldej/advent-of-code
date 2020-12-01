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
			for ic, c := range ints {
				if ia == ic || ib == ic {
					continue
				}
				if a + b + c == expected {
					fmt.Printf("%d + %d + %d = %d\n", a, b, c, expected)
					fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
					return a * b * c
				}
			}
		}
	}
	return -1
}
