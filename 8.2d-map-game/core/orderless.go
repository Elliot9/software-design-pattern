package core

import (
	"fmt"
	"math/rand"
)

type Orderless struct {
	BaseState
}

func NewOrderless(role Role) *Orderless {
	state := &Orderless{BaseState: BaseState{Role: role, Name: "混亂"}}
	state.BaseState.State = state
	return state
}

func (s *Orderless) RoundStart() {
	if s.Role.GetStateDuration() <= 0 {
		s.Role.SetState(NewNormal(s.Role))
	} else {
		s.Role.SetStateDuration(s.Role.GetStateDuration() - 1)
	}
}

func (s *Orderless) EnterState() {
	s.Role.SetStateDuration(3)
}

func (s *Orderless) Action() {
	fmt.Println(s.GetRole().GetName(), "在混亂狀態下行動！")

	directions := getDirectionPair(rand.Intn(2) == 0)
	str := "請選擇移動方向: "
	for i, direction := range directions {
		str += fmt.Sprintf("%d. %s, ", i+1, string(direction))
	}

	fmt.Println(str)

	// 玩家隨機移動
	if _, isCharacter := s.GetRole().(*Character); isCharacter {
		for {
			var input int
			fmt.Scanln(&input)
			direction := directions[input-1]
			if s.GetRole().Move(direction) {
				break
			}
		}
	} else {
		direction := directions[rand.Intn(2)]
		for {
			if s.GetRole().Move(direction) {
				break
			}
		}
	}
}

func (s *Orderless) Attack() bool {
	// 混亂狀態無法攻擊
	fmt.Println(s.Role.GetName(), "在混亂狀態下無法攻擊！")
	return false
}
