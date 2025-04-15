package core

import (
	"slices"
)

type Player interface {
	Play() []Card
	AddHands(cards []Card)
	NameSelf()
	GetName() string
	GetHands() []Card
	PlayCards() []Card
}

type BasePlayer struct {
	name  string
	hands []Card
	Player
}

func (p *BasePlayer) NameSelf() {
	panic("player name self not implemented")
}

func (p *BasePlayer) SetHands(cards []Card) {
	p.hands = cards
	SortCards(p.hands)
}

func (p *BasePlayer) GetHands() []Card {
	return p.hands
}

func (p *BasePlayer) SetName(name string) {
	p.name = name
}

func (p *BasePlayer) GetName() string {
	return p.name
}

func (p *BasePlayer) AddHands(cards []Card) {
	p.hands = append(p.hands, cards...)
	SortCards(p.hands)
}

func (p *BasePlayer) Play() []Card {
	cards := p.Player.PlayCards()
	if len(cards) == 0 {
		return cards
	}

	// 剩餘手牌
	newHands := []Card{}
	for _, hand := range p.GetHands() {
		if !slices.Contains(cards, hand) {
			newHands = append(newHands, hand)
		}
	}

	// 更新玩家的手牌
	p.SetHands(newHands)
	return cards
}

func (p *BasePlayer) PlayCards() []Card {
	panic("player play cards not implemented")
}

var _ Player = &BasePlayer{}
