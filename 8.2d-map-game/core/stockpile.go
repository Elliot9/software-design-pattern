package core

type Stockpile struct {
	BaseState
}

func NewStockpile(role Role) *Stockpile {
	state := &Stockpile{BaseState: BaseState{Role: role, Name: "蓄力"}}
	state.BaseState.State = state
	return state
}

func (s *Stockpile) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewErupting(s.Role))
	} else {
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Stockpile) EnterState() {
	s.Role.SetStateDuration(2)
}

func (s *Stockpile) Damage(damage int) {
	s.Role.SetHp(s.Role.GetHp() - damage)
	s.Role.SetState(NewNormal(s.Role))
}
