package main

import "math/rand/v2"

// 存放牌, 給予抽取
// 打亂牌堆順序
type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	// 生成 40 張牌
	cards := []Card{}
	for _, color := range []Color{Red, Yellow, Green, Blue} {
		for _, number := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			cards = append(cards, Card{number: number, color: color})
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

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *Deck) AddCards(cards []Card) {
	d.cards = append(d.cards, cards...)
}
