package collisions

import "github/elliot9/world/core"

type HeroHeroCollision struct {
	core.BaseCollision
}

func NewHeroHeroCollision() *HeroHeroCollision {
	h := &HeroHeroCollision{
		BaseCollision: core.BaseCollision{
			SpriteAType: core.HeroType,
			SpriteBType: core.HeroType,
		},
	}
	return h
}

func (h *HeroHeroCollision) Handle(spriteA core.Sprite, spriteB core.Sprite) core.CollisionResult {
	return core.CollisionResult{
		SpriteA: spriteA,
		SpriteB: spriteB,
		Success: false,
	}
}

var _ core.Collision = &HeroHeroCollision{}
