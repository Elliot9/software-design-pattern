package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
	"strings"
)

type CheerUp struct {
	BaseAction
}

func (c *CheerUp) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s", role.GetTroop().GetTroopIndex(), role.GetName())

	if len(targets) > 0 {
		str += " 對 "
	}

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + c.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		state := state.NewCheerUp()
		target.SetState(state)
	}
}

func NewCheerUp() *CheerUp {
	return &CheerUp{
		BaseAction: BaseAction{
			Name:        "鼓舞",
			MPCost:      100,
			TargetCount: 3,
			TargetType:  interfaces.TargetAllyWithoutSelf,
		},
	}
}

var _ interfaces.Action = &CheerUp{}
