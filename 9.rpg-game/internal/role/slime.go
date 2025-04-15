package role

import (
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/interfaces"
)

type Slime struct {
	*AI
	summoner interfaces.Role
}

func NewSlime(summoner interfaces.Role, HP, MP, STR int, CLI cli.CLI, actions []interfaces.Action) *Slime {
	slime := &Slime{
		AI:       NewSeedAI("Slime", HP, MP, STR, CLI, actions),
		summoner: summoner,
	}

	slime.BaseRole.Role = slime
	slime.State.SetRole(slime)

	return slime
}

func (s *Slime) SetHP(hp int) {
	if hp <= 0 && s.summoner != nil && s.summoner.IsAlive() {
		s.summoner.Heal(30)
	}
	s.BaseRole.SetHP(hp)
}
