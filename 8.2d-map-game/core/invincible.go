package core

import "fmt"

type Invincible struct {
	BaseState
}

func (s *Invincible) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Invincible) EnterState() {
	s.Role.SetStateDuration(2)
}

func (s *Invincible) Damage(damage int) {
	fmt.Printf("%s 無敵狀態，受到傷害無效\n", s.Role.GetName())
}

func NewInvincible(role Role) *Invincible {
	state := &Invincible{BaseState: BaseState{Role: role, Name: "無敵"}}
	state.BaseState.State = state
	return state
}
