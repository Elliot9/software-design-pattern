package sprites

import (
	"github/elliot9/world/core"
)

type Fire struct {
	core.BaseSprite
}

func NewFire() *Fire {
	fire := &Fire{
		BaseSprite: core.BaseSprite{
			Type: core.FireType,
		},
	}
	fire.Sprite = fire
	return fire
}

var _ core.Sprite = &Fire{}
