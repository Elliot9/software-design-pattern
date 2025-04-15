package guards

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type OrGuard struct {
	guards []cores.Guard
}

func NewOrGuard(guards []cores.Guard) *OrGuard {
	return &OrGuard{guards: guards}
}

func (g *OrGuard) Evaluate(context context.Context) bool {
	for _, guard := range g.guards {
		if guard.Evaluate(context) {
			return true
		}
	}
	return false
}

var _ cores.Guard = &OrGuard{}
