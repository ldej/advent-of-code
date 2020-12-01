package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code/2019/common"
)

func main() {
	scanner := common.ReadLines("./day01/input.txt")

	fuel := 0
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		fuel += (a / 3) - 2
	}
	fmt.Println(fuel)
}
