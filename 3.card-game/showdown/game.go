package main

import "fmt"

const (
	// 需求中未表明實際真人玩家數量，但通常是1人
	RealPlayerCount = 1
	// 交換卡牌的回合數
	ExchangeDuration = 3
	// 遊戲回合數
	GameRound = 13
	// 玩家數量
	PlayerCount = 4
)

// 管理整體流程，回合控制
type Game struct {
	players      [PlayerCount]Player
	currentRound int
	exchanges    []*Exchange
	deck         *Deck
}

func NewGame() *Game {
	return &Game{
		players:      [PlayerCount]Player{},
		currentRound: 0,
		exchanges:    []*Exchange{},
		deck:         nil,
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
	for len(g.deck.GetCards()) > 0 {
		for _, player := range g.players {
			card := g.deck.Draw()
			player.AddHands(card)
		}
	}

	// 開始回合
	for g.getRound() < GameRound {
		g.playRound()
	}
	g.printWinner()
}

func (g *Game) playRound() {
	g.increaseRound()
	fmt.Println("================")
	fmt.Printf("第 %d 回合\n", g.getRound())

	// 檢查交換卡牌時間
	g.checkExchanges()

	// 玩家輪流回合操作
	currentRoundCards := map[*Player]Card{}
	for _, player := range g.players {
		card, err := g.takeTurn(player)
		if err != nil {
			fmt.Printf("回合操作失敗: %s\n", err)
		}

		if card != nil {
			currentRoundCards[&player] = *card
		}
	}

	// 顯示回合出牌
	err := g.ShowCards(currentRoundCards)
	if err != nil {
		fmt.Printf("顯示牌組失敗: %s\n", err)
	}

	// 比較回合出牌
	winner, err := g.CompareCards(currentRoundCards)
	if err != nil {
		fmt.Printf("比大小失敗: %s\n", err)
	}

	fmt.Printf("這回合贏家是 %s\n\n", (*winner).GetName())
	(*winner).AddScore()
}

func (g *Game) takeTurn(player Player) (*Card, error) {
	if !player.UsedPermission() && player.DecideToExchange() {
		var otherPlayers []Player
		for _, p := range g.players {
			if p != player {
				otherPlayers = append(otherPlayers, p)
			}
		}

		playerB, err := player.UseExchange(otherPlayers)
		if err != nil {
			return nil, err
		}

		exchange := exeuteExchange(player, playerB, g.getRound())
		g.addExchange(&exchange)
	}

	return player.Show(), nil
}

func (g *Game) increaseRound() {
	g.currentRound++
}

func (g *Game) getRound() int {
	return g.currentRound
}

func (g *Game) addExchange(exchange *Exchange) {
	g.exchanges = append(g.exchanges, exchange)
}

func (g *Game) getExchanges() []*Exchange {
	return g.exchanges
}

func (g *Game) checkExchanges() {
	for _, exchange := range g.getExchanges() {
		if exchange.GetRound() == g.getRound()-ExchangeDuration {
			exchange.ReturnBackCards()
		}
	}
}

func (g *Game) ShowCards(playerCards map[*Player]Card) error {
	if len(playerCards) == 0 {
		return fmt.Errorf("沒有玩家出牌")
	}

	for player, card := range playerCards {
		fmt.Printf("%s 出了 %s\n", (*player).GetName(), card.String())
	}
	return nil
}

func (g *Game) CompareCards(playerCards map[*Player]Card) (*Player, error) {
	if len(playerCards) == 0 {
		return nil, fmt.Errorf("沒有玩家出牌")
	}

	var maxCard Card
	var maxPlayer *Player
	isFirst := true

	// 遍歷所有玩家的牌
	for player, card := range playerCards {
		if isFirst {
			maxCard = card
			maxPlayer = player
			isFirst = false
			continue
		}

		if card.getRank() > maxCard.getRank() {
			maxCard = card
			maxPlayer = player
		} else if card.getRank() == maxCard.getRank() {
			if card.getSuit() > maxCard.getSuit() {
				maxCard = card
				maxPlayer = player
			}
		}
	}

	return maxPlayer, nil
}

func (g *Game) printWinner() {
	winner := g.players[0]
	for i := 1; i < len(g.players); i++ {
		if g.players[i].GetScore() > winner.GetScore() {
			winner = g.players[i]
		}
	}
	fmt.Printf("勝利者： %s 最終分數: %d\n", winner.GetName(), winner.GetScore())
}
