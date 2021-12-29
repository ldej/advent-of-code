package main

import (
	"fmt"
	"image"
	"sort"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	field := tools.ReadIntGrid()

	lowestPoints := getLowestPoints(field)

	var basinSizes []int
	for _, point := range lowestPoints {
		_, _, size := getBasinSize(field, []image.Point{point}, map[image.Point]bool{}, 0)
		basinSizes = append(basinSizes, size)
	}
	sort.Ints(basinSizes)

	result := 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		result *= size
	}
	return result
}

func getLowestPoints(field [][]int) []image.Point {
	var lowestPoints []image.Point
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			lowest := true
			v := field[i][j]

			if i-1 >= 0 {
				lowest = lowest && v < field[i-1][j]
			}
			if i+1 < len(field) {
				lowest = lowest && v < field[i+1][j]
			}
			if j-1 >= 0 {
				lowest = lowest && v < field[i][j-1]
			}
			if j+1 < len(field[i]) {
				lowest = lowest && v < field[i][j+1]
			}

			if lowest {
				lowestPoints = append(lowestPoints, image.Point{X: i, Y: j})
			}
		}
	}
	return lowestPoints
}

func getBasinSize(field [][]int, pointsToConsider []image.Point, considered map[image.Point]bool, size int) ([][]int, []image.Point, int) {
	if len(pointsToConsider) == 0 {
		return field, pointsToConsider, size
	}

	var point image.Point
	point, pointsToConsider = pointsToConsider[0], pointsToConsider[1:]
	v := field[point.X][point.Y]

	p := image.Point{X: point.X - 1, Y: point.Y}
	if point.X-1 >= 0 && field[point.X-1][point.Y] > v && field[point.X-1][point.Y] < 9 && !considered[p] {
		pointsToConsider = append(pointsToConsider, p)
		considered[p] = true
	}
	p = image.Point{X: point.X + 1, Y: point.Y}
	if point.X+1 < len(field) && field[point.X+1][point.Y] > v && field[point.X+1][point.Y] < 9 && !considered[p] {
		pointsToConsider = append(pointsToConsider, p)
		considered[p] = true
	}
	p = image.Point{X: point.X, Y: point.Y - 1}
	if point.Y-1 >= 0 && field[point.X][point.Y-1] > v && field[point.X][point.Y-1] < 9 && !considered[p] {
		pointsToConsider = append(pointsToConsider, p)
		considered[p] = true
	}
	p = image.Point{X: point.X, Y: point.Y + 1}
	if point.Y+1 < len(field[point.X]) && field[point.X][point.Y+1] > v && field[point.X][point.Y+1] < 9 && !considered[p] {
		pointsToConsider = append(pointsToConsider, p)
		considered[p] = true
	}

	return getBasinSize(field, pointsToConsider, considered, size+1)
}
