package one_punch_handlers

import "github/elliot9/class9/internal/interfaces"

type HPOver500 struct {
	BaseOnePunchHandler
}

func (h *HPOver500) Handle(attacker, target interfaces.Role) {
	if target.GetHP() < 500 {
		h.Next(attacker, target)
	} else {
		attacker.Attack(300, target)
	}
}

func NewHPOver500() *HPOver500 {
	return &HPOver500{
		BaseOnePunchHandler: BaseOnePunchHandler{
			next: nil,
		},
	}
}

var _ interfaces.OnePunchHandler = &HPOver500{}
