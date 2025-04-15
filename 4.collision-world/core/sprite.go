package core

type SpriteType string

const (
	HeroType  SpriteType = "H"
	WaterType SpriteType = "W"
	FireType  SpriteType = "F"
)

type Sprite interface {
	GetType() SpriteType
	SetCollision(collisions []Collision)
	HandleCollision(sprite Sprite) CollisionResult
}

type BaseSprite struct {
	Type       SpriteType
	collisions []Collision
	Sprite
}

func (s *BaseSprite) GetType() SpriteType {
	return s.Type
}

func (s *BaseSprite) SetCollision(collisions []Collision) {
	s.collisions = collisions
}

func (s *BaseSprite) HandleCollision(sprite Sprite) CollisionResult {
	for _, collision := range s.collisions {
		if collision.Match(s.Sprite, sprite) {
			return collision.Handle(s.Sprite, sprite)
		}
	}
	return CollisionResult{
		SpriteA: s.Sprite,
		SpriteB: sprite,
		Success: false,
	}
}

var _ Sprite = &BaseSprite{}
