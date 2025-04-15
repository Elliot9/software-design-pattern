package collisions

import (
	"fmt"
	"github/elliot9/world/core"
	"github/elliot9/world/entities/sprites"
)

type HeroWaterCollision struct {
	core.BaseCollision
}

func NewHeroWaterCollision() *HeroWaterCollision {
	h := &HeroWaterCollision{
		BaseCollision: core.BaseCollision{
			SpriteAType: core.HeroType,
			SpriteBType: core.WaterType,
		},
	}
	return h
}

func (h *HeroWaterCollision) Handle(spriteA core.Sprite, spriteB core.Sprite) core.CollisionResult {
	var hero *sprites.Hero
	isHeroInitiator := true

	fmt.Printf("spriteA: %+v\n", spriteA)
	fmt.Printf("spriteB: %+v\n", spriteB)

	if spriteA.GetType() == core.HeroType {
		hero, _ = spriteA.(*sprites.Hero)
	} else {
		hero, _ = spriteB.(*sprites.Hero)
		isHeroInitiator = false
	}

	// Hero 生命值增加 10 滴
	hero.GetHeal(10)

	fmt.Printf("hero: %+v\n", hero)

	var result core.CollisionResult
	// 如果 c1 為 Hero，則 c1 移動成功
	// Water 從世界中被移除
	if isHeroInitiator {
		result = core.CollisionResult{
			SpriteA: hero,
			SpriteB: nil,
			Success: true,
		}
	} else {
		result = core.CollisionResult{
			SpriteA: nil,
			SpriteB: hero,
			Success: false,
		}
	}

	return result
}

var _ core.Collision = &HeroFireCollision{}
