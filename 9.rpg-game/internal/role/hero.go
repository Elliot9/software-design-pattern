package role

import (
	"fmt"
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
)

type Hero struct {
	BaseRole
}

func (h *Hero) MakeDecisionImpl() interfaces.Action {
	actions := h.CLI.ReadNumber()
	return h.Actions[actions[0]]
}

func (h *Hero) SelectTargets(targets []interfaces.Role, count int) []interfaces.Role {
	str := fmt.Sprintf("選擇 %d 位目標: ", count)
	for i, target := range targets {
		str += fmt.Sprintf("(%d) [%d]%s ", i, target.GetTroop().GetTroopIndex(), target.GetName())
	}
	h.GetCLI().Println(str)
	numbers := h.GetCLI().ReadNumber()

	if len(numbers) != count {
		return h.SelectTargets(targets, count)
	}

	selectedTargets := []interfaces.Role{}
	for _, number := range numbers {
		selectedTargets = append(selectedTargets, targets[number])
	}
	return selectedTargets
}

func NewHero(name string, HP, MP, STR int, CLI cli.CLI, actions []interfaces.Action) *Hero {
	state := state.NewNormal()
	hero := &Hero{
		BaseRole: BaseRole{
			HP:      HP,
			MP:      MP,
			STR:     STR,
			Name:    name,
			CLI:     CLI,
			Actions: append([]interfaces.Action{}, actions...),
			State:   state,
			Cursed:  make(map[interfaces.Role]bool),
		},
	}
	hero.BaseRole.Role = hero
	hero.State.SetRole(hero)
	return hero
}

var _ interfaces.Role = &Hero{}
