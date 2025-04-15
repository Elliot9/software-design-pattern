package collisions

import (
	"github/elliot9/world/core"
)

type FireFireCollision struct {
	core.BaseCollision
}

func NewFireFireCollision() *FireFireCollision {
	f := &FireFireCollision{
		BaseCollision: core.BaseCollision{
			SpriteAType: core.FireType,
			SpriteBType: core.FireType,
		},
	}
	return f
}

func (f *FireFireCollision) Handle(spriteA core.Sprite, spriteB core.Sprite) core.CollisionResult {
	return core.CollisionResult{
		SpriteA: spriteA,
		SpriteB: spriteB,
		Success: false,
	}
}

var _ core.Collision = &FireFireCollision{}
