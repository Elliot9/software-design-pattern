package interfaces

type Action interface {
	GetName() string
	GetMPCost() int
	Excute(role Role, targets []Role)
	GetTargetCount() int
	GetTargetType() TargetType
}

type TargetType int

const (
	TargetEnemy TargetType = iota
	TargetEnemyAll
	TargetAllyWithoutSelf
	TargetSelf
	TargetAllWithoutSelf
	TargetNil
)
