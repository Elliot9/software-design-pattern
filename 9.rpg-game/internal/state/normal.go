package state

import (
	"github/elliot9/class9/internal/interfaces"
)

type Normal struct {
	BaseState
}

func (n *Normal) EnterState() {
	n.SetStateDuration(0)
}

func NewNormal() *Normal {
	state := &Normal{
		BaseState: BaseState{
			Name: "正常",
		},
	}
	state.BaseState.State = state
	return state
}

var _ interfaces.State = &Normal{}
