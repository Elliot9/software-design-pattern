package triggers

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type BaseTrigger struct {
	Event cores.Event
}

func NewBaseTrigger(event cores.Event) *BaseTrigger {
	return &BaseTrigger{Event: event}
}

func (t *BaseTrigger) Match(event cores.Event, context context.Context) bool {
	return t.Event.GetName() == event.GetName()
}

var _ cores.Trigger = &BaseTrigger{}
