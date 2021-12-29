package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}

type Line []image.Point

func run() int {
	inputs := tools.ReadStrings()

	var lines []Line
	for _, input := range inputs {
		parts := strings.Split(input, " -> ")
		line := Line{}
		for _, part := range parts {
			values := myints.ParseCsv(part)[0]
			line = append(line, image.Point{X: values[0], Y: values[1]})
		}
		lines = append(lines, line)
	}

	field := map[int]map[int]int{}
	for _, line := range lines {
		if line[0].X == line[1].X {
			if line[0].Y < line[1].Y {
				for i := line[0].Y; i <= line[1].Y; i++ {
					if field[line[0].X] == nil {
						field[line[0].X] = map[int]int{}
					}
					field[line[0].X][i]++
				}
			} else {
				for i := line[1].Y; i <= line[0].Y; i++ {
					if field[line[0].X] == nil {
						field[line[0].X] = map[int]int{}
					}
					field[line[0].X][i]++
				}
			}
		}
		if line[0].Y == line[1].Y {
			if line[0].X < line[1].X {
				for i := line[0].X; i <= line[1].X; i++ {
					if field[i] == nil {
						field[i] = map[int]int{}
					}
					field[i][line[0].Y]++
				}
			} else {
				for i := line[1].X; i <= line[0].X; i++ {
					if field[i] == nil {
						field[i] = map[int]int{}
					}
					field[i][line[0].Y]++
				}
			}
		}
	}

	count := 0
	for _, column := range field {
		for _, value := range column {
			if value >= 2 {
				count++
			}
		}
	}

	return count
}
