package main

import (
	"github/elliot9/world/core"
	"github/elliot9/world/entities/collisions"
	"github/elliot9/world/entities/sprites"
	"math/rand"
)

const (
	InitialSpriteNumber = 10
	WorldLength         = 30
)

func main() {
	sprites := initSprites()

	world := core.NewWorld(WorldLength)
	world.SetInitialSprites(sprites[:])

	world.Start()
}

func initSprites() [InitialSpriteNumber]core.Sprite {
	sprites := [InitialSpriteNumber]core.Sprite{}
	for i := 0; i < InitialSpriteNumber; i++ {
		sprites[i] = randomSprite()
	}
	return sprites
}

func createHero() core.Sprite {
	hero := sprites.NewHero()
	hero.SetCollision([]core.Collision{
		collisions.NewHeroHeroCollision(),
		collisions.NewHeroFireCollision(),
		collisions.NewHeroWaterCollision(),
	})
	return hero
}

func createFire() core.Sprite {
	fire := sprites.NewFire()
	fire.SetCollision([]core.Collision{
		collisions.NewFireFireCollision(),
		collisions.NewHeroFireCollision(),
		collisions.NewWaterFireCollision(),
	})
	return fire
}

func createWater() core.Sprite {
	water := sprites.NewWater()
	water.SetCollision([]core.Collision{
		collisions.NewWaterWaterCollision(),
		collisions.NewHeroWaterCollision(),
		collisions.NewWaterFireCollision(),
	})
	return water
}

func randomSprite() core.Sprite {
	spriteType := rand.Intn(3)

	switch spriteType {
	case 0:
		return createHero()
	case 1:
		return createFire()
	case 2:
		return createWater()
	default:
		return createHero()
	}
}
