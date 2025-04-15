package core

import (
	"fmt"
	"math/rand"
)

type Map struct {
	grid [][]MapObject
	game *Game
}

func NewMap(width, height int) *Map {
	m := &Map{
		grid: make([][]MapObject, height),
	}
	for i := 0; i < height; i++ {
		m.grid[i] = make([]MapObject, width)
	}
	return m
}

func (m *Map) AddObject(x, y int, obj MapObject) {
	m.grid[y][x] = obj
}

func (m *Map) RemoveObject(x, y int) {
	m.grid[y][x] = nil
}

func (m *Map) GetRandomEmptyPosition() (int, int) {
	for {
		x, y := m.getRandomPosition()
		if m.grid[y][x] == nil {
			return x, y
		}
	}
}

func (m *Map) RandomMoveToEmptyGrid(obj MapObject) {
	for {
		x, y := m.GetRandomEmptyPosition()
		if m.IsValidPosition(x, y) {
			fromX, fromY := obj.GetPosition()
			m.MoveObject(fromX, fromY, x, y)
			return
		}
	}
}

func (m *Map) MoveObject(fromX, fromY, toX, toY int) {
	m.grid[toY][toX] = m.grid[fromY][fromX]
	m.grid[fromY][fromX] = nil
}

func (m *Map) IsValidPosition(x, y int) bool {
	return x >= 0 && x < len(m.grid[0]) && y >= 0 && y < len(m.grid)
}

func (m *Map) getRandomPosition() (int, int) {
	return rand.Intn(len(m.grid[0])), rand.Intn(len(m.grid))
}

func (m *Map) GetGrid(x, y int) MapObject {
	return m.grid[y][x]
}

func (m *Map) Print() {
	for _, row := range m.grid {
		for _, obj := range row {
			if obj == nil {
				fmt.Print("_")
			} else {
				fmt.Print(obj.GetSymbol())
			}
		}
		fmt.Println()
	}
}

func (m *Map) SetGame(game *Game) {
	m.game = game
}

func (m *Map) GetGame() *Game {
	return m.game
}
