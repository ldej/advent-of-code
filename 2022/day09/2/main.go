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
	knots := make([]tools.Point, 10)

	visited := map[tools.Point]bool{
		knots[9]: true,
	}

	input := tools.ReadStrings()
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction, steps := parts[0], myints.ToInt(parts[1])

		for i := 1; i <= steps; i++ {
			switch direction {
			case "L":
				knots[0].X += 1
			case "R":
				knots[0].X -= 1
			case "U":
				knots[0].Y -= 1
			case "D":
				knots[0].Y += 1
			}

			for j := 0; j < len(knots)-1; j++ {
				rightDiff := knots[j].X - knots[j+1].X
				leftDiff := knots[j+1].X - knots[j].X
				downDiff := knots[j].Y - knots[j+1].Y
				upDiff := knots[j+1].Y - knots[j].Y

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
					knots[j+1].X -= 1
					knots[j+1].Y -= 1
				case upRight:
					knots[j+1].X += 1
					knots[j+1].Y -= 1
				case downLeft:
					knots[j+1].X -= 1
					knots[j+1].Y += 1
				case downRight:
					knots[j+1].X += 1
					knots[j+1].Y += 1
				case right:
					knots[j+1].X += 1
				case left:
					knots[j+1].X -= 1
				case down:
					knots[j+1].Y += 1
				case up:
					knots[j+1].Y -= 1
				}
			}
			visited[knots[9]] = true
		}
	}

	return len(visited)
}
