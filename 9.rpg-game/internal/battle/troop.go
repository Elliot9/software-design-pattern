package battle

import (
	"github/elliot9/class9/internal/interfaces"
)

type BaseTroop struct {
	Roles  []interfaces.Role
	seed   int
	index  int
	battle interfaces.Battle
}

func NewTroop(index int) *BaseTroop {
	return &BaseTroop{
		Roles: []interfaces.Role{},
		seed:  0,
		index: index,
	}
}

func (t *BaseTroop) AddRole(role interfaces.Role) {
	t.Roles = append(t.Roles, role)
	role.SetTroop(t)
}

func (t *BaseTroop) GetRoles() []interfaces.Role {
	roles := make([]interfaces.Role, len(t.Roles))
	copy(roles, t.Roles)
	return roles
}

func (t *BaseTroop) RemoveRole(role interfaces.Role) {
	for i, r := range t.Roles {
		if r == role {
			t.Roles = append(t.Roles[:i], t.Roles[i+1:]...)
			t.seed = max(0, t.seed-1)
			break
		}
	}
}

func (t *BaseTroop) GetNextRole() interfaces.Role {
	if t.seed == len(t.GetRoles()) {
		t.seed = 0
		return nil
	}

	current := t.GetRoles()[t.seed]
	t.seed++
	return current
}

func (t *BaseTroop) GetTroopIndex() int {
	return t.index
}

func (t *BaseTroop) SetBattle(battle interfaces.Battle) {
	t.battle = battle
}

func (t *BaseTroop) GetBattle() interfaces.Battle {
	return t.battle
}

var _ interfaces.Troop = &BaseTroop{}
