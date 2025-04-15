package cores

import (
	"context"
)

type State interface {
	GetName() string
	OnEntry(context context.Context) context.Context
	OnExit(context context.Context) context.Context
	HandleEvent(event Event, context context.Context)
	AddListeners(stateListeners map[Event][]StateListener)
}
