package state

import (
	"fmt"
	"github/elliot9/class9/internal/interfaces"
)

type BaseState struct {
	Role          interfaces.Role
	Name          string
	StateDuration int
	State         interfaces.State
}

func (s *BaseState) GetName() string {
	return s.Name
}

func (s *BaseState) OnTurnStart() {
	s.SetStateDuration(s.GetStateDuration() - 1)
}

func (s *BaseState) OnTurnEnd() {
	if s.GetStateDuration() <= 0 {
		s.GetRole().SetState(NewNormal())
	}
}

func (s *BaseState) TakeTurn() {
	action := s.GetRole().MakeDecision()
	targets := s.GetRole().GetTargets(action.GetTargetType(), action.GetTargetCount())
	action.Excute(s.GetRole(), targets)
	s.GetRole().SetMP(s.GetRole().GetMP() - action.GetMPCost())
}

func (s *BaseState) Attack(damage int, target interfaces.Role) {
	s.GetRole().GetCLI().Println(fmt.Sprintf("[%d]%s 對 [%d]%s 造成 %d 點傷害。", s.GetRole().GetTroop().GetTroopIndex(), s.GetRole().GetName(), target.GetTroop().GetTroopIndex(), target.GetName(), damage))
	target.SetHP(target.GetHP() - damage)
}

func (s *BaseState) EnterState() {
}

func (s *BaseState) ExitState() {
}

func (s *BaseState) SetStateDuration(duration int) {
	s.StateDuration = duration
}

func (s *BaseState) GetStateDuration() int {
	return s.StateDuration
}

func (s *BaseState) SetRole(role interfaces.Role) {
	s.Role = role
}

func (s *BaseState) GetRole() interfaces.Role {
	return s.Role
}
