package main

// 做出遊戲決策
type Player interface {
	GetHands() []Card
	GetName() string
	AddHands(card Card)
	Show() *Card
	NameSelf()
}

type BasePlayer struct {
	hands []Card
	name  string
}

func (p *BasePlayer) GetHands() []Card {
	return p.hands
}

func (p *BasePlayer) setName(name string) {
	p.name = name
}

func (p *BasePlayer) GetName() string {
	return p.name
}

func (p *BasePlayer) AddHands(card Card) {
	p.hands = append(p.hands, card)
}

func (p *BasePlayer) setCards(cards []Card) {
	p.hands = cards
}
