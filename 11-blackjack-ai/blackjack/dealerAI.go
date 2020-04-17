package blackjack

import deck "github.com/ritabc/gophercises/09-deck"

func DealerAI() AI {
	return dealerAI{}
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
	// noop
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealerCard deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	} else {
		return MoveStand
	}
}

func (ai dealerAI) Results(hands [][]deck.Card, dealerHand []deck.Card) {
	// noop

}
