package core

type Teleport struct {
	BaseState
}

func NewTeleport(role Role) *Teleport {
	state := &Teleport{BaseState: BaseState{Role: role, Name: "瞬身"}}
	state.BaseState.State = state
	return state
}

func (s *Teleport) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Teleport) EnterState() {
	s.Role.SetStateDuration(1)
}

func (s *Teleport) ExitState() {
	s.Role.GetMap().RandomMoveToEmptyGrid(s.Role.(MapObject))
}
