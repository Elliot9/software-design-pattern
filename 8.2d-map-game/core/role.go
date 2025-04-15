package core

import (
	"fmt"
)

type Role interface {
	Action()
	Attack() bool
	Damage(damage int)
	SetState(state State)
	RoundStart()
	GetHp() int
	GetMaxHp() int
	SetHp(hp int)
	SetStateDuration(duration int)
	GetStateDuration() int
	GetName() string
	GetMap() *Map
	Move(direction Direction) bool
	LineAttack() bool
	SurroundAttack() bool
	GlobalAttack() bool
}

type BaseRole struct {
	BaseMapObject
	maxHp         int
	hp            int
	state         State
	stateDuration int
	direction     Direction
	Role
}

func (r *BaseRole) SetState(state State) {
	if r.state != nil {
		r.state.ExitState()
	}

	r.state = state

	if r.state != nil {
		r.state.EnterState()
	}
}

func (r *BaseRole) RoundStart() {
	r.state.RoundStart()
}

func (r *BaseRole) Action() {
	r.state.Action()
}

func (r *BaseRole) Attack() bool {
	return r.state.Attack()
}

func (r *BaseRole) Move(direction Direction) bool {
	originX, originY := r.GetPosition()
	nextX, nextY := originX, originY

	switch direction {
	case Up:
		nextY--
	case Down:
		nextY++
	case Left:
		nextX--
	case Right:
		nextX++
	}

	if !r.GetMap().IsValidPosition(nextX, nextY) {
		return false
	}

	nextGrid := r.GetMap().GetGrid(nextX, nextY)
	if nextGrid == nil {
		fmt.Printf("%s 往 %s 方向移動！\n", r.GetName(), string(direction))
		r.GetMap().MoveObject(originX, originY, nextX, nextY)
	} else {
		r.touch(nextGrid)
	}

	r.SetDirection(direction)
	return true
}

func (r *BaseRole) GetHp() int {
	return r.hp
}

func (r *BaseRole) GetMaxHp() int {
	return r.maxHp
}

func (r *BaseRole) SetHp(hp int) {
	if hp >= r.maxHp {
		r.hp = r.maxHp
	} else if hp <= 0 {
		r.hp = 0

		if target, ok := r.Role.(*Monster); ok {
			for i, monster := range r.GetMap().GetGame().monsters {
				if monster == target {
					fmt.Printf("怪物死亡，移除怪物\n")
					r.GetMap().RemoveObject(target.GetPosition())
					r.GetMap().GetGame().monsters = append(r.GetMap().GetGame().monsters[:i], r.GetMap().GetGame().monsters[i+1:]...)
					break
				}
			}
		}

	} else {
		r.hp = hp
	}
}

func (r *BaseRole) GetState() State {
	return r.state
}

func (r *BaseRole) SetStateDuration(duration int) {
	r.stateDuration = duration
}

func (r *BaseRole) GetStateDuration() int {
	return r.stateDuration
}

func (r *BaseRole) GetDirection() Direction {
	return r.direction
}

func (r *BaseRole) SetDirection(direction Direction) {
	r.direction = direction
}

func (r *BaseRole) Damage(damage int) {
	r.state.Damage(damage)
}

func (r *BaseRole) touch(target MapObject) {
	name := r.GetName()

	if _, ok := target.(*Obstacle); ok {
		// 如果移動到障礙物，則不移動
		fmt.Printf("%s 移動到障礙物，無法通過！\n", name)
	} else if treasure, ok := target.(*Treasure); ok {
		// 如果移動到寶藏，則獲得效果
		state := treasure.State
		state.SetRole(r.Role)
		r.Role.SetState(state)
		r.GetMap().RemoveObject(treasure.GetPosition())
		fmt.Printf("%s 碰觸到寶藏，獲得效果：%s\n", name, state.GetName())
	}
}

func (r *BaseRole) GetMap() *Map {
	return r.BaseMapObject.GetMap()
}

func (r *BaseRole) GetName() string {
	if _, ok := r.Role.(*Monster); ok {
		return "怪物"
	} else {
		return "玩家"
	}
}

func getDirectionPair(isVertical bool) []Direction {
	if isVertical {
		return []Direction{Up, Down}
	}
	return []Direction{Left, Right}
}

func (r *BaseRole) LineAttack() bool {
	x, y := r.GetPosition()
	hasHit := false

	for {
		switch r.direction {
		case Up:
			y--
		case Down:
			y++
		case Left:
			x--
		case Right:
			x++
		}

		// 檢查是否超出邊界
		if !r.GetMap().IsValidPosition(x, y) {
			break
		}

		// 檢查目標格子
		target := r.GetMap().GetGrid(x, y)
		if target == nil {
			continue
		}

		// 如果是障礙物，停止攻擊
		if _, isObstacle := target.(*Obstacle); isObstacle {
			break
		}

		// 如果是怪物，造成傷害
		if monster, ok := target.(*Monster); ok {
			monster.Damage(50)
			hasHit = true
		}
	}
	return hasHit
}

func (r *BaseRole) SurroundAttack() bool {
	x, y := r.GetPosition()
	name := r.GetName()
	hasHit := false

	// 檢查上下左右四個方向
	directions := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}

	for _, d := range directions {
		newX, newY := x+d.dx, y+d.dy
		if !r.GetMap().IsValidPosition(newX, newY) {
			continue
		}

		if target := r.GetMap().GetGrid(newX, newY); target != nil {
			if character, ok := target.(*Character); ok {
				fmt.Printf("%s 對 %s 發動攻擊\n", name, character.GetName())
				character.Damage(50)
				hasHit = true
			}
		}
	}
	return hasHit
}

func (r *BaseRole) GlobalAttack() bool {
	hasHit := false

	if character, ok := r.Role.(*Character); ok {
		fmt.Printf("玩家 對 怪物 發動全圖攻擊\n")
		monsters := character.GetMap().GetGame().monsters
		for i := len(monsters) - 1; i >= 0; i-- {
			monsters[i].Damage(50)
			hasHit = true
		}
	} else {
		fmt.Printf("怪物 對 玩家 發動全圖攻擊\n")
		r.GetMap().GetGame().player.Damage(50)
		hasHit = true
	}

	return hasHit
}

var _ Role = &BaseRole{}
var _ MapObject = &BaseRole{}
