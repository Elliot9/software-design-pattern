package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"strings"
)

type Fireball struct {
	BaseAction
}

func (w *Fireball) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 對 ", role.GetTroop().GetTroopIndex(), role.GetName())

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + w.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		role.Attack(50, target)
	}
}

func NewFireball() *Fireball {
	return &Fireball{
		BaseAction: BaseAction{
			Name:        "火球",
			MPCost:      50,
			TargetCount: 0,
			TargetType:  interfaces.TargetEnemyAll,
		},
	}
}

var _ interfaces.Action = &Fireball{}
