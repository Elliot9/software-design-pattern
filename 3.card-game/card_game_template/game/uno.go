package game

import (
	"fmt"
	"github/elliot9/card_game_template/card"
	"github/elliot9/card_game_template/deck"
	"github/elliot9/card_game_template/player"
	"strings"
)

type UnoGame struct {
	BaseGame
	boardCards []card.CardStrategy
}

func NewUnoGame() *UnoGame {
	game := &UnoGame{
		BaseGame: BaseGame{
			players: make([]player.Player, PlayerCount),
		},
		boardCards: make([]card.CardStrategy, 0),
	}

	game.Game = game
	return game
}

func (g *UnoGame) InitDeck() {
	cards := []card.CardStrategy{}

	for _, color := range []card.Color{card.Red, card.Yellow, card.Green, card.Blue} {
		for _, number := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			cards = append(cards, card.NewUnoCard(number, color))
		}
	}

	g.deck = deck.NewUnoDeck(cards)
}

func (g *UnoGame) InitPlayers() {
	for i := 0; i < PlayerCount; i++ {
		if i < RealPlayerCount {
			g.players[i] = &player.HumanUnoPlayer{Cli: player.NewConsoleIO()}
		} else {
			g.players[i] = &player.AIUnoPlayer{}
		}
	}
}

func (g *UnoGame) GetDefaultHandsCount() int {
	return 5
}

func (g *UnoGame) InitPlayRound() {
	// 遊戲開始時，從牌堆中翻出第一張牌到檯面上
	g.addBoardCard(g.deck.Draw())
}

func (g *UnoGame) IsGameOver() bool {
	return g.GetWinner() != nil
}

func (g *UnoGame) PlayRound() {
	fmt.Println("================")

	for _, player := range g.players {
		card := g.takeTurn(player)
		g.addBoardCard(card)
	}
}

func (g *UnoGame) GetWinner() player.Player {
	for _, player := range g.players {
		if len(player.GetHands()) == 0 {
			return player
		}
	}
	return nil
}

func (g *UnoGame) takeTurn(player player.Player) card.CardStrategy {
	fmt.Printf("當前檯面上的牌: %v\n", g.getTopBoardCard())
	for !g.hasValidCard(player) {
		unoDeck := g.deck.(*deck.UnoDeck)
		if unoDeck.IsEmpty() {
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

		player.AddHands(card)
	}
}

func (g *UnoGame) addBoardCard(card card.CardStrategy) {
	g.boardCards = append(g.boardCards, card)
}

func (g *UnoGame) getTopBoardCard() card.CardStrategy {
	return g.boardCards[len(g.boardCards)-1]
}

func (g *UnoGame) hasValidCard(player player.Player) bool {
	for _, card := range player.GetHands() {
		if g.isValidCard(card) {
			return true
		}
	}
	return false
}

func (g *UnoGame) isValidCard(card card.CardStrategy) bool {
	if card == nil {
		return false
	}
	topCard := g.getTopBoardCard()

	validNumber := strings.Split(card.Identify(), "-")[1] == strings.Split(topCard.Identify(), "-")[1]
	validColor := strings.Split(card.Identify(), "-")[0] == strings.Split(topCard.Identify(), "-")[0]

	return validNumber || validColor
}

func (g *UnoGame) resetBoardCards() {
	unoDeck := g.deck.(*deck.UnoDeck)
	unoDeck.AddCards(g.boardCards[:len(g.boardCards)-1])
	g.boardCards = g.boardCards[:len(g.boardCards)-1]
	unoDeck.Shuffle()
}

var _ Game = &UnoGame{}
