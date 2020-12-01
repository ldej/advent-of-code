package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code-2017/common"
)

func main() {
	scanner := common.ReadLines("./4/input1.txt")

	valid := 0
	for scanner.Scan() {
		if isValidPassphrase(scanner.Text()) {
			valid++
		}
	}
	fmt.Println(valid)
}

func isValidPassphrase(passphrase string) bool {
	words := strings.Fields(passphrase)
	found := []string{}

	for _, word := range words {
		for _, f := range found {
			if word == f {
				return false
			}
		}
		found = append(found, word)
	}
	return true
}
