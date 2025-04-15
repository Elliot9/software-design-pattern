package collisions

import "github/elliot9/world/core"

type WaterWaterCollision struct {
	core.BaseCollision
}

func NewWaterWaterCollision() *WaterWaterCollision {
	w := &WaterWaterCollision{
		BaseCollision: core.BaseCollision{
			SpriteAType: core.WaterType,
			SpriteBType: core.WaterType,
		},
	}
	return w
}

func (w *WaterWaterCollision) Handle(spriteA core.Sprite, spriteB core.Sprite) core.CollisionResult {
	return core.CollisionResult{
		SpriteA: spriteA,
		SpriteB: spriteB,
		Success: false,
	}
}

var _ core.Collision = &WaterWaterCollision{}
