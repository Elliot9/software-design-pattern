package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	r "github/elliot9/class9/internal/role"
)

type Summon struct {
	BaseAction
}

func (w *Summon) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 使用了 %s。", role.GetTroop().GetTroopIndex(), role.GetName(), w.GetName())
	role.GetCLI().Println(str)

	slime := r.NewSlime(role, 100, 0, 50, role.GetCLI(), []interfaces.Action{NewBasicAttack()})
	role.GetTroop().AddRole(slime)
}

func NewSummon() *Summon {
	return &Summon{
		BaseAction: BaseAction{
			Name:        "召喚",
			MPCost:      150,
			TargetCount: 0,
			TargetType:  interfaces.TargetNil,
		},
	}
}

var _ interfaces.Action = &Summon{}
