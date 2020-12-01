package main

import (
	"fmt"
	"io/ioutil"
	"github.com/ldej/advent-of-code/2017/common"
)

func main() {
	bytes := common.ReadBytes("./1/input.txt")

	off := len(bytes) / 2
	total := 0
	for i, c := range bytes {
		if c == bytes[(i+off)%(len(bytes)-1)] {
			total += int(c) - 48 // or - '0'
		}
	}
	fmt.Println(total)
}
