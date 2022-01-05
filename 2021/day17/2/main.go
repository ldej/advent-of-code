package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadRegex(`.*x=(?P<x1>-?\d+)\.\.(?P<x2>-?\d+), y=(?P<y1>-?\d+)\.\.(?P<y2>-?\d+)`)[0]
	targetMin := tools.Point{X: myints.ToInt(input["x1"]), Y: myints.ToInt(input["y2"])}
	targetMax := tools.Point{X: myints.ToInt(input["x2"]), Y: myints.ToInt(input["y1"])}

	var position tools.Point
	var velocity tools.Point

	var count int
	for i := 1; i <= targetMax.X; i++ {
		for j := -300; j < 300; j++ {
			velocity.X = i
			velocity.Y = j

			y := HighestY(position, velocity, targetMin, targetMax)
			if y >= 0 {
				count++
			}
		}
	}

	return count
}

func HighestY(position tools.Point, velocity tools.Point, targetMin tools.Point, targetMax tools.Point) int {
	highestY := 0
	hit := isWithinTarget(position, targetMin, targetMax)
	for hit < 1 {
		position.Add(velocity)
		if position.Y > highestY {
			highestY = position.Y
		}

		if velocity.X > 0 {
			velocity.X -= 1
		} else if velocity.X < 0 {
			velocity.X += 1
		}
		velocity.Y -= 1

		if position.Y < targetMax.Y {
			return -1
		}

		hit = isWithinTarget(position, targetMin, targetMax)
		if hit == 0 {
			return highestY
		}
	}
	return -1
}

func isWithinTarget(position tools.Point, targetMin tools.Point, targetMax tools.Point) int {
	if position.X >= targetMin.X && position.X <= targetMax.X && position.Y <= targetMin.Y && position.Y >= targetMax.Y {
		return 0
	}
	if position.X < targetMin.X || position.Y > targetMin.Y {
		return -1
	}
	if position.X > targetMax.X || position.Y < targetMax.Y {
		return 1
	}
	return -1
}
