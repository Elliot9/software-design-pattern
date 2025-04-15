package core

type Accelerated struct {
	BaseState
}

func NewAccelerated(role Role) *Accelerated {
	state := &Accelerated{BaseState: BaseState{Role: role, Name: "加速"}}
	state.BaseState.State = state
	return state
}

func (s *Accelerated) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Accelerated) EnterState() {
	s.Role.SetStateDuration(3)
}

func (s *Accelerated) Damage(damage int) {
	s.Role.SetHp(s.Role.GetHp() - damage)
	s.Role.SetState(NewNormal(s.Role))
}

func (s *Accelerated) Action() {
	// 執行兩次行動
	for i := 0; i < 2; i++ {
		s.BaseState.Action()
	}
}
