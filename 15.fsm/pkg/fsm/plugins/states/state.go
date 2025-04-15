package states

import (
	"context"
	"github/elliot9/class15/pkg/fsm/cores"
)

type BaseState struct {
	Name      string
	Listeners map[string][]cores.StateListener
}

func NewBaseState(name string) *BaseState {
	return &BaseState{
		Name:      name,
		Listeners: make(map[string][]cores.StateListener),
	}
}

func (s *BaseState) GetName() string {
	return s.Name
}

// <<wait for implement>>
func (s *BaseState) OnEntry(context context.Context) context.Context {
	return context
}

// <<wait for implement>>
func (s *BaseState) OnExit(context context.Context) context.Context {
	return context
}

func (s *BaseState) HandleEvent(event cores.Event, context context.Context) {
	if listeners, ok := s.Listeners[event.GetName()]; ok {
		for _, listener := range listeners {
			listener.Handle(context)
		}
	}
}

func (s *BaseState) AddListeners(stateListeners map[cores.Event][]cores.StateListener) {
	for event, listeners := range stateListeners {
		if s.Listeners == nil {
			s.Listeners = make(map[string][]cores.StateListener)
		}
		s.Listeners[event.GetName()] = append(s.Listeners[event.GetName()], listeners...)
	}
}

var _ cores.State = &BaseState{}
