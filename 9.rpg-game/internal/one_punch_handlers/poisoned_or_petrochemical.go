package one_punch_handlers

import (
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
)

type PoisonedOrPetrochemical struct {
	BaseOnePunchHandler
}

func (p *PoisonedOrPetrochemical) Handle(attacker, target interfaces.Role) {
	_, isPoisoned := target.GetState().(*state.Poisoned)
	_, isPetrochemical := target.GetState().(*state.Petrochemical)

	if isPoisoned || isPetrochemical {
		for i := 0; i < 3; i++ {
			if target.IsAlive() {
				attacker.Attack(80, target)
			}
		}
	} else {
		p.Next(attacker, target)
	}
}

func NewPoisonedOrPetrochemical() *PoisonedOrPetrochemical {
	return &PoisonedOrPetrochemical{
		BaseOnePunchHandler: BaseOnePunchHandler{
			next: nil,
		},
	}
}

var _ interfaces.OnePunchHandler = &PoisonedOrPetrochemical{}
