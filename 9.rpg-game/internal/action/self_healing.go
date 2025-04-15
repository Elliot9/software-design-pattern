package action

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
)

type SelfHealing struct {
	BaseAction
}

func (w *SelfHealing) Excute(role interfaces.Role, targets []interfaces.Role) {
	str := fmt.Sprintf("[%d]%s 使用了 %s。", role.GetTroop().GetTroopIndex(), role.GetName(), w.GetName())
	role.GetCLI().Println(str)
	role.Heal(150)
}

func NewSelfHealing() *SelfHealing {
	return &SelfHealing{
		BaseAction: BaseAction{
			Name:        "自我治療",
			MPCost:      50,
			TargetCount: 0,
			TargetType:  interfaces.TargetSelf,
		},
	}
}

var _ interfaces.Action = &SelfHealing{}
