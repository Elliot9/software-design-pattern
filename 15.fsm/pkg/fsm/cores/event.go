package cores

type Event interface {
	GetName() string
}

func NewEvent(name string) Event {
	return &BaseEvent{Name: name}
}

type BaseEvent struct {
	Name string
}

func (e *BaseEvent) GetName() string {
	return e.Name
}

var _ Event = &BaseEvent{}
