package collisions

import (
	"github/elliot9/world/core"
)

type WaterFireCollision struct {
	core.BaseCollision
}

func NewWaterFireCollision() *WaterFireCollision {
	w := &WaterFireCollision{
		BaseCollision: core.BaseCollision{
			SpriteAType: core.WaterType,
			SpriteBType: core.FireType,
		},
	}
	return w
}

func (w *WaterFireCollision) Handle(spriteA core.Sprite, spriteB core.Sprite) core.CollisionResult {
	return core.CollisionResult{
		SpriteA: nil,
		SpriteB: nil,
		Success: true,
	}
}

var _ core.Collision = &WaterFireCollision{}
