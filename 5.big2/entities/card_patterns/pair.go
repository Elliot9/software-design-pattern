package card_patterns

import "github/elliot9/big2/core"

type Pair struct {
	core.BaseCardPattern
}

func NewPair() *Pair {
	p := &Pair{
		BaseCardPattern: core.BaseCardPattern{
			Name: "對子",
		},
	}
	p.CardPattern = p
	return p
}

func (s *Pair) IsValid(cards []core.Card) bool {
	return len(cards) == 2 && cards[0].Rank == cards[1].Rank
}

func (s *Pair) IsStrongerThan(cards []core.Card, topPlay []core.Card) bool {
	card := core.GetMaxOfCards(cards)
	topCard := core.GetMaxOfCards(topPlay)

	if card.Rank == topCard.Rank {
		return card.Suit > topCard.Suit
	}

	return card.Rank > topCard.Rank
}

var _ core.CardPattern = (*Pair)(nil)
