package card_patterns

import (
	"github/elliot9/big2/core"
)

type FullHouse struct {
	core.BaseCardPattern
}

func NewFullHouse() *FullHouse {
	f := &FullHouse{
		BaseCardPattern: core.BaseCardPattern{
			Name: "葫蘆",
		},
	}
	f.CardPattern = f
	return f
}

func (s *FullHouse) IsValid(cards []core.Card) bool {
	if len(cards) != 5 {
		return false
	}

	rankCount := make(map[core.Rank]int)
	for _, card := range cards {
		rankCount[card.Rank]++
	}

	hasThree := false
	hasTwo := false
	for _, count := range rankCount {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}

	return hasThree && hasTwo
}

func (s *FullHouse) IsStrongerThan(cards []core.Card, topPlay []core.Card) bool {
	findThreePattern := func(cards []core.Card) core.Rank {
		rankCount := make(map[core.Rank]int)
		for _, card := range cards {
			rankCount[card.Rank]++
		}

		for rank, count := range rankCount {
			if count == 3 {
				return rank
			}
		}
		return core.Rank(0)
	}

	return findThreePattern(cards) > findThreePattern(topPlay)
}

var _ core.CardPattern = (*FullHouse)(nil)
