package core

type Collision interface {
	Match(spriteA Sprite, spriteB Sprite) bool
	Handle(spriteA Sprite, spriteB Sprite) CollisionResult
}

type CollisionResult struct {
	SpriteA Sprite
	SpriteB Sprite
	Success bool
}

type BaseCollision struct {
	SpriteAType SpriteType
	SpriteBType SpriteType
}

func (c *BaseCollision) GetSpriteAType() SpriteType {
	return c.SpriteAType
}

func (c *BaseCollision) GetSpriteBType() SpriteType {
	return c.SpriteBType
}

func (c *BaseCollision) Match(spriteA Sprite, spriteB Sprite) bool {
	return (c.GetSpriteAType() == spriteA.GetType() && c.GetSpriteBType() == spriteB.GetType()) ||
		(c.GetSpriteAType() == spriteB.GetType() && c.GetSpriteBType() == spriteA.GetType())
}

func (c *BaseCollision) Handle(spriteA Sprite, spriteB Sprite) CollisionResult {
	panic("Handle not implemented")
}

var _ Collision = &BaseCollision{}
