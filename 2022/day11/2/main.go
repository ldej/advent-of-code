package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadStringsDoubleNewlines(file...)

	type monkey struct {
		inspected int
		items     []int
		operation func(a int) int
		test      int
		true      int
		false     int
	}
	var monkeys []monkey

	for _, block := range input {
		lines := strings.Split(block, "\n")
		monkeys = append(monkeys, monkey{
			items: tools.FindInts(lines[1]),
			operation: func(value int, multiply bool) func(a int) int {
				if value == -1 {
					return func(a int) int {
						if multiply {
							return a * a
						} else {
							return a + a
						}
					}
				}
				return func(a int) int {
					if multiply {
						return a * value
					} else {
						return a + value
					}
				}
			}(
				tools.FindIntOr(lines[2], -1),
				strings.Contains(lines[2], "*"),
			),
			test:  tools.FindInt(lines[3]),
			true:  tools.FindInt(lines[4]),
			false: tools.FindInt(lines[5]),
		})
	}

	var divisors []int
	for _, m := range monkeys {
		divisors = append(divisors, m.test)
	}
	divisor := myints.Product(divisors)

	rounds := 10000
	for i := 0; i < rounds; i++ {
		for j := range monkeys {
			itemsToInspect := monkeys[j].items
			monkeys[j].items = []int{}
			for _, item := range itemsToInspect {
				worryLevel := monkeys[j].operation(item) % divisor
				if worryLevel%monkeys[j].test == 0 {
					monkeys[monkeys[j].true].items = append(monkeys[monkeys[j].true].items, worryLevel)
				} else {
					monkeys[monkeys[j].false].items = append(monkeys[monkeys[j].false].items, worryLevel)
				}
				monkeys[j].inspected++
			}
		}
	}

	var inspected []int
	for _, m := range monkeys {
		inspected = append(inspected, m.inspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return inspected[0] * inspected[1]
}
