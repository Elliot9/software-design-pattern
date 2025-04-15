package interfaces

type Troop interface {
	AddRole(role Role)
	GetRoles() []Role
	RemoveRole(role Role)
	GetNextRole() Role
	GetTroopIndex() int
	GetBattle() Battle
	SetBattle(battle Battle)
}
