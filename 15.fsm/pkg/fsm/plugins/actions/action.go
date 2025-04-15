package actions

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type BaseAction struct {
	Action func(context context.Context) context.Context
}

func NewBaseAction(action func(context context.Context) context.Context) *BaseAction {
	return &BaseAction{Action: action}
}

func (a *BaseAction) Execute(context context.Context) context.Context {
	return a.Action(context)
}

var _ cores.Action = &BaseAction{}
