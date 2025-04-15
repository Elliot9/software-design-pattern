package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
	"strings"
)

type Petrochemical struct {
	BaseAction
}

func (w *Petrochemical) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 對 ", role.GetTroop().GetTroopIndex(), role.GetName())

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + w.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		state := state.NewPetrochemical()
		target.SetState(state)
	}
}

func NewPetrochemical() *Petrochemical {
	return &Petrochemical{
		BaseAction: BaseAction{
			Name:        "石化",
			MPCost:      100,
			TargetCount: 1,
			TargetType:  interfaces.TargetEnemy,
		},
	}
}

var _ interfaces.Action = &Petrochemical{}
