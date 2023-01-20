package main

import (
	"fmt"
	"sort"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadIntDoubleNewlines()

	var elves []int
	for _, elf := range input {
		elves = append(elves, myints.Sum(elf))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return myints.Sum(elves[0:3])
}
