package main

import "fmt"

const (
	// 需求中未表明實際真人玩家數量，但通常是1人
	RealPlayerCount = 1
	// 玩家數量
	PlayerCount = 4
	// 抽牌階段預設抽牌數量
	DrawCardCount = 5
)

// 管理整體流程，回合控制，遊戲規則
type Game struct {
	players    [PlayerCount]Player
	deck       *Deck
	boardCards []Card
}

func NewGame() *Game {
	return &Game{
		players:    [PlayerCount]Player{},
		deck:       nil,
		boardCards: []Card{},
	}
}

func (g *Game) initDeck() {
	g.deck = NewDeck()
}

func (g *Game) Start() {
	// 初始化牌堆
	g.initDeck()

	// 初始化玩家
	for i := 0; i < PlayerCount; i++ {
		if i < RealPlayerCount {
			g.players[i] = &HumanPlayer{cli: NewConsoleIO()}
		} else {
			g.players[i] = &AIPlayer{}
		}
	}

	// 玩家為自己取名
	for _, player := range g.players {
		player.NameSelf()
	}

	// 洗牌
	g.deck.Shuffle()

	// 發牌
	for i := 0; i < DrawCardCount; i++ {
		for _, player := range g.players {
			card := g.deck.Draw()
			player.AddHands(card)
		}
	}

	// 遊戲開始時，從牌堆中翻出第一張牌到檯面上
	g.boardCards = append(g.boardCards, g.deck.Draw())

	// 開始回合
	for !g.isGameOver() {
		g.playRound()
	}
	g.printWinner()
}

func (g *Game) isGameOver() bool {
	for _, player := range g.players {
		if len(player.GetHands()) == 0 {
			return true
		}
	}
	return false
}

func (g *Game) playRound() {
	fmt.Println("================")

	for _, player := range g.players {
		card := g.takeTurn(player)
		g.boardCards = append(g.boardCards, *card)
	}
}

func (g *Game) takeTurn(player Player) *Card {
	fmt.Printf("當前檯面上的牌: %v\n", g.getTopBoardCard())
	for !g.hasValidCard(player) {
		if g.deck.IsEmpty() {
			g.resetBoardCards()
		}

		newCard := g.deck.Draw()
		player.AddHands(newCard)
	}

	for {
		card := player.Show()
		if g.isValidCard(card) {
			fmt.Printf("玩家 %s 打出了 %v\n", player.GetName(), card)
			return card
		}

		player.AddHands(*card)
	}
}

func (g *Game) hasValidCard(player Player) bool {
	for _, card := range player.GetHands() {
		if g.isValidCard(&card) {
			return true
		}
	}
	return false
}

func (g *Game) isValidCard(card *Card) bool {
	if card == nil {
		return false
	}
	topCard := g.getTopBoardCard()
	return card.getNumber() == topCard.getNumber() || card.getColor() == topCard.getColor()
}

func (g *Game) printWinner() {
	for _, player := range g.players {
		if len(player.GetHands()) == 0 {
			fmt.Printf("勝利者： %s\n", player.GetName())
		}
	}
}

func (g *Game) getTopBoardCard() *Card {
	return &g.boardCards[len(g.boardCards)-1]
}

func (g *Game) resetBoardCards() {
	g.deck.AddCards(g.boardCards[:len(g.boardCards)-1])
	g.boardCards = g.boardCards[:len(g.boardCards)-1]
	g.deck.Shuffle()
}
