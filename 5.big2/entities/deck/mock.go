package deck

import "github/elliot9/big2/core"

type MockDeck struct {
	core.BaseDeck
}

func NewMockDeck(cards []core.Card) *MockDeck {
	m := &MockDeck{
		BaseDeck: core.BaseDeck{},
	}
	m.BaseDeck.SetCards(cards)
	return m
}
