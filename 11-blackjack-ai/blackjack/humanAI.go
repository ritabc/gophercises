package blackjack

import (
	"fmt"

	deck "github.com/ritabc/gophercises/09-deck"
)

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(playerHand []deck.Card, dealerCard deck.Card) Move {
	for {
		fmt.Println("Player:", playerHand)
		fmt.Println("Dealer:", dealerCard)
		fmt.Println("What will you do? (h)it, (s)tand, (d)ouble, s(p)lit")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid Option", input)
		}
	}
}

func (ai humanAI) Results(hands [][]deck.Card, dealerHand []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	// TODO: print scores
	fmt.Println("Dealer:", dealerHand)
}
