package deck

import "github/elliot9/card_game_template/card"

type ShowdownDeck struct {
	BaseDeck
}

func NewShowdownDeck(cards []card.CardStrategy) *ShowdownDeck {
	return &ShowdownDeck{BaseDeck: BaseDeck{cards: cards}}
}

var _ Deck = &ShowdownDeck{}
