package interfaces

type State interface {
	OnTurnStart()
	TakeTurn()
	OnTurnEnd()
	Attack(damage int, target Role)
	EnterState()
	ExitState()
	SetRole(role Role)
	GetName() string
	GetRole() Role
}
