package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code-2017/common"
)

func main() {
	scanner := common.ReadLines("./2/input2.txt")

	total := 0
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		for _, _i := range vals {
			i, _ := strconv.Atoi(_i)
			for _, _j := range vals {
				j, _ := strconv.Atoi(_j)
				if i > j {
					res := i / j
					if res*j == i {
						total += res
					}
				}
			}
		}
	}
	fmt.Println(total)
}
