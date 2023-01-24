package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	head := tools.Point{X: 0, Y: 0}
	tail := tools.Point{X: 0, Y: 0}

	visited := map[tools.Point]bool{
		tail: true,
	}

	input := tools.ReadStrings()
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction, steps := parts[0], myints.ToInt(parts[1])

		for i := 1; i <= steps; i++ {
			switch direction {
			case "L":
				head.X += 1
			case "R":
				head.X -= 1
			case "U":
				head.Y -= 1
			case "D":
				head.Y += 1
			}

			rightDiff := head.X - tail.X
			leftDiff := tail.X - head.X
			downDiff := head.Y - tail.Y
			upDiff := tail.Y - head.Y

			right := rightDiff == 2
			left := leftDiff == 2
			down := downDiff == 2
			up := upDiff == 2

			upRight := upDiff+rightDiff >= 3
			upLeft := upDiff+leftDiff >= 3
			downLeft := downDiff+leftDiff >= 3
			downRight := downDiff+rightDiff >= 3

			switch {
			case upLeft:
				tail.X -= 1
				tail.Y -= 1
			case upRight:
				tail.X += 1
				tail.Y -= 1
			case downLeft:
				tail.X -= 1
				tail.Y += 1
			case downRight:
				tail.X += 1
				tail.Y += 1
			case right:
				tail.X += 1
			case left:
				tail.X -= 1
			case down:
				tail.Y += 1
			case up:
				tail.Y -= 1
			}

			visited[tail] = true
		}
	}

	return len(visited)
}
