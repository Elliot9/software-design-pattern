package sprites

import "github/elliot9/world/core"

const (
	HeroHP = 30
)

type Hero struct {
	core.BaseSprite
	HP int
}

func NewHero() *Hero {
	hero := &Hero{
		BaseSprite: core.BaseSprite{
			Type: core.HeroType,
		},
		HP: HeroHP,
	}
	hero.Sprite = hero
	return hero
}

func (h *Hero) GetDamage(damage int) {
	h.setHP(h.HP - damage)
}

func (h *Hero) GetHeal(heal int) {
	h.setHP(h.HP + heal)
}

func (h *Hero) GetHP() int {
	return h.HP
}

func (h *Hero) IsDead() bool {
	return h.GetHP() <= 0
}

func (h *Hero) setHP(hp int) {
	h.HP = hp
}

var _ core.Sprite = &Hero{}
