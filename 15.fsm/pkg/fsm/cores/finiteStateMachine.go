package cores

import (
	"context"
	"fmt"
)

type FiniteStateMachine struct {
	CurrentState         State
	Transitions          []Transition
	InitialStateResolver func() State
}

func NewFiniteStateMachine(initialStateResolver func() State, transitions []Transition) *FiniteStateMachine {
	return &FiniteStateMachine{InitialStateResolver: initialStateResolver, Transitions: transitions}
}

func (fsm *FiniteStateMachine) Current() State {
	if fsm.CurrentState == nil {
		fsm.CurrentState = fsm.InitialStateResolver()
	}

	return fsm.CurrentState
}

// fire event and trigger transition
func (fsm *FiniteStateMachine) SendEvent(event Event, context context.Context) {
	fsm.Current().HandleEvent(event, context)

	for _, transition := range fsm.Transitions {
		if transition.GetFrom().GetName() == fsm.Current().GetName() && transition.GetTrigger().Match(event, context) {
			fsm.Transition(transition, context)
		}
	}
}

// transition from current state to target state
// delegate transition logic to <<transition>>
func (fsm *FiniteStateMachine) Transition(transition Transition, context context.Context) error {
	if transition.GetFrom().GetName() != fsm.Current().GetName() {
		return fmt.Errorf("invalid transition: current state is %s, but transition is from %s",
			fsm.Current().GetName(), transition.GetFrom().GetName())
	}

	if transition.GetGuard() == nil || transition.GetGuard().Evaluate(context) {
		return transition.StartTransition(fsm, context)
	}

	return nil
}

func (fsm *FiniteStateMachine) AddTransitions(transitions []Transition) {
	fsm.Transitions = append(fsm.Transitions, transitions...)
}

func (fsm *FiniteStateMachine) SetInitialStateResolver(resolver func() State) {
	fsm.InitialStateResolver = resolver
}

func (fsm *FiniteStateMachine) SetCurrent(state State) {
	fsm.CurrentState = state
}
