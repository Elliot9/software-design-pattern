package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"strings"
)

type Waterball struct {
	BaseAction
}

func (w *Waterball) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 對 ", role.GetTroop().GetTroopIndex(), role.GetName())

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + w.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		role.Attack(120, target)
	}
}

func NewWaterball() *Waterball {
	return &Waterball{
		BaseAction: BaseAction{
			Name:        "水球",
			MPCost:      50,
			TargetCount: 1,
			TargetType:  interfaces.TargetEnemy,
		},
	}
}

var _ interfaces.Action = &Waterball{}
