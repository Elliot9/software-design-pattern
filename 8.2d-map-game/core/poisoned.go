package core

type Poisoned struct {
	BaseState
}

func NewPoisoned(role Role) *Poisoned {
	state := &Poisoned{BaseState: BaseState{Role: role, Name: "中毒"}}
	state.BaseState.State = state
	return state
}

func (s *Poisoned) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetHp(s.Role.GetHp() - 15)
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Poisoned) EnterState() {
	s.Role.SetStateDuration(3)
}
