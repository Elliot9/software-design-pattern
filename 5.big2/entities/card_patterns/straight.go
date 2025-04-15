package card_patterns

import (
	"github/elliot9/big2/core"
)

type Straight struct {
	core.BaseCardPattern
}

func NewStraight() *Straight {
	s := &Straight{
		BaseCardPattern: core.BaseCardPattern{
			Name: "順子",
		},
	}
	s.CardPattern = s
	return s
}

func (s *Straight) IsValid(cards []core.Card) bool {
	if len(cards) != 5 {
		return false
	}

	exists := make([]bool, 13)
	for _, card := range cards {
		exists[int(card.Rank)] = true
	}

	temp := make([]bool, 26)
	copy(temp, exists)
	copy(temp[13:], exists)

	// 找連續的 5 個 true
	count := 0
	for _, has := range temp {
		if has {
			count++
			if count == 5 {
				return true
			}
		} else {
			count = 0
		}
	}

	return false
}

func (s *Straight) IsStrongerThan(cards []core.Card, topPlay []core.Card) bool {
	card := core.GetMaxOfCards(cards)
	topCard := core.GetMaxOfCards(topPlay)

	if card.Rank == topCard.Rank {
		return card.Suit > topCard.Suit
	}

	return card.Rank > topCard.Rank
}

var _ core.CardPattern = (*Straight)(nil)
