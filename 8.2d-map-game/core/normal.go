package core

type Normal struct {
	BaseState
}

func (s *Normal) RoundStart() {}
func (s *Normal) EnterState() { s.Role.SetStateDuration(0) }
func (s *Normal) ExitState()  {}

func NewNormal(role Role) *Normal {
	state := &Normal{BaseState: BaseState{Role: role, Name: "正常"}}
	state.BaseState.State = state
	return state
}
