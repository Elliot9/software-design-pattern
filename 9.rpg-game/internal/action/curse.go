package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"strings"
)

type Curse struct {
	BaseAction
}

func (c *Curse) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 對 ", role.GetTroop().GetTroopIndex(), role.GetName())

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + c.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		target.AddCursed(role)
	}
}

func NewCurse() *Curse {
	return &Curse{
		BaseAction: BaseAction{
			Name:        "詛咒",
			MPCost:      100,
			TargetCount: 1,
			TargetType:  interfaces.TargetEnemy,
		},
	}
}

var _ interfaces.Action = &Curse{}
