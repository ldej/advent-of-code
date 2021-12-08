package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	fishes := tools.ReadIntCsvOneLine("./2021/day06/input.txt")

	for i := 0; i < 80; i++ {
		var newFishes []int
		for _, fish := range fishes {
			if fish == 0 {
				newFishes = append(newFishes, 6)
				newFishes = append(newFishes, 8)
			} else {
				newFishes = append(newFishes, fish-1)
			}
		}
		fishes = newFishes
		fmt.Println(len(fishes))
	}

	return len(fishes)
}
