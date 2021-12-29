package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/queue"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day22/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day22/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStringsDoubleNewlines()
	var player1 = tools.FindInts(strings.Split(lines[0], ":\n")[1])
	var player2 = tools.FindInts(strings.Split(lines[1], ":\n")[1])

	_, cards := PlayGame(player1, player2, 1, 1)

	var score int
	for i := len(cards); i > 0; i-- {
		card, _ := cards.Pop()
		score += card * i
	}

	return score
}

var subGameCounter = 1

func PlayGame(player1 queue.IntQueue, player2 queue.IntQueue, game int, lastGame int) (int, queue.IntQueue) {
	var history1 = make(map[string]bool)
	var history2 = make(map[string]bool)

	if game > 1 && myints.Max(player1...) > myints.Max(player2...) && myints.Max(player1...) > len(player1)+len(player2)-2 {
		// https://www.reddit.com/r/adventofcode/comments/khyjgv/2020_day_22_solutions/ggpcsnd/
		fmt.Println("Sub-game cheat, player 1 wins!")
		return 0, player1
	}

	fmt.Printf("=== Game %d ===\n\n", game)

	var round = 1
	for !player1.IsEmpty() && !player2.IsEmpty() {

		fmt.Printf("-- Round %d (Game %d) --\n", round, game)

		cards1 := myints.ToCsv(player1)
		cards2 := myints.ToCsv(player2)
		fmt.Printf("Player 1's deck: %s\n", cards1)
		fmt.Printf("Player 2's deck: %s\n", cards2)
		_, found1 := history1[cards1]
		_, found2 := history2[cards2]
		if found1 || found2 {
			fmt.Println("Seen these cards before, player 1 wins")
			return 0, player1
		}
		history1[cards1] = true
		history2[cards2] = true

		card1, _ := player1.Pop()
		card2, _ := player2.Pop()
		fmt.Printf("Player 1 plays: %d\n", card1)
		fmt.Printf("Player 2 plays: %d\n", card2)

		var winner int
		if len(player1) >= card1 && len(player2) >= card2 {
			fmt.Printf("Playing a sub-game to determine the winner...\n\n")
			subGameCounter++
			winner, _ = PlayGame(player1.Copy()[:card1], player2.Copy()[:card2], subGameCounter, game)
		} else if card1 > card2 {
			winner = 0
		} else {
			winner = 1
		}

		fmt.Printf("Player %d wins round %d of game %d!\n\n", winner+1, round, game)
		if winner == 0 {
			player1.Push(card1, card2)
		} else {
			player2.Push(card2, card1)
		}
		round++
	}

	var winner int
	var cards queue.IntQueue
	if !player1.IsEmpty() {
		winner = 0
		cards = player1
	} else {
		winner = 1
		cards = player2
	}
	fmt.Printf("The winner of game %d is player %d!\n", game, winner+1)
	if game > 1 {
		fmt.Printf("...anyway, back to game %d.\n", lastGame)
	}
	return winner, cards
}
