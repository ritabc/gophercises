package blackjack

import (
	"errors"

	deck "github.com/ritabc/gophercises/09-deck"
)

type state int8

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout int
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: DealerAI(),
		balance:  0,
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BlackjackPayout == 0 {
		opts.BlackjackPayout = 2
	}
	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout
	return g
}

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player[g.handIdx].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

type hand struct {
	cards []deck.Card
	bet   int
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	g.playerBet = bet
}

func deal(g *Game) {
	playerHand := make([]deck.Card, 0, 5)
	g.handIdx = 0
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		playerHand = append(playerHand, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.player = []hand{
		{
			cards: playerHand,
			bet:   g.playerBet,
		},
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = nil

	// Reshuffle after once get to 1/4 of deck
	min := 52 * g.nDecks / 4
	for i := 0; i < g.nHands; i++ {
		shuffled := false
		if len(g.deck) < min {
			g.deck = deck.New(deck.DeckMultiplier(g.nDecks), deck.Shuffle)
			shuffled = true
		}
		bet(g, ai, shuffled)
		deal(g)
		if Blackjack(g.dealer...) {
			endRound(g, ai)
			continue
		}
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(*g.currentHand()))
			copy(hand, *g.currentHand())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			if err != nil {
				switch err {
				case errBust:
					MoveStand(g)
				default:
					panic(err)
				}
			}
		}
		// If dealer score <= 16, we hit
		// If dealer has a soft 17 (score == 17 && minScore == 7)
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endRound(g, ai)
	}
	return g.balance

}

type Game struct {

	// unexported fields
	nDecks          int
	nHands          int
	blackjackPayout int

	state state
	deck  []deck.Card

	player    []hand
	handIdx   int
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

var errBust = errors.New("hand score over 21")

type Move func(*Game) error

func MoveHit(g *Game) error {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
	return nil
}

func MoveSplit(g *Game) error {
	cards := g.currentHand()
	if len(*cards) != 2 {
		return errors.New("you can only split with two cards in your hand")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("both cards must have the same rank to split")
	}
	g.player = append(g.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   g.player[g.handIdx].bet,
	})
	g.player[g.handIdx].cards = (*cards)[:1]
	return nil
}

func MoveDouble(g *Game) error {
	if len(*g.currentHand()) != 2 {
		return errors.New("can only double on a hand with 2 cards")
	}
	g.playerBet *= 2
	MoveHit(g)
	return MoveStand(g)
}

func MoveStand(g *Game) error {
	if g.state == stateDealerTurn {
		g.state++
		return nil
	}
	if g.state == statePlayerTurn {

		g.handIdx++
		if g.handIdx == len(g.player) {
			g.state++
		}
		return nil
	}
	return errors.New("invalid state")
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// Score takes in a hand of cards and returns the best blackjack score possible with the hand
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to be worth 11
			return minScore + 10
		}
	}
	return minScore
}

// Soft returns true if an Ace is counting as 11 points
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

// Blackjack returns true if a hand is a blackjack
func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

// Assume all Aces are 1 point
func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

// Useful to determine that J/Q/Ks which have rank of 11/12/13, actually have score of 10 in BlackJack
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func endRound(g *Game, playerAI AI) {
	dScore := Score(g.dealer...)
	dBlackjack := Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))
	for hi, hand := range g.player {
		cards := hand.cards
		allHands[hi] = cards
		winnings := hand.bet
		pScore, pBlackjack := Score(cards...), Blackjack(cards...)
		switch {
		case pBlackjack && dBlackjack:
			winnings = 0
		case dBlackjack:
			winnings *= -1
		case pBlackjack:
			winnings = int(winnings * g.blackjackPayout)
		case pScore > 21:
			winnings *= -1
		case dScore > 21:
			// win
		case pScore > dScore:
			// win
		case dScore > pScore:
			winnings *= -1
		case dScore == pScore:
			winnings = 0
		}
		g.balance += winnings
	}
	playerAI.Results(allHands, g.dealer)
	g.player = nil
	g.dealer = nil
}
