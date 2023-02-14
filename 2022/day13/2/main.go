package main

import (
	"fmt"
	"sort"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadStrings(file...)

	divider1, _ := parse("[[2]]")
	divider2, _ := parse("[[6]]")

	lines := NodesList{divider1, divider2}

	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		parsed, _ := parse(line)
		lines = append(lines, parsed)
	}
	sort.Sort(lines)

	result := 1
	for i, n := range lines {
		if n.Print() == divider1.Print() || n.Print() == divider2.Print() {
			result *= i + 1
		}
	}
	return result
}

type NodesList []Nodes

func (n NodesList) Len() int { return len(n) }
func (n NodesList) Less(i, j int) bool {
	b, _ := inRightOrder(n[i], n[j])
	return b
}
func (n NodesList) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

type Node struct {
	value int
	nodes *Nodes
}

type Nodes []Node

func (n *Node) Print() string {
	if n.nodes != nil {
		return fmt.Sprintf("[%s]", n.nodes.Print())
	} else {
		return fmt.Sprintf("%d", n.value)
	}
}

func (n *Nodes) Print() string {
	var a string
	for i, node := range *n {
		a = a + node.Print()
		if i < len(*n)-1 {
			a = a + ","
		}
	}
	return a
}

func parse(line string) (Nodes, string) {
	var nodes Nodes
	var number string

	for len(line) > 0 {
		var c byte
		c, line = line[0], line[1:]
		switch c {
		case '[':
			var parsed Nodes
			parsed, line = parse(line)
			nodes = append(nodes, Node{nodes: &parsed, value: -1})
		case ',', ']':
			if len(number) > 0 {
				nodes = append(nodes, Node{value: myints.ToInt(number)})
				number = ""
			}
		default:
			// must be a number
			number = number + string(c)
		}
		if c == ']' {
			break
		}
	}
	return nodes, line
}

func inRightOrder(left Nodes, right Nodes) (correct bool, found bool) {
	for i := 0; i < len(left) && i < len(right); i++ {
		// if at least one of them is a slice, do a slice comparison
		if left[i].nodes != nil || right[i].nodes != nil {
			leftNodes := left[i].nodes
			rightNodes := right[i].nodes
			if left[i].nodes == nil {
				leftNodes = &Nodes{Node{value: left[i].value}}
			}
			if right[i].nodes == nil {
				rightNodes = &Nodes{Node{value: right[i].value}}
			}
			answer, found := inRightOrder(*leftNodes, *rightNodes)
			if found {
				return answer, found
			}
		} else if left[i].value < right[i].value {
			// if they are both numbers, do a number comparison
			return true, true
		} else if left[i].value > right[i].value {
			return false, true
		}
	}

	// No comparisons were decisive, check the lengths
	if len(right) > len(left) {
		return true, true
	} else if len(right) < len(left) {
		return false, true
	}

	return false, false
}
