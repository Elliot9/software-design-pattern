package role

import (
	"fmt"
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/interfaces"
	"strings"
)

type BaseRole struct {
	HP      int
	MP      int
	STR     int
	State   interfaces.State
	Name    string
	Actions []interfaces.Action
	Role    interfaces.Role
	Troop   interfaces.Troop
	CLI     cli.CLI
	Cursed  map[interfaces.Role]bool
}

func (r *BaseRole) Attack(damage int, target interfaces.Role) {
	r.Role.GetState().Attack(damage, target)
}

func (r *BaseRole) Heal(heal int) {
	r.Role.SetHP(r.Role.GetHP() + heal)
}

func (r *BaseRole) TakeTurn() {
	str := fmt.Sprintf("輪到 [%d]%s (HP: %d, MP: %d, STR: %d, State: %s)。", r.GetTroop().GetTroopIndex(), r.GetName(), r.GetHP(), r.GetMP(), r.GetSTR(), r.GetState().GetName())
	r.GetCLI().Println(str)

	r.Role.GetState().OnTurnStart()

	if !r.IsAlive() {
		return
	}

	r.Role.GetState().TakeTurn()
	r.Role.GetState().OnTurnEnd()
}

func (r *BaseRole) MakeDecisionImpl() interfaces.Action {
	panic("MakeDecisionImpl not implemented")
}

func (r *BaseRole) MakeDecision() interfaces.Action {
	for {
		str := "選擇行動："
		actions := []string{}
		for i, action := range r.Actions {
			actions = append(actions, fmt.Sprintf("(%d) %s", i, action.GetName()))
		}
		str += strings.Join(actions, " ")

		r.CLI.Println(str)
		action := r.Role.MakeDecisionImpl()

		if action.GetMPCost() > r.GetMP() {
			r.CLI.Println("你缺乏 MP，不能進行此行動。")
			continue
		}
		return action
	}
}

func (r *BaseRole) GetTargets(targetType interfaces.TargetType, count int) []interfaces.Role {
	targets := []interfaces.Role{}

	switch targetType {
	case interfaces.TargetAllWithoutSelf:
		targets = append(r.GetTroop().GetRoles(), r.GetTroop().GetBattle().GetEnemyTroop(r.GetTroop()).GetRoles()...)
		for i, target := range targets {
			if target.GetTroop().GetTroopIndex() == r.GetTroop().GetTroopIndex() {
				targets = append(targets[:i], targets[i+1:]...)
				break
			}
		}
	case interfaces.TargetEnemy:
		targets = r.GetTroop().GetBattle().GetEnemyTroop(r.GetTroop()).GetRoles()
		if count < len(targets) {
			targets = r.SelectTargets(targets, count)
		}
	case interfaces.TargetEnemyAll:
		targets = r.GetTroop().GetBattle().GetEnemyTroop(r.GetTroop()).GetRoles()
	case interfaces.TargetAllyWithoutSelf:
		targets = r.GetTroop().GetRoles()
		for i, target := range targets {
			if target == r.Role {
				targets = append(targets[:i], targets[i+1:]...)
				break
			}
		}
		if count < len(targets) {
			targets = r.SelectTargets(targets, count)
		}
	case interfaces.TargetSelf:
		targets = []interfaces.Role{r.Role}
	case interfaces.TargetNil:
		targets = []interfaces.Role{}
	}

	return targets
}

func (r *BaseRole) SelectTargets(targets []interfaces.Role, count int) []interfaces.Role {
	return r.Role.SelectTargets(targets, count)
}

func (r *BaseRole) GetState() interfaces.State {
	return r.State
}

func (r *BaseRole) SetState(state interfaces.State) {
	r.State.ExitState()
	r.State = state
	state.SetRole(r.Role)
	r.State.EnterState()
}

func (r *BaseRole) GetHP() int {
	return r.HP
}

func (r *BaseRole) GetMP() int {
	return r.MP
}

func (r *BaseRole) GetSTR() int {
	return r.STR
}

func (r *BaseRole) GetName() string {
	return r.Name
}

func (r *BaseRole) SetHP(hp int) {
	r.HP = hp
	if !r.IsAlive() {
		r.GetCLI().Println(fmt.Sprintf("[%d]%s 死亡。", r.GetTroop().GetTroopIndex(), r.GetName()))
		for role := range r.Cursed {
			role.Heal(r.GetMP())
		}
		r.Troop.RemoveRole(r.Role)
	}
}

func (r *BaseRole) SetMP(mp int) {
	r.MP = mp
}

func (r *BaseRole) IsAlive() bool {
	return r.HP > 0
}

func (r *BaseRole) SetTroop(troop interfaces.Troop) {
	r.Troop = troop
}

func (r *BaseRole) GetTroop() interfaces.Troop {
	return r.Troop
}

func (r *BaseRole) GetCLI() cli.CLI {
	return r.CLI
}

func (r *BaseRole) AddCursed(from interfaces.Role) {
	r.Cursed[from] = true
}

var _ interfaces.Role = &BaseRole{}
