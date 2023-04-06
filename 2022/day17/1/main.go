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

type Rock []tools.Point

func (r Rock) Add(a tools.Point) {
	for i, p := range r {
		p.Add(a)
		r[i] = p
	}
}

func (r Rock) MaxX() int {
	maxX := 0
	for i := range r {
		if r[i].X > maxX {
			maxX = r[i].X
		}
	}
	return maxX
}

func (r Rock) MinX() int {
	minX := 7
	for i := range r {
		if r[i].X < minX {
			minX = r[i].X
		}
	}
	return minX
}

func (r Rock) Is(p tools.Point) bool {
	for _, a := range r {
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
	withinBoundaries := true
	for _, p := range r {
		withinBoundaries = withinBoundaries &&
			p.X+move.X > c.left && p.X+move.X < c.right &&
			!c.resting[tools.Point{X: p.X + move.X, Y: p.Y + move.Y}]

	}
	if withinBoundaries {
		r.Add(move)
	}
}

func (c *Chamber) MoveVertical(r Rock, move tools.Point) bool {
	willRest := false
	for _, p := range r {
		if c.resting[tools.Point{X: p.X + move.X, Y: p.Y + move.Y}] || p.Y+move.Y <= 0 {
			willRest = true
			break
		}
	}

	if willRest {
		for _, p := range r {
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
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
		},
		{
			{X: 1, Y: 2},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 1, Y: 0},
		},
		{
			{X: 2, Y: 2},
			{X: 2, Y: 1},
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
		},
		{
			{X: 0, Y: 3},
			{X: 0, Y: 2},
			{X: 0, Y: 1},
			{X: 0, Y: 0},
		},
		{
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 0, Y: 0},
			{X: 1, Y: 0},
		},
	}

	maxNumberOfRocks := 2022
	jetPushes := 0

	for i := 0; i < maxNumberOfRocks; i++ {
		start := tools.Point{X: 3, Y: chamber.height + 4}
		r := Rock{}
		for _, p := range rocks[i%5] {
			r = append(r, p)
		}
		r.Add(start)

		resting := false
		for !resting {
			// chamber.Print(r)
			switch jetPattern[jetPushes%len(jetPattern)] {
			case '>':
				chamber.MoveHorizontal(r, tools.Point{X: 1, Y: 0})
			case '<':
				chamber.MoveHorizontal(r, tools.Point{X: -1, Y: 0})
			}
			jetPushes++

			// chamber.Print(r)
			resting = chamber.MoveVertical(r, tools.Point{X: 0, Y: -1})
		}
	}

	return chamber.height
}
