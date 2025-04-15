package core

import (
	"fmt"
	"math/rand/v2"
)

func NewWorld(length int) *World {
	world := &World{
		positions: make([]Sprite, length),
	}
	return world
}

type World struct {
	positions []Sprite
}

func (w *World) SetInitialSprites(sprites []Sprite) {
	copy(w.positions, sprites)

	// 隨機賦予初始座標
	rand.Shuffle(len(w.positions), func(i, j int) {
		w.positions[i], w.positions[j] = w.positions[j], w.positions[i]
	})
}

func (w *World) Start() {
	for {
		w.printfWorld()
		fmt.Println("請使用者輸入兩個數字（以空白隔開）")
		var x1, x2 int
		fmt.Scanln(&x1, &x2)

		// 取得兩個不同座標的生命
		spriteA := w.positions[x1]
		spriteB := w.positions[x2]

		if spriteA == nil || spriteB == nil {
			// 如果其中一個是 nil，則不會觸發碰撞
			w.positions[x1], w.positions[x2] = w.positions[x2], w.positions[x1]
			continue
		}

		// 發生碰撞
		collisionResult := spriteA.HandleCollision(spriteB)

		// 更新世界
		w.positions[x1], w.positions[x2] = collisionResult.SpriteA, collisionResult.SpriteB

		// 如果移動成功, 交換位置
		if collisionResult.Success {
			w.positions[x1], w.positions[x2] = w.positions[x2], w.positions[x1]
		}
	}
}

func (w *World) printfWorld() {
	for i, sprite := range w.positions {
		if sprite == nil {
			fmt.Printf("%d:_\n", i)
		} else {
			fmt.Printf("%d:%s\n", i, sprite.GetType())
		}
	}
}
