package stateListeners

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type BaseStateListener struct {
	Action func(context context.Context)
}

func NewBaseStateListener(action func(context context.Context)) *BaseStateListener {
	return &BaseStateListener{Action: action}
}

func (l *BaseStateListener) Handle(context context.Context) {
	l.Action(context)
}

var _ cores.StateListener = &BaseStateListener{}
