package game

import (
	"fmt"
	"github/elliot9/card_game_template/card"
	"github/elliot9/card_game_template/deck"
	"github/elliot9/card_game_template/player"
	"strings"
)

const (
	GameRound        = 13
	ExchangeDuration = 3
)

type ShowdownGame struct {
	BaseGame
	currentRound int
	exchanges    []*Exchange
}

func NewShowdownGame() *ShowdownGame {
	game := &ShowdownGame{
		BaseGame: BaseGame{
			players: make([]player.Player, PlayerCount),
		},
		currentRound: 0,
		exchanges:    make([]*Exchange, 0),
	}

	game.Game = game
	return game
}

func (g *ShowdownGame) InitDeck() {
	cards := []card.CardStrategy{}

	for _, suit := range []card.Suit{card.Club, card.Diamond, card.Heart, card.Spade} {
		for _, rank := range []card.Rank{card.Two, card.Three, card.Four, card.Five, card.Six, card.Seven, card.Eight, card.Nine, card.Ten, card.Jack, card.Queen, card.King, card.Ace} {
			cards = append(cards, card.NewPokerCard(rank, suit))
		}
	}

	g.deck = deck.NewShowdownDeck(cards)
}

func (g *ShowdownGame) InitPlayers() {
	for i := 0; i < PlayerCount; i++ {
		if i < RealPlayerCount {
			g.players[i] = &player.HumanShowdownPlayer{Cli: player.NewConsoleIO()}
		} else {
			g.players[i] = &player.AIShowdownPlayer{}
		}
	}
}

func (g *ShowdownGame) GetDefaultHandsCount() int {
	return 13
}

func (g *ShowdownGame) InitPlayRound() {}

func (g *ShowdownGame) IsGameOver() bool {
	return g.currentRound > GameRound
}

func (g *ShowdownGame) PlayRound() {
	g.increaseRound()
	fmt.Println("================")
	fmt.Printf("第 %d 回合\n", g.getRound())

	// 檢查交換卡牌時間
	g.checkExchanges()

	// 玩家輪流回合操作
	currentRoundCards := map[player.Player]card.CardStrategy{}
	for _, player := range g.players {
		card, err := g.takeTurn(player)
		if err != nil {
			fmt.Printf("回合操作失敗: %s\n", err)
		}

		if card != nil {
			currentRoundCards[player] = card
		}
	}

	// 顯示回合出牌
	err := g.showCards(currentRoundCards)
	if err != nil {
		fmt.Printf("顯示牌組失敗: %s\n", err)
	}

	// 比較回合出牌
	winner, err := g.compareCards(currentRoundCards)
	if err != nil {
		fmt.Printf("比大小失敗: %s\n", err)
	}

	fmt.Printf("這回合贏家是 %s\n\n", winner.GetName())
	showdownPlayer := winner.(player.ShowdownPlayer)
	showdownPlayer.AddScore()
}

func (g *ShowdownGame) GetWinner() player.Player {
	winner := g.players[0].(player.ShowdownPlayer)
	for i := 1; i < len(g.players); i++ {
		showdownPlayer := g.players[i].(player.ShowdownPlayer)
		if showdownPlayer.GetScore() > winner.GetScore() {
			winner = showdownPlayer
		}
	}
	return winner
}
func (g *ShowdownGame) takeTurn(p player.Player) (card.CardStrategy, error) {
	showdownPlayer := p.(player.ShowdownPlayer)

	if !showdownPlayer.UsedPermission() && showdownPlayer.DecideToExchange() {
		var otherPlayers []player.ShowdownPlayer
		for _, p := range g.players {
			if p != showdownPlayer {
				otherPlayers = append(otherPlayers, p.(player.ShowdownPlayer))
			}
		}
		playerB, err := showdownPlayer.UseExchange(otherPlayers)
		if err != nil {
			return nil, err
		}
		exchange := executeExchange(showdownPlayer, playerB, g.getRound())
		g.addExchange(&exchange)
	}
	return showdownPlayer.Show(), nil
}

func (g *ShowdownGame) increaseRound() {
	g.currentRound++
}

func (g *ShowdownGame) getRound() int {
	return g.currentRound
}

func (g *ShowdownGame) getExchanges() []*Exchange {
	return g.exchanges
}

func (g *ShowdownGame) addExchange(exchange *Exchange) {
	g.exchanges = append(g.exchanges, exchange)
}

func (g *ShowdownGame) checkExchanges() {
	for _, exchange := range g.getExchanges() {
		if exchange.GetRound() == g.getRound()-ExchangeDuration {
			exchange.ReturnBackCards()
		}
	}
}

func (g *ShowdownGame) showCards(playerCards map[player.Player]card.CardStrategy) error {
	if len(playerCards) == 0 {
		return fmt.Errorf("沒有玩家出牌")
	}

	for player, card := range playerCards {
		fmt.Printf("%s 出了 %s\n", player.GetName(), card.Identify())
	}
	return nil
}
func (g *ShowdownGame) compareCards(playerCards map[player.Player]card.CardStrategy) (player.Player, error) {
	if len(playerCards) == 0 {
		return nil, fmt.Errorf("沒有玩家出牌")
	}

	var maxCard card.CardStrategy
	var maxPlayer player.Player
	isFirst := true

	rankMap := map[string]card.Rank{
		"2":  card.Two,
		"3":  card.Three,
		"4":  card.Four,
		"5":  card.Five,
		"6":  card.Six,
		"7":  card.Seven,
		"8":  card.Eight,
		"9":  card.Nine,
		"10": card.Ten,
		"J":  card.Jack,
		"Q":  card.Queen,
		"K":  card.King,
		"A":  card.Ace,
	}

	suitMap := map[string]card.Suit{
		"♣": card.Club,
		"♦": card.Diamond,
		"♥": card.Heart,
		"♠": card.Spade,
	}

	// 遍歷所有玩家的牌
	for player, card := range playerCards {
		if isFirst {
			maxCard = card
			maxPlayer = player
			isFirst = false
			continue
		}

		maxCardRank, maxCardSuit := strings.Split(maxCard.Identify(), "-")[0], strings.Split(maxCard.Identify(), "-")[1]
		cardRank, cardSuit := strings.Split(card.Identify(), "-")[0], strings.Split(card.Identify(), "-")[1]

		if rankMap[cardRank] > rankMap[maxCardRank] {
			maxCard = card
			maxPlayer = player
		} else if rankMap[cardRank] == rankMap[maxCardRank] {
			if suitMap[cardSuit] > suitMap[maxCardSuit] {
				maxCard = card
				maxPlayer = player
			}
		}
	}

	return maxPlayer, nil
}

var _ Game = &ShowdownGame{}
