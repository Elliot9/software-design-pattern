package sprites

import "github/elliot9/world/core"

type Water struct {
	core.BaseSprite
}

func NewWater() *Water {
	water := &Water{
		BaseSprite: core.BaseSprite{
			Type: core.WaterType,
		},
	}
	water.Sprite = water
	return water
}

var _ core.Sprite = &Water{}
