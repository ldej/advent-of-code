package main

import (
	"github.com/ldej/advent-of-code/2018/common"
		"strconv"
		"strings"
	"fmt"
)

type point struct {
	X int
	Y int
	DX int
	DY int
}

func main() {
	results := common.ReadAllLines("./day10/input.txt", `position=<(?P<x>.*?), (?P<y>.*?)> velocity=<(?P<dx>.*?), (?P<dy>.*)>`)

	points := []*point{}
	minx := 100000
	maxx := 0
	miny := 100000
	maxy := 0
	for _, result := range results {
		x, _ := strconv.Atoi(strings.TrimSpace(result["x"]))
		y, _ := strconv.Atoi(strings.TrimSpace(result["y"]))
		dx, _ := strconv.Atoi(strings.TrimSpace(result["dx"]))
		dy, _ := strconv.Atoi(strings.TrimSpace(result["dy"]))

		if x < minx {
			minx = x
		} else if x > maxx {
			maxx = x
		}
		if y < miny {
			miny = y
		} else if y > maxy {
			maxy = y
		}

		points = append(points, &point{
			X: x,
			Y: y,
			DX: dx,
			DY: dy,
		})
	}

	// TODO fail
	for i := 0; i < 100000; i ++ {
		if i % 1000 == 0 {
			printPoints(points, minx, maxx, miny, maxy)
		}
		minx, maxx, miny, maxy = movePoints(points)
		//movePoints(points)
	}
}

func printPoints(points []*point, minx, maxx, miny, maxy int) {
	xpoint := (maxx - minx) / 600.0
	ypoint := (maxy - miny) / 200.0

	toPrint := [][]string{}

	for y := miny; y < maxy; y += ypoint {
		line := []string{}
		for x := minx; x < maxx; x += xpoint {
			printed := false
			for _, point := range points {
				if point.X > x && point.X <= x + xpoint && point.Y > y && point.Y <= y + ypoint {
					if !printed {
						line = append(line, "X")
					}
					printed = true
				}
			}
			if !printed {
				line = append(line, ".")
			}
		}
		toPrint = append(toPrint, line)
	}

	total := ""
	for _, line := range toPrint {
		total += strings.Join(line, "") + "\n"
	}
	fmt.Println(total)
}

func movePoints(points []*point) (int, int, int, int){
	minx := 100000
	maxx := 0
	miny := 100000
	maxy := 0

	for _, p := range points {
		p.X += p.DX
		p.Y += p.DY

		if p.X < minx {
			minx = p.X
		} else if p.X > maxx {
			maxx = p.X
		}
		if p.Y < miny {
			miny = p.Y
		} else if p.Y > maxy {
			maxy = p.Y
		}
	}
	return minx, maxx, miny, maxy
}
