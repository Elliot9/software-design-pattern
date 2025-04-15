package guards

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type BaseGuard struct {
	Condition func(context context.Context) bool
}

func NewBaseGuard(condition func(context context.Context) bool) *BaseGuard {
	return &BaseGuard{Condition: condition}
}

func (g *BaseGuard) Evaluate(context context.Context) bool {
	return g.Condition(context)
}

var _ cores.Guard = &BaseGuard{}
