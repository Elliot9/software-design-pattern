package collisions

import (
	"github/elliot9/world/core"
	"github/elliot9/world/entities/sprites"
)

type HeroFireCollision struct {
	core.BaseCollision
}

func NewHeroFireCollision() *HeroFireCollision {
	h := &HeroFireCollision{
		BaseCollision: core.BaseCollision{
			SpriteAType: core.HeroType,
			SpriteBType: core.FireType,
		},
	}
	return h
}

func (h *HeroFireCollision) Handle(spriteA core.Sprite, spriteB core.Sprite) core.CollisionResult {
	var hero *sprites.Hero
	isHeroInitiator := true

	if spriteA.GetType() == core.HeroType {
		hero, _ = spriteA.(*sprites.Hero)
	} else {
		hero, _ = spriteB.(*sprites.Hero)
		isHeroInitiator = false
	}

	// Hero 生命值減少 10 滴
	hero.GetDamage(10)

	// 如果 HP ≤ 0 時，Hero 死亡，並且會從世界中被移除
	if hero.IsDead() {
		if isHeroInitiator {
			return core.CollisionResult{
				SpriteA: nil,
				SpriteB: nil,
				Success: true,
			}
		}
		return core.CollisionResult{
			SpriteA: nil,
			SpriteB: nil,
			Success: false,
		}
	}

	var result core.CollisionResult
	// 如果 c1 為 Hero，則 c1 移動成功
	// Fire 從世界中被移除
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
