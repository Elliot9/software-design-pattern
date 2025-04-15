package card_patterns

import "github/elliot9/big2/core"

type Single struct {
	core.BaseCardPattern
}

func NewSingle() *Single {
	s := &Single{
		BaseCardPattern: core.BaseCardPattern{
			Name: "單張",
		},
	}
	s.CardPattern = s
	return s
}

func (s *Single) IsValid(cards []core.Card) bool {
	return len(cards) == 1
}

func (s *Single) IsStrongerThan(cards []core.Card, topPlay []core.Card) bool {
	card := cards[0]
	topCard := topPlay[0]

	if card.Rank == topCard.Rank {
		return card.Suit > topCard.Suit
	}

	return card.Rank > topCard.Rank
}

var _ core.CardPattern = (*Single)(nil)
