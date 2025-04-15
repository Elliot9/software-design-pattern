package guards

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type AndGuard struct {
	guards []cores.Guard
}

func NewAndGuard(guards []cores.Guard) *AndGuard {
	return &AndGuard{guards: guards}
}

func (g *AndGuard) Evaluate(context context.Context) bool {
	for _, guard := range g.guards {
		if !guard.Evaluate(context) {
			return false
		}
	}
	return true
}

var _ cores.Guard = &AndGuard{}
