package transitions

import (
	"context"
	"fmt"
	"github/elliot9/class15/pkg/fsm/cores"
)

type BaseTransition struct {
	Trigger cores.Trigger
	Guard   cores.Guard
	From    cores.State
	To      cores.State
	Actions []cores.Action
}

func NewBaseTransition(trigger cores.Trigger, guard cores.Guard, from cores.State, to cores.State, actions []cores.Action) *BaseTransition {
	return &BaseTransition{Trigger: trigger, Guard: guard, From: from, To: to, Actions: actions}
}

func (t *BaseTransition) GetTrigger() cores.Trigger {
	return t.Trigger
}

func (t *BaseTransition) GetGuard() cores.Guard {
	return t.Guard
}

func (t *BaseTransition) GetFrom() cores.State {
	return t.From
}

func (t *BaseTransition) GetTo() cores.State {
	return t.To
}

func (t *BaseTransition) GetActions() []cores.Action {
	return t.Actions
}

func (t *BaseTransition) StartTransition(fsm *cores.FiniteStateMachine, context context.Context) error {
	if t.GetFrom().GetName() != fsm.Current().GetName() {
		return fmt.Errorf("invalid transition: current state is %s, but transition is from %s",
			fsm.Current().GetName(), t.GetFrom().GetName())
	}

	if t.GetGuard() == nil || t.GetGuard().Evaluate(context) {
		context = fsm.Current().OnExit(context)

		for _, action := range t.Actions {
			context = action.Execute(context)
		}

		fsm.SetCurrent(t.To)
		fsm.Current().OnEntry(context)
	}

	return nil
}

var _ cores.Transition = &BaseTransition{}
