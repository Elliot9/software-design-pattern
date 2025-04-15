package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"strings"
)

type OnePunch struct {
	BaseAction
	handler interfaces.OnePunchHandler
}

func (o *OnePunch) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 對 ", role.GetTroop().GetTroopIndex(), role.GetName())

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + o.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		o.handler.Handle(role, target)
	}
}

func NewOnePunch(handler interfaces.OnePunchHandler) *OnePunch {
	return &OnePunch{
		BaseAction: BaseAction{
			Name:        "一拳攻擊",
			MPCost:      180,
			TargetCount: 1,
			TargetType:  interfaces.TargetEnemy,
		},
		handler: handler,
	}
}

var _ interfaces.Action = &OnePunch{}
