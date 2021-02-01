package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("./day2/input.txt")
	lines := strings.Split(string(content), "\n")

	for _, line1 := range lines {
		for _, line2 := range lines {
			if result := differByOne(line1, line2); result != "" {
				fmt.Println(result)
				return
			}
		}
	}
}

func differByOne(left, right string) string {
	if left == "" || right == "" {
		return ""
	}
	differences := 0
	index := 0
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			differences += 1
			index = i
		}
		if differences > 1 {
			return ""
		}
	}
	if differences == 1 {
		res := ""
		for i := 0; i < len(left); i++ {
			if i != index {
				res = res + string(left[i])
			}
		}
		return res
	}
	return ""
}
