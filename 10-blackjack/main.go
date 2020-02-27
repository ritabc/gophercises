package main

import (
	"fmt"
	"strings"

	deck "github.com/ritabc/gophercises/09-deck"
)

// get a deck of cards with 3 decks
// shuffle deck
// draw cards
// Create player
// Create dealer

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	// create a deck and shuffle it
	cards := deck.New(deck.DeckMultiplier(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("Player:", player.String())
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	fmt.Println("==FINAL HANDDS==")
	fmt.Println("Player:", player.String())
	fmt.Println("Dealer:", dealer.String())
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
