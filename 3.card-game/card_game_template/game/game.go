package game

import (
	"fmt"
	"github/elliot9/card_game_template/deck"
	"github/elliot9/card_game_template/player"
)

const (
	PlayerCount     = 4
	RealPlayerCount = 1
)

type Game interface {
	Start()
	InitDeck()
	InitPlayers()
	GetDefaultHandsCount() int
	InitPlayRound()
	PlayRound()
	IsGameOver() bool
	GetWinner() player.Player
}

type BaseGame struct {
	deck    deck.Deck
	players []player.Player
	Game
}

func (g *BaseGame) Start() {
	// 初始化牌堆 ????
	g.Game.InitDeck()

	// 初始化玩家 ????
	g.Game.InitPlayers()

	// 玩家為自己取名
	for _, player := range g.players {
		player.NameSelf()
	}

	// 牌堆進行洗牌
	g.deck.Shuffle()

	// 玩家依序抽牌，直到玩家擁有 ???? 張牌為止
	for i := 0; i < g.Game.GetDefaultHandsCount(); i++ {
		for _, player := range g.players {
			card := g.deck.Draw()
			player.AddHands(card)
		}
	}

	// 初始化回合階段 ????
	g.Game.InitPlayRound()

	// 遊戲採用 ???? 遊玩規則遊玩，直到達成 ???? 勝利條件
	for !g.Game.IsGameOver() {
		g.Game.PlayRound()
	}

	// 遊戲結束，顯示獲勝者
	fmt.Printf("遊戲結束，獲勝者是: %s\n", g.Game.GetWinner().GetName())
}

// rewrite abstract method
func (g *BaseGame) InitDeck() {
	panic("InitDeck must be implemented")
}

func (g *BaseGame) InitPlayers() {
	panic("InitPlayers must be implemented")
}

func (g *BaseGame) InitPlayRound() {
	panic("InitPlayRound must be implemented")
}

func (g *BaseGame) GetDefaultHandsCount() int {
	panic("GetDefaultHandsCount must be implemented")
}

func (g *BaseGame) PlayRound() {
	panic("PlayRound must be implemented")
}

func (g *BaseGame) IsGameOver() bool {
	panic("IsGameOver must be implemented")
}

func (g *BaseGame) GetWinner() player.Player {
	panic("GetWinner must be implemented")
}

var _ Game = &BaseGame{}
