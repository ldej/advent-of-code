package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code-2018/common"
)

type node struct {
	Header   [2]int
	Children []node
	Metadata []int
}

func main() {
	if calculate("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2") != 138 {
		fmt.Println("Error")
		return
	}
	results := common.ReadAllLines("./day8/input.txt", `(?P<data>.*)`)
	fmt.Println(calculate(results[0]["data"]))
}

func calculate(input string) int {
	values := strings.Split(input, " ")

	numbers := []int{}
	for _, value := range values {
		a, _ := strconv.Atoi(value)
		numbers = append(numbers, a)
	}

	_, n := getNode(numbers)

	return getMetadataTotal([]node{n})
}

func getMetadataTotal(nodes []node) int {
	total := 0
	for _, n := range nodes {
		for _, m := range n.Metadata {
			total += m
		}
		total += getMetadataTotal(n.Children)
	}
	return total
}

func getNode(numbers []int) ([]int, node) {
	numberOfChildren, numbers := numbers[0], numbers[1:]
	numberOfMetadataEntries, numbers := numbers[0], numbers[1:]

	var n node
	children := []node{}
	for i := 0; i < numberOfChildren; i++ {
		numbers, n = getNode(numbers)
		children = append(children, n)
	}

	var metadata []int
	metadataEntriesIndex := numberOfMetadataEntries
	numbers, metadata = numbers[metadataEntriesIndex:], numbers[:metadataEntriesIndex]

	return numbers, node{
		Header:   [2]int{numberOfChildren, numberOfMetadataEntries},
		Children: children,
		Metadata: metadata,
	}
}
