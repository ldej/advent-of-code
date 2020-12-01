package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code-2017/common"
)

func main() {
	scanner := common.ReadLines("./2/input.txt")

	total := 0
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		min := math.MaxInt16
		max := 0
		for _, v := range vals {
			if c, err := strconv.Atoi(v); err == nil {
				if c > max {
					max = c
				}
				if c < min {
					min = c
				}
			}
		}
		total += max - min
	}
	fmt.Println(total)
}
