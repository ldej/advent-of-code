package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/queue"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day22/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day22/input.txt")
	fmt.Println("Result:", result)
}

type Player struct {
	Number int
	Cards  queue.IntQueue
}

func run(input string) int {
	lines := tools.ReadStringsDoubleNewlines(input)
	var players []Player
	for _, line := range lines {
		parts := strings.Split(line, ":\n")
		players = append(players, Player{
			Number: tools.FindInt(parts[0]),
			Cards:  tools.FindInts(parts[1]),
		})
	}

	for !players[0].Cards.IsEmpty() && !players[1].Cards.IsEmpty() {
		card1, _ := players[0].Cards.Pop()
		card2, _ := players[1].Cards.Pop()
		if card1 > card2 {
			players[0].Cards.Push(card1)
			players[0].Cards.Push(card2)
		} else {
			players[1].Cards.Push(card2)
			players[1].Cards.Push(card1)
		}
	}

	var winner Player
	if players[0].Cards.IsEmpty() {
		winner = players[1]
	} else {
		winner = players[0]
	}

	var score int
	for i := len(winner.Cards); i > 0; i-- {
		card, _ := winner.Cards.Pop()
		score += card * i
	}

	return score
}
