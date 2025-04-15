package deck

import (
	"github/elliot9/big2/core"
	"math/rand/v2"
)

type NormalDeck struct {
	core.BaseDeck
}

func NewNormalDeck() *NormalDeck {
	n := &NormalDeck{}
	n.BaseDeck = core.BaseDeck{}

	cards := make([]core.Card, 0)
	for _, rank := range []core.Rank{core.Three, core.Four, core.Five, core.Six, core.Seven, core.Eight, core.Nine, core.Ten, core.Jack, core.Queen, core.King, core.Ace, core.Two} {
		for _, suit := range []core.Suit{core.Club, core.Diamond, core.Heart, core.Spade} {
			cards = append(cards, core.NewCard(rank, suit))
		}
	}

	// 洗牌
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	n.BaseDeck.SetCards(cards)
	return n
}
