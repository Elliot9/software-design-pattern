package core

import (
	"math/rand"
)

type Monster struct {
	BaseRole
}

func NewMonster() *Monster {
	m := &Monster{
		BaseRole: BaseRole{
			BaseMapObject: BaseMapObject{
				symbol: "M",
			},
			maxHp:         1,
			hp:            1,
			state:         nil,
			stateDuration: 0,
			direction:     []Direction{Up, Right, Down, Left}[rand.Intn(4)],
		},
	}
	m.SetState(NewNormal(m))
	m.MapObject = m
	m.Role = m
	return m
}
