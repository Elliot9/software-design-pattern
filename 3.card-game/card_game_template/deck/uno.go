package deck

import "github/elliot9/card_game_template/card"

type UnoDeck struct {
	BaseDeck
}

func (d *UnoDeck) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *UnoDeck) AddCards(cards []card.CardStrategy) {
	d.cards = append(d.cards, cards...)
	d.Shuffle()
}

func NewUnoDeck(cards []card.CardStrategy) *UnoDeck {
	return &UnoDeck{BaseDeck: BaseDeck{cards: cards}}
}

var _ Deck = &UnoDeck{}
