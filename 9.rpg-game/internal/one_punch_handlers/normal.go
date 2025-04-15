package one_punch_handlers

import (
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
)

type Normal struct {
	BaseOnePunchHandler
}

func (n *Normal) Handle(attacker, target interfaces.Role) {
	_, isNormal := target.GetState().(*state.Normal)

	if isNormal {
		attacker.Attack(100, target)
	} else {
		n.Next(attacker, target)
	}
}

func NewNormal() *Normal {
	return &Normal{
		BaseOnePunchHandler: BaseOnePunchHandler{
			next: nil,
		},
	}
}

var _ interfaces.OnePunchHandler = &Normal{}
