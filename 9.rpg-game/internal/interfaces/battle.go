package interfaces

type Battle interface {
	Start()
	IsGameOver() bool
	GetEnemyTroop(troop Troop) Troop
	PrintResult()
}
