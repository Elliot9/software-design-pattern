package role

import (
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/state"
)

type AI struct {
	BaseRole
	AIStrategy interfaces.AIStrategy
}

func (a *AI) MakeDecisionImpl() interfaces.Action {
	return a.AIStrategy.MakeDecisionImpl()
}

func (a *AI) SelectTargets(targets []interfaces.Role, count int) []interfaces.Role {
	return a.AIStrategy.SelectTargets(targets, count)
}

func NewSeedAI(name string, HP, MP, STR int, CLI cli.CLI, actions []interfaces.Action) *AI {
	state := state.NewNormal()
	ai := &AI{
		BaseRole: BaseRole{
			Actions: append([]interfaces.Action{}, actions...),
			HP:      HP,
			MP:      MP,
			STR:     STR,
			Name:    name,
			CLI:     CLI,
			State:   state,
			Cursed:  make(map[interfaces.Role]bool),
		},
		AIStrategy: &SeedAIStrategy{
			seed: 0,
		},
	}
	ai.BaseRole.Role = ai
	ai.State.SetRole(ai)
	ai.AIStrategy.(*SeedAIStrategy).AI = ai
	return ai
}

type SeedAIStrategy struct {
	seed int
	AI   *AI
}

func (s *SeedAIStrategy) MakeDecisionImpl() interfaces.Action {
	action := s.AI.BaseRole.Actions[s.seed%len(s.AI.BaseRole.Actions)]
	s.IncreaseSeed()
	return action
}

func (s *SeedAIStrategy) SelectTargets(targets []interfaces.Role, count int) []interfaces.Role {
	selectedTargets := []interfaces.Role{}
	seed := s.seed

	for i := 0; i < count; i++ {
		selectedTargets = append(selectedTargets, targets[seed%len(targets)])
		seed++
	}
	s.IncreaseSeed()
	return selectedTargets
}

func (s *SeedAIStrategy) IncreaseSeed() {
	s.seed++
}

var _ interfaces.Role = &AI{}
var _ interfaces.AIStrategy = &SeedAIStrategy{}
