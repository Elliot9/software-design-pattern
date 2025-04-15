package card

import "fmt"

type PokerCard struct {
	rank Rank
	suit Suit
}

func NewPokerCard(rank Rank, suit Suit) *PokerCard {
	return &PokerCard{rank: rank, suit: suit}
}

func (p *PokerCard) Identify() string {
	return fmt.Sprintf("%s-%s", p.rank.string(), p.suit.string())
}

type Rank int

func (r Rank) string() string {
	return []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[r]
}

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) string() string {
	return []string{"♣", "♦", "♥", "♠"}[s]
}
