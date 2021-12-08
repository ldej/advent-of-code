package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int64 {
	fishes := tools.ReadIntCsvOneLine("./2021/day06/input.txt")

	fishCounter := [9]int64{}
	for _, fish := range fishes {
		fishCounter[fish]++
	}

	for i := 0; i < 256; i++ {
		newFishCounter := [9]int64{}
		for days, count := range fishCounter {
			if days == 0 {
				newFishCounter[6] += count
				newFishCounter[8] += count
			} else {
				newFishCounter[days-1] += count
			}
		}
		fishCounter = newFishCounter

		sum := int64(0)
		for _, count := range fishCounter {
			sum += count
		}
	}

	sum := int64(0)
	for _, count := range fishCounter {
		sum += count
	}

	return sum
}
