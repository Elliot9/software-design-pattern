package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
	"strings"
)

type SelfExplosion struct {
	BaseAction
}

func (s *SelfExplosion) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 對 ", role.GetTroop().GetTroopIndex(), role.GetName())

	targetArray := []string{}
	for _, target := range targets {
		targetArray = append(targetArray, fmt.Sprintf("[%d]%s", target.GetTroop().GetTroopIndex(), target.GetName()))
	}
	targetsStr := strings.Join(targetArray, ", ")
	str += targetsStr + " 使用了 " + s.GetName() + "。"

	role.GetCLI().Println(str)

	for _, target := range targets {
		role.Attack(150, target)
	}
	role.SetHP(0)
}

func NewSelfExplosion() *SelfExplosion {
	return &SelfExplosion{
		BaseAction: BaseAction{
			Name:        "自爆",
			MPCost:      200,
			TargetCount: 0,
			TargetType:  interfaces.TargetAllWithoutSelf,
		},
	}
}

var _ interfaces.Action = &SelfExplosion{}
