package one_punch_handlers

import (
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
)

type CheerUp struct {
	BaseOnePunchHandler
}

func (c *CheerUp) Handle(attacker, target interfaces.Role) {
	_, isCheerUp := target.GetState().(*state.CheerUp)

	if isCheerUp {
		attacker.Attack(100, target)
		target.SetState(state.NewNormal())
	} else {
		c.Next(attacker, target)
	}
}

func NewCheerUp() *CheerUp {
	return &CheerUp{
		BaseOnePunchHandler: BaseOnePunchHandler{
			next: nil,
		},
	}
}

var _ interfaces.OnePunchHandler = &CheerUp{}
