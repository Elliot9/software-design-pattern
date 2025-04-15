package core

import (
	"math/rand/v2"
)

type Character struct {
	BaseRole
}

func NewCharacter() *Character {
	c := &Character{
		BaseRole: BaseRole{
			BaseMapObject: BaseMapObject{
				symbol: "",
			},
			maxHp:         300,
			hp:            300,
			state:         nil,
			stateDuration: 0,
			direction:     []Direction{Up, Right, Down, Left}[rand.IntN(4)],
		},
	}
	c.SetState(NewOrderless(c))
	c.MapObject = c
	c.Role = c
	return c
}

func (c *Character) GetSymbol() string {
	return string(c.direction)
}
