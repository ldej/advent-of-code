package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadIntCsvOneLine()

	min, max := myints.MinAndMax(input)

	minFuel := -1

	for i := min; i <= max; i++ {
		fuel := 0
		for _, x := range input {
			fuel += myints.Abs(x - i)
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}
