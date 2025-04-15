package community

type Listener interface {
	OnEvent(event Event[any])
}
