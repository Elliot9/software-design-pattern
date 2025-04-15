package core

import (
	"fmt"
	"math/rand"
)

type State interface {
	RoundStart()
	Damage(damage int)
	EnterState()
	ExitState()
	GetName() string

	SetRole(role Role)
	GetRole() Role

	Action()
	Attack() bool
}

type BaseState struct {
	Name string
	Role Role
	State
}

func (s *BaseState) RoundStart() {
	s.State.RoundStart()
}

func (s *BaseState) Damage(damage int) {
	s.GetRole().SetHp(s.GetRole().GetHp() - damage)
	if s.GetRole().GetHp() > 0 {
		s.GetRole().SetState(NewInvincible(s.GetRole()))
	}
}

func (s *BaseState) EnterState() {
	s.State.EnterState()
}

func (s *BaseState) ExitState() {}

func (s *BaseState) SetRole(role Role) {
	s.Role = role
}

func (s *BaseState) GetRole() Role {
	return s.Role
}

func (s *BaseState) Action() {
	// 預設行動邏輯
	if _, isCharacter := s.GetRole().(*Character); isCharacter {
		// 玩家行動
		fmt.Println("請選擇欲讓主角執行的動作: (1. 移動, 2. 攻擊)")
		var action string
		fmt.Scanln(&action)

		if action == "1" {
			directions := []Direction{Up, Down, Left, Right}
			str := "請選擇移動方向: "
			for i, direction := range directions {
				str += fmt.Sprintf("%d. %s, ", i+1, string(direction))
			}
			fmt.Println(str)

			for {
				var input int
				fmt.Scanln(&input)
				direction := directions[input-1]

				if !s.GetRole().Move(direction) {
					fmt.Println("玩家移動失敗，請重新輸入移動方向")
				} else {
					break
				}
			}
		} else {
			s.GetRole().Attack()
		}
	} else {
		// 怪物行動
		if !s.GetRole().Attack() {
			for {
				if s.GetRole().Move([]Direction{Up, Down, Left, Right}[rand.Intn(4)]) {
					break
				}
			}
		}
	}
}

func (s *BaseState) Attack() bool {
	if _, isCharacter := s.GetRole().(*Character); isCharacter {
		return s.GetRole().LineAttack()
	} else {
		return s.GetRole().SurroundAttack()
	}
}

func (s *BaseState) GetName() string {
	return s.Name
}

var _ State = &BaseState{}
var _ State = &Normal{}
var _ State = &Invincible{}
var _ State = &Poisoned{}
var _ State = &Orderless{}
var _ State = &Stockpile{}
var _ State = &Erupting{}
var _ State = &Teleport{}
var _ State = &Healing{}
