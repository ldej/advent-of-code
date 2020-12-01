package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/2017/common"
)

var configurations [][]int

func main() {
	scanner := common.ReadLines("./6/input1.txt")

	var banks []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		for _, part := range parts {
			if number, err := strconv.Atoi(part); err == nil {
				banks = append(banks, number)
			}
		}
	}

	numBanks := len(banks)
	redistributions := 0

	for {
		startIndex := findLargestBank(banks)
		blocks := banks[startIndex]
		banks[startIndex] = 0
		for i := 1; i <= blocks; i++ {
			banks[(startIndex+i)%(numBanks)]++
		}
		redistributions++
		if isKnown(banks) {
			break
		}
	}
	fmt.Println(banks)
	fmt.Println(redistributions)
}

func findLargestBank(banks []int) int {
	largest := -1
	largestIndex := -1
	for index, bank := range banks {
		if bank > largest {
			largest = bank
			largestIndex = index
		}
	}
	return largestIndex
}

func isKnown(banks []int) bool {
	newBank := make([]int, len(banks))
	copy(newBank, banks)

	if len(configurations) == 0 {
		configurations = append(configurations, newBank)
		return false
	}
	for _, configuration := range configurations {
		if isEqual(newBank, configuration) {
			return true
		}
	}
	configurations = append(configurations, newBank)
	return false
}

func isEqual(bank1 []int, bank2 []int) bool {
	for i := range bank1 {
		if bank1[i] != bank2[i] {
			return false
		}
	}
	return true
}
