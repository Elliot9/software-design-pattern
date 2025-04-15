package state

import (
	"github/elliot9/class9/internal/interfaces"
)

type Poisoned struct {
	BaseState
}

func (n *Poisoned) EnterState() {
	n.SetStateDuration(3)
}

func (n *Poisoned) OnTurnStart() {
	n.BaseState.OnTurnStart()
	if n.BaseState.Role.GetState() == n {
		n.BaseState.Role.SetHP(n.BaseState.Role.GetHP() - 30)
	}
}

func NewPoisoned() *Poisoned {
	state := &Poisoned{
		BaseState: BaseState{
			Name: "中毒",
		},
	}
	state.BaseState.State = state
	return state
}

var _ interfaces.State = &Petrochemical{}
