package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

type Rock struct {
	points []tools.Point
	maxX   int
	minX   int
}

func (r Rock) Add(a tools.Point) {
	for i, p := range r.points {
		p.Add(a)
		r.points[i] = p
	}
	r.minX += a.X
	r.maxX += a.X
}

func (r Rock) Is(p tools.Point) bool {
	for _, a := range r.points {
		if a.X == p.X && a.Y == p.Y {
			return true
		}
	}
	return false
}

type Chamber struct {
	resting map[tools.Point]bool
	left    int
	right   int
	height  int
}

func (c *Chamber) MoveHorizontal(r Rock, move tools.Point) {
	for _, p := range r.points {
		withinBoundaries := p.X+move.X > c.left && p.X+move.X < c.right &&
			!c.resting[tools.Point{X: p.X + move.X, Y: p.Y + move.Y}]
		if !withinBoundaries {
			return
		}

	}
	r.Add(move)
}

func (c *Chamber) MoveVertical(r Rock, move tools.Point) bool {
	willRest := false
	for _, p := range r.points {
		if c.resting[tools.Point{X: p.X + move.X, Y: p.Y + move.Y}] || p.Y+move.Y <= 0 {
			willRest = true
			break
		}
	}

	if willRest {
		for _, p := range r.points {
			c.resting[p] = true
			if p.Y > c.height {
				c.height = p.Y
			}
		}
		return true
	}
	r.Add(move)
	return false
}

func (c *Chamber) Print(r Rock) {
	for j := myints.Max(c.height, 4) + 6; j >= 0; j-- {
		for i := c.left; i <= c.right; i++ {
			if r.Is(tools.Point{X: i, Y: j}) {
				fmt.Printf("@")
			} else if i == c.left || i == c.right {
				if j == 0 {
					fmt.Printf("+")
				} else {
					fmt.Printf("|")
				}
			} else if j == 0 {
				fmt.Printf("-")
			} else if c.resting[tools.Point{X: i, Y: j}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func run(file ...string) int {
	jetPattern := tools.ReadString(file...)

	chamber := Chamber{
		resting: map[tools.Point]bool{},
		left:    0,
		right:   8,
	}

	rocks := []Rock{
		{
			points: []tools.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
			},
			minX: 0,
			maxX: 3,
		},
		{
			points: []tools.Point{
				{X: 1, Y: 2},
				{X: 0, Y: 1},
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 1, Y: 0},
			},
			minX: 1,
			maxX: 2,
		},
		{
			points: []tools.Point{
				{X: 2, Y: 2},
				{X: 2, Y: 1},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
			},
			minX: 0,
			maxX: 2,
		},
		{
			points: []tools.Point{
				{X: 0, Y: 3},
				{X: 0, Y: 2},
				{X: 0, Y: 1},
				{X: 0, Y: 0},
			},
			minX: 0,
			maxX: 0,
		},
		{
			points: []tools.Point{
				{X: 0, Y: 1},
				{X: 1, Y: 1},
				{X: 0, Y: 0},
				{X: 1, Y: 0},
			},
			minX: 0,
			maxX: 1,
		},
	}

	maxNumberOfRocks := 1000000000000
	jetPushes := 0

	history := map[[2]int][]int{}

	for rockNumber := 0; rockNumber < maxNumberOfRocks; rockNumber++ {

		state := [2]int{rockNumber % len(rocks), jetPushes % len(jetPattern)}

		if h, found := history[state]; found {
			lastRepeatingRockNumber, lastHeight := h[0], h[1]
			rocksToGo := maxNumberOfRocks - rockNumber
			cycleLength := rockNumber - lastRepeatingRockNumber

			// continue adding rocks until we can finish it off with full cycles
			if rocksToGo%cycleLength == 0 {
				heightInCycle := chamber.height - lastHeight
				cyclesToGo := rocksToGo / cycleLength
				return chamber.height + cyclesToGo*heightInCycle
			}
		}
		history[state] = []int{rockNumber, chamber.height}

		start := tools.Point{X: 3, Y: chamber.height + 4}
		r := Rock{}
		for _, p := range rocks[rockNumber%len(rocks)].points {
			r.points = append(r.points, p)
		}
		r.Add(start)

		resting := false
		for !resting {
			jetNumber := jetPushes % len(jetPattern)

			switch jetPattern[jetNumber] {
			case '>':
				chamber.MoveHorizontal(r, tools.Point{X: 1, Y: 0})
			case '<':
				chamber.MoveHorizontal(r, tools.Point{X: -1, Y: 0})
			}
			jetPushes++

			resting = chamber.MoveVertical(r, tools.Point{X: 0, Y: -1})
		}
	}

	return chamber.height
}
