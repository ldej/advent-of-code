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
	field := tools.ReadIntGrid("./2021/day11/input.txt")
	totalFlashes := 0
	for i := 0; i < 100; i++ {
		newField, flashes := step(field)
		field = newField
		totalFlashes += flashes
	}
	return totalFlashes
}

func step(field tools.IntGrid) (tools.IntGrid, int) {
	newField := field.Copy()
	newFlash := false
	for cell := range field.Cells() {
		newField[cell.X][cell.Y] += 1
		if newField[cell.X][cell.Y] > 9 {
			newFlash = true
		}
	}
	if !newFlash {
		return newField, 0
	}

	hasFlashed := field.Copy()

	for newFlash {
		newFlash = false
		for cell := range newField.Cells() {
			if cell.Value > 9 && hasFlashed[cell.X][cell.Y] != -1 {
				hasFlashed[cell.X][cell.Y] = -1

				for i := myints.Max(cell.X-1, 0); i <= myints.Min(cell.X+1, len(newField)-1); i++ {
					for j := myints.Max(cell.Y-1, 0); j <= myints.Min(cell.Y+1, len(newField[cell.X])-1); j++ {
						if !(i == cell.X && j == cell.Y) {
							newField[i][j] += 1
							if newField[i][j] > 9 {
								newFlash = true
							}
						}
					}
				}
			}
		}
	}

	flashes := 0
	for cell := range hasFlashed.Cells() {
		if cell.Value == -1 {
			flashes += 1
			newField[cell.X][cell.Y] = 0
		}
	}
	return newField, flashes
}
