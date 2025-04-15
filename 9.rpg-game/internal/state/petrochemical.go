package state

import (
	"github/elliot9/class9/internal/interfaces"
)

type Petrochemical struct {
	BaseState
}

func (n *Petrochemical) EnterState() {
	n.SetStateDuration(3)
}

func (n *Petrochemical) TakeTurn() {
	// 無法進行操作
}

func NewPetrochemical() *Petrochemical {
	state := &Petrochemical{
		BaseState: BaseState{
			Name: "石化",
		},
	}
	state.BaseState.State = state
	return state
}

var _ interfaces.State = &Petrochemical{}
