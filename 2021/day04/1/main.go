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

type Board [][]map[int]bool

func run() int {
	input := tools.ReadStringsDoubleNewlines("./2021/day04/input.txt")
	valuesToDraw, boardsInput := myints.ParseCsv(input[0])[0], input[1:]

	var boards []Board

	for _, b := range boardsInput {
		lines := strings.Split(b, "\n")

		board := Board{}
		for _, line := range lines {
			row := make([]map[int]bool, 0)
			numbers := myints.ToInts(strings.Fields(line))
			for _, number := range numbers {
				row = append(row, map[int]bool{number: false})
			}
			board = append(board, row)
		}
		boards = append(boards, board)
	}

	for _, value := range valuesToDraw {
		boards = UpdateBoards(boards, value)
		if winningBoard := GetWinningBoard(boards); winningBoard != nil {
			sum := SumUnmarkedNumbers(*winningBoard)
			return sum * value
		}
	}

	return -1
}

func UpdateBoards(boards []Board, toUpdate int) []Board {
	var newBoards []Board
	for _, board := range boards {
		newBoard := Board{}
		for _, row := range board {
			var newRow []map[int]bool
			for _, number := range row {
				for key, value := range number {
					if key == toUpdate {
						newRow = append(newRow, map[int]bool{key: true})
					} else {
						newRow = append(newRow, map[int]bool{key: value})
					}
				}
			}
			newBoard = append(newBoard, newRow)
		}
		newBoards = append(newBoards, newBoard)
	}
	return newBoards
}

func GetWinningBoard(boards []Board) *Board {
	for _, board := range boards {
		for _, row := range board {
			if RowWins(row) {
				return &board
			}
		}

		for i := 0; i < len(board[0]); i++ {
			if ColumnWins(board, i) {
				return &board
			}
		}
	}
	return nil
}

func RowWins(row []map[int]bool) bool {
	for _, number := range row {
		for _, value := range number {
			if !value {
				return false
			}
		}
	}
	return true
}

func ColumnWins(board Board, index int) bool {
	for _, row := range board {
		for _, value := range row[index] {
			if !value {
				return false
			}
		}
	}
	return true
}

func SumUnmarkedNumbers(board Board) int {
	sum := 0
	for _, row := range board {
		for _, number := range row {
			for key, value := range number {
				if !value {
					sum += key
				}
			}
		}
	}
	return sum
}
