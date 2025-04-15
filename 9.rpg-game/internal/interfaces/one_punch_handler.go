package interfaces

type OnePunchHandler interface {
	SetNext(handler OnePunchHandler)
	Next(attacker, target Role)
	Handle(attacker, target Role)
}
