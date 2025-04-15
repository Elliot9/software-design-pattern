package state

import (
	"github/elliot9/class9/internal/interfaces"
)

type CheerUp struct {
	BaseState
}

func (n *CheerUp) EnterState() {
	n.SetStateDuration(3)
}

func (n *CheerUp) Attack(damage int, target interfaces.Role) {
	n.BaseState.Attack(damage+50, target)
}

func NewCheerUp() *CheerUp {
	state := &CheerUp{
		BaseState: BaseState{
			Name: "受到鼓舞",
		},
	}
	state.BaseState.State = state
	return state
}

var _ interfaces.State = &CheerUp{}
