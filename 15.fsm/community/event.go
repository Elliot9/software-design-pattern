package community

type EventType string

const (
	NewMessage        EventType = "NewMessage"
	NewPost           EventType = "NewPost"
	NewComment        EventType = "NewComment"
	NewSpeak          EventType = "NewNewSpeak"
	GoBroadcasting    EventType = "GoBroadcasting"
	StopBroadcasting  EventType = "StopBroadcasting"
	OnlineUserChanged EventType = "OnlineUserChanged"
)

type Event[T any] struct {
	eventType EventType
	data      T
}

func NewEvent[T any](eventType EventType, data T) Event[T] {
	return Event[T]{
		eventType: eventType,
		data:      data,
	}
}

func (e *Event[T]) GetType() EventType {
	return e.eventType
}

func (e *Event[T]) GetData() T {
	return e.data
}
