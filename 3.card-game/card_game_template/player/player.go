package player

import "github/elliot9/card_game_template/card"

type Player interface {
	GetName() string
	GetHands() []card.CardStrategy
	AddHands(card card.CardStrategy)
	Show() card.CardStrategy
	NameSelf()
}

type BasePlayer struct {
	hands []card.CardStrategy
	name  string
}

func (p *BasePlayer) setName(name string) {
	p.name = name
}

func (p *BasePlayer) GetName() string {
	return p.name
}

func (p *BasePlayer) setHands(cards []card.CardStrategy) {
	p.hands = cards
}

func (p *BasePlayer) GetHands() []card.CardStrategy {
	return p.hands
}

func (p *BasePlayer) AddHands(card card.CardStrategy) {
	p.hands = append(p.hands, card)
}
