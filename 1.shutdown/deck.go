package main

import "math/rand/v2"

// 存放牌, 給予抽取
// 打亂牌堆順序
type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	// 生成 52 張牌
	cards := []Card{}
	for _, suit := range []Suit{Club, Diamond, Heart, Spade} {
		for _, rank := range []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace} {
			cards = append(cards, Card{rank: rank, suit: suit})
		}
	}
	return &Deck{cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) GetCards() []Card {
	return d.cards
}
