package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code/2019/common"
)

func main() {
	scanner := common.ReadLines("./day01/input.txt")

	fuel_total := 0
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())

		for {
			fuel := (a / 3) - 2

			if fuel < 0 {
				break
			}

			fuel_total += fuel
			a = fuel
		}
	}
	fmt.Println(fuel_total)
}
