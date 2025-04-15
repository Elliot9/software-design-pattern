package core

import (
	"fmt"
	"sort"
)

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) string() string {
	return []string{"C", "D", "H", "S"}[s]
}

type Rank int

const (
	Three Rank = iota
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
	Two
)

func (r Rank) string() string {
	return []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}[r]
}

type Card struct {
	Suit Suit
	Rank Rank
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%s[%s]", c.Suit.string(), c.Rank.string())
}

func SortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank == cards[j].Rank {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Rank < cards[j].Rank
	})
}

func GetMaxOfCards(cards []Card) Card {
	maxCard := cards[0]
	for _, card := range cards {
		if card.Rank > maxCard.Rank {
			maxCard = card
		}
		if card.Rank == maxCard.Rank && card.Suit > maxCard.Suit {
			maxCard = card
		}
	}
	return maxCard
}
