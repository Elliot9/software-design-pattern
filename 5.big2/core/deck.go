package core

type Deck interface {
	Draw() Card
	IsEmpty() bool
}

type BaseDeck struct {
	cards []Card
}

func (b *BaseDeck) Draw() Card {
	card := b.cards[0]
	b.cards = b.cards[1:]
	return card
}

func (b *BaseDeck) IsEmpty() bool {
	return len(b.cards) == 0
}

func (b *BaseDeck) SetCards(cards []Card) {
	b.cards = cards
}

func (b *BaseDeck) GetCards() []Card {
	return b.cards
}
