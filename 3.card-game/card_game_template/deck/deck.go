package deck

import (
	"github/elliot9/card_game_template/card"
	"math/rand/v2"
)

type Deck interface {
	Shuffle()
	Draw() card.CardStrategy
}

type BaseDeck struct {
	cards []card.CardStrategy
}

func (d *BaseDeck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *BaseDeck) Draw() card.CardStrategy {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

var _ Deck = &BaseDeck{}
