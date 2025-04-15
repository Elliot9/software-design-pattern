package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
)

type BasicAttack struct {
	BaseAction
}

func (b *BasicAttack) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 攻擊 [%d]%s。", role.GetTroop().GetTroopIndex(), role.GetName(), targets[0].GetTroop().GetTroopIndex(), targets[0].GetName())
	role.GetCLI().Println(str)

	for _, target := range targets {
		role.Attack(role.GetSTR(), target)
	}
}

func NewBasicAttack() *BasicAttack {
	return &BasicAttack{
		BaseAction: BaseAction{
			Name:        "普通攻擊",
			MPCost:      0,
			TargetCount: 1,
			TargetType:  interfaces.TargetEnemy,
		},
	}
}

var _ interfaces.Action = &BasicAttack{}
