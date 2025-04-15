package core

type Erupting struct {
	BaseState
}

func NewErupting(role Role) *Erupting {
	state := &Erupting{BaseState: BaseState{Role: role, Name: "爆發"}}
	state.BaseState.State = state
	return state
}

func (s *Erupting) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Erupting) EnterState() {
	s.Role.SetStateDuration(3)
}

func (s *Erupting) Attack() bool {
	return s.Role.GlobalAttack()
}
