package blackjack

import (
	deck "github.com/ritabc/gophercises/09-deck"
)

type AI interface {
	Play(playerHand []deck.Card, dealerCard deck.Card) Move
	Bet(shuffled bool) int
	Results(hands [][]deck.Card, dealerHand []deck.Card)
}
