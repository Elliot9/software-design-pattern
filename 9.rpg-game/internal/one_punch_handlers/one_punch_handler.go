package one_punch_handlers

import "github/elliot9/class9/internal/interfaces"

type BaseOnePunchHandler struct {
	next interfaces.OnePunchHandler
}

func (b *BaseOnePunchHandler) SetNext(handler interfaces.OnePunchHandler) {
	b.next = handler
}

func (b *BaseOnePunchHandler) Next(attacker, target interfaces.Role) {
	if b.next != nil {
		b.next.Handle(attacker, target)
	}
}

func (b *BaseOnePunchHandler) Handle(attacker, target interfaces.Role) {
	panic("not implemented")
}

var _ interfaces.OnePunchHandler = &BaseOnePunchHandler{}
