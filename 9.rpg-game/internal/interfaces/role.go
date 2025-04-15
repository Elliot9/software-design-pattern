package interfaces

import "github/elliot9/class9/infra/cli"

type Role interface {
	Attack(damage int, target Role)
	MakeDecisionImpl() Action
	TakeTurn()
	SelectTargets(targets []Role, count int) []Role
	Heal(heal int)
	GetState() State
	SetState(state State)
	IsAlive() bool
	GetHP() int
	GetMP() int
	GetSTR() int
	GetName() string
	SetHP(hp int)
	SetMP(mp int)
	SetTroop(troop Troop)
	GetTroop() Troop
	GetCLI() cli.CLI
	MakeDecision() Action
	GetTargets(targetType TargetType, count int) []Role
	AddCursed(from Role)
}

type AIStrategy interface {
	MakeDecisionImpl() Action
	SelectTargets(targets []Role, count int) []Role
}
