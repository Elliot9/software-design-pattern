package core

import (
	"fmt"
	"math/rand/v2"
)

type Game struct {
	mapper   *Map
	player   *Character
	monsters []*Monster
}

func NewGame(width, height int) *Game {
	g := &Game{
		mapper: NewMap(width, height),
	}
	g.initialize()
	g.mapper.SetGame(g)
	return g
}

func (g *Game) Start() {
	for !g.isGameOver() {
		// 寶物跟怪物，在各個回合都有機率在隨機地點生成
		g.generateNewTreasure()
		if rand.IntN(10) < 2 {
			g.generateNewMonster(1)
		}

		// 回合開始前更新角色狀態
		g.player.BaseRole.RoundStart()
		for _, monster := range g.monsters {
			monster.BaseRole.RoundStart()
		}

		// 每一回合開始時先印出二維地圖 以及印出主角的生命值、狀態
		g.printMap()
		g.printCharacter()

		// 請玩家選擇欲讓主角執行的動作
		if !g.isGameOver() {
			g.player.Action()
		}

		// 怪物行動
		for i := 0; i < len(g.monsters); i++ {
			if g.isGameOver() {
				break
			}
			monster := g.monsters[i]
			monster.Action()
		}
	}

	fmt.Println("遊戲結束")
}

func (g *Game) initialize() {
	// 初始化玩家
	g.generateNewPlayer()

	// 初始化怪物
	g.generateNewMonster(10)

	// 初始化寶藏
	g.generateNewTreasure()

	// 初始化障礙物
	g.generateNewObstacle(rand.IntN(5))
}

func (g *Game) isGameOver() bool {
	return g.player.BaseRole.GetHp() <= 0 || len(g.monsters) == 0
}

func (g *Game) generateNewPlayer() {
	player := NewCharacter()
	player.SetMap(g.mapper)
	emptyX, emptyY := g.mapper.GetRandomEmptyPosition()
	g.mapper.AddObject(emptyX, emptyY, player)
	g.player = player
}

func (g *Game) generateNewMonster(count int) {
	for i := 0; i < count; i++ {
		monster := NewMonster()
		monster.SetMap(g.mapper)
		emptyX, emptyY := g.mapper.GetRandomEmptyPosition()
		g.mapper.AddObject(emptyX, emptyY, monster)
		g.monsters = append(g.monsters, monster)
	}
}
func (g *Game) generateNewTreasure() {
	var treasureConfigs = map[string]struct {
		Probability float64
		State       func() State
	}{
		"Super Star":          {Probability: 0.1, State: func() State { return NewInvincible(nil) }},
		"Poison":              {Probability: 0.25, State: func() State { return NewPoisoned(nil) }},
		"Accelerating Potion": {Probability: 0.2, State: func() State { return NewAccelerated(nil) }},
		"Healing Potion":      {Probability: 0.15, State: func() State { return NewHealing(nil) }},
		"Devil Fruit":         {Probability: 0.1, State: func() State { return NewOrderless(nil) }},
		"King's Rock":         {Probability: 0.1, State: func() State { return NewStockpile(nil) }},
		"Dokodemo Door":       {Probability: 0.1, State: func() State { return NewTeleport(nil) }},
	}

	r := rand.Float64()
	for name, treasure := range treasureConfigs {
		if r < treasure.Probability {
			treasure := NewTreasure(name, treasure.Probability, treasure.State())
			emptyX, emptyY := g.mapper.GetRandomEmptyPosition()
			g.mapper.AddObject(emptyX, emptyY, treasure)
			treasure.SetMap(g.mapper)
		}
	}
}

func (g *Game) generateNewObstacle(count int) {
	for i := 0; i < count; i++ {
		obstacle := NewObstacle()
		emptyX, emptyY := g.mapper.GetRandomEmptyPosition()
		g.mapper.AddObject(emptyX, emptyY, obstacle)
		obstacle.SetMap(g.mapper)
	}
}

func (g *Game) printCharacter() {
	fmt.Printf("[Player] HP: %d, State: %s, Direction: %s\n", g.player.BaseRole.GetHp(), g.player.BaseRole.GetState().GetName(), string(g.player.BaseRole.GetDirection()))
}

func (g *Game) printMap() {
	g.mapper.Print()
}
