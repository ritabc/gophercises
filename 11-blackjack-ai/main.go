package main

import (
	"fmt"

	"github.com/ritabc/gophercises/11-blackjack-ai/blackjack"
)

type AI struct{}

func main() {

	opts := blackjack.Options{
		Hands:           2,
		Decks:           3,
		BlackjackPayout: 2,
	}

	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())

	fmt.Println("Our AI won/lost:", winnings)
}
