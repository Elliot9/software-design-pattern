package cores

import "context"

type Transition interface {
	StartTransition(fsm *FiniteStateMachine, context context.Context) error
	GetFrom() State
	GetTo() State
	GetTrigger() Trigger
	GetGuard() Guard
	GetActions() []Action
}
