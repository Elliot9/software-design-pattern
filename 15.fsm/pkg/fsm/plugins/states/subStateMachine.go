package states

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type SubStateMachine struct {
	Fsm  *cores.FiniteStateMachine
	Name string
}

func NewSubStateMachine(fsm *cores.FiniteStateMachine, name string) *SubStateMachine {
	return &SubStateMachine{
		Fsm:  fsm,
		Name: name,
	}
}

func (s *SubStateMachine) GetName() string {
	return s.Name
}

func (s *SubStateMachine) OnEntry(context context.Context) context.Context {
	return s.Fsm.Current().OnEntry(context)
}

func (s *SubStateMachine) OnExit(context context.Context) context.Context {
	return s.Fsm.Current().OnExit(context)
}

func (s *SubStateMachine) HandleEvent(event cores.Event, context context.Context) {
	s.Fsm.SendEvent(event, context)
}

// <<wait for implement>>
func (s *SubStateMachine) AddListeners(stateListeners map[cores.Event][]cores.StateListener) {
}

var _ cores.State = &SubStateMachine{}
