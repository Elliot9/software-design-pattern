package player

import "github/elliot9/big2/core"

type AI struct {
	core.BasePlayer
}

func NewAI() *AI {
	ai := &AI{}
	ai.Player = ai
	return ai
}

func (a *AI) PlayCards() []core.Card {
	panic("ai 出牌尚未實作")
}

func (a *AI) NameSelf() {
	panic("ai 取名尚未實作")
}

var _ core.Player = (*AI)(nil)
