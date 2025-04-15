package core

type Healing struct {
	BaseState
}

func NewHealing(role Role) *Healing {
	state := &Healing{BaseState: BaseState{Role: role, Name: "恢復"}}
	state.BaseState.State = state
	return state
}

func (s *Healing) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetHp(s.Role.GetHp() + 30)
		if s.Role.GetHp() == s.Role.GetMaxHp() {
			s.Role.SetState(NewNormal(s.Role))
		} else {
			s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
		}
	}
}

func (s *Healing) EnterState() {
	s.Role.SetStateDuration(5)
}
