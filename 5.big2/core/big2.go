package core

import (
	"fmt"
	"strings"
)

const (
	PlayerCount = 4
)

type Big2 interface {
	// Setting Game Rules, Card Patterns
	SetCardPattern(pattern CardPattern)
	SetPlayValidator(validator PlayValidator)
	GetPlayValidator() PlayValidator
	RemovePlayValidator(validator PlayValidator)

	// Game LifeCycle
	Start()
	InitialPlayerIndex()
	IsGameOver() bool
	GetWinner() Player
	PrintWinner()

	// Round LifeCycle
	IsRoundOver() bool
	BeforeRound()
	AfterRound() // Currently not needed for Normal rules

	// Turn LifeCycle
	BeforeTakeTurn()
	TakeTurn() ([]Card, CardPattern)
	AfterTakeTurn(playCards []Card, pattern CardPattern)

	// Game Rules
	IsValidPlay(cards []Card) error
	IsValidCardPattern(cards []Card) (bool, CardPattern)

	// Game State
	GetRound() int
	GetCurrentPlayer() Player
	GetTopPlayer() Player
	GetTopPlay() []Card
}

type BaseBig2 struct {
	Cli                CLI
	currentRound       int
	players            []Player
	currentPlayerIndex int
	topPlay            []Card
	topPlayerIndex     int
	deck               Deck
	cardPattern        CardPattern
	playValidator      PlayValidator
	Big2
}

// 遊戲主流程
func (b *BaseBig2) Start() {
	// 1. 玩家命名
	b.namePlayer()

	// 1-2.發牌
	b.dealCards()

	// 1-3. 設定第一個出牌玩家
	b.Big2.InitialPlayerIndex()

	// 2. 遊戲主循環
	for !b.Big2.IsGameOver() {

		// 清除牌頂玩家 頂牌
		b.resetTopPlay()
		b.resetTopPlayerIndex()
		b.increaseRound()

		// 3. 回合循環
		b.Big2.BeforeRound()
		for !b.Big2.IsGameOver() && !b.Big2.IsRoundOver() {
			// 4. 玩家操作前
			b.Big2.BeforeTakeTurn()
			// 4-2. 玩家操作
			playCards, pattern := b.TakeTurn()
			// 4-3. 玩家操作後
			b.Big2.AfterTakeTurn(playCards, pattern)
			// 4-4. 更新 validator
			b.updateValidator(playCards)
			// 4-5. 換下一個玩家
			b.nextPlayer()
		}
		b.Big2.AfterRound()
	}

	// 5. 遊戲結束
	b.Big2.PrintWinner()
}

func (b *BaseBig2) namePlayer() {
	for _, player := range b.players {
		player.NameSelf()
	}
}

func (b *BaseBig2) dealCards() {
	currentIndex := 0
	for !b.GetDeck().IsEmpty() {
		b.players[currentIndex].AddHands([]Card{b.GetDeck().Draw()})
		currentIndex = (currentIndex + 1) % len(b.players)
	}
}

// @TODO: 設定第一個出牌玩家
func (b *BaseBig2) InitialPlayerIndex() {
	panic("big 2 InitialPlayerIndex not implemented")
}

// @TODO: 遊戲是否結束
func (b *BaseBig2) IsGameOver() bool {
	panic("big 2 IsGameOver not implemented")
}

// @TODO: 回合開始前
func (b *BaseBig2) BeforeRound() {
	panic("big 2 BeforeRound not implemented")
}

// @TODO: 回合是否結束
func (b *BaseBig2) IsRoundOver() bool {
	panic("big 2 IsRoundOver not implemented")
}

// @TODO: 玩家操作前
func (b *BaseBig2) BeforeTakeTurn() {
	panic("big 2 BeforeTakeTurn not implemented")
}

func (b *BaseBig2) printHands(player Player) {
	strIndex := ""
	strCards := ""

	for i, card := range player.GetHands() {
		strIndex += fmt.Sprintf("%d", i)
		strCard := fmt.Sprintf("%s ", card.String())
		strCards += strCard
		strIndex += strings.Repeat(" ", len(strCard)-len(fmt.Sprintf("%d", i)))
	}
	b.Cli.Println(strIndex)
	b.Cli.Println(strCards)
}

func (b *BaseBig2) TakeTurn() ([]Card, CardPattern) {
	currentPlayer := b.GetCurrentPlayer()

	// 印出玩家手牌
	b.printHands(currentPlayer)

	// 玩家出牌
	playCards := currentPlayer.Play()

	// 驗證有效操作
	if err := b.IsValidPlay(playCards); err != nil {
		currentPlayer.AddHands(playCards)
		b.Cli.Println(err.Error())
		return b.TakeTurn()
	}

	var pattern CardPattern
	var isValidPattern bool

	// 玩家有出牌
	if len(playCards) != 0 {
		// 判斷牌型
		isValidPattern, pattern = b.IsValidCardPattern(playCards)

		if !isValidPattern {
			currentPlayer.AddHands(playCards)
			b.Cli.Println("此牌型不合法，請再嘗試一次。")
			return b.TakeTurn()
		}

		b.setTopPlay(playCards)
		b.setTopPlayerIndex(b.GetCurrentPlayerIndex())
	}

	return playCards, pattern
}

func (b *BaseBig2) updateValidator(playCards []Card) {
	if b.GetPlayValidator() != nil {
		validator := b.GetPlayValidator()
		for validator != nil {
			if validator.ShouldRemove(b, playCards) {
				b.RemovePlayValidator(validator)
			}
			validator = validator.GetNext()
		}
	}
}

// @TODO: 玩家操作後
func (b *BaseBig2) AfterTakeTurn(playCards []Card, pattern CardPattern) {
	panic("big 2 AfterTakeTurn not implemented")
}

// @TODO: 回合結束
func (b *BaseBig2) AfterRound() {
	panic("big 2 AfterRound not implemented")
}

// @TODO: 取得勝利者
func (b *BaseBig2) GetWinner() Player {
	panic("big 2 GetWinner not implemented")
}

func (b *BaseBig2) PrintWinner() {
	b.Cli.Println(fmt.Sprintf("遊戲結束，遊戲的勝利者為 %s", b.GetWinner().GetName()))
}

func (b *BaseBig2) IsValidPlay(cards []Card) error {
	if b.playValidator != nil {
		return b.playValidator.IsValid(b, cards)
	}

	return nil
}

func (b *BaseBig2) IsValidCardPattern(cards []Card) (bool, CardPattern) {
	if b.cardPattern != nil {
		return b.cardPattern.Check(cards, b.GetTopPlay())
	}

	return false, nil
}

func (b *BaseBig2) GetRound() int {
	return b.currentRound
}

func (b *BaseBig2) GetDeck() Deck {
	return b.deck
}

func (b *BaseBig2) increaseRound() {
	b.currentRound++
}

func (b *BaseBig2) GetCurrentPlayerIndex() int {
	return b.currentPlayerIndex
}

func (b *BaseBig2) GetCurrentPlayer() Player {
	return b.players[b.GetCurrentPlayerIndex()]
}

func (b *BaseBig2) SetCurrentPlayerIndex(index int) {
	b.currentPlayerIndex = index
}

func (b *BaseBig2) setTopPlayerIndex(index int) {
	b.topPlayerIndex = index
}

func (b *BaseBig2) resetTopPlayerIndex() {
	b.setTopPlayerIndex(-1)
}

func (b *BaseBig2) GetTopPlayerIndex() int {
	return b.topPlayerIndex
}

func (b *BaseBig2) resetTopPlay() {
	b.setTopPlay([]Card{})
}

func (b *BaseBig2) setTopPlay(cards []Card) {
	b.topPlay = cards
}

func (b *BaseBig2) GetTopPlay() []Card {
	return b.topPlay
}

func (b *BaseBig2) GetTopPlayer() Player {
	if b.GetTopPlayerIndex() == -1 {
		return nil
	}
	return b.players[b.GetTopPlayerIndex()]
}

func (b *BaseBig2) nextPlayer() {
	b.currentPlayerIndex = (b.currentPlayerIndex + 1) % len(b.players)
}

func (b *BaseBig2) SetPlayers(players []Player) {
	b.players = players
}

func (b *BaseBig2) GetPlayers() []Player {
	return b.players
}

func (b *BaseBig2) SetDeck(deck Deck) {
	b.deck = deck
}

func (b *BaseBig2) SetCardPattern(pattern CardPattern) {
	b.cardPattern = pattern
}

func (b *BaseBig2) SetPlayValidator(validator PlayValidator) {
	b.playValidator = validator
}

func (b *BaseBig2) GetPlayValidator() PlayValidator {
	return b.playValidator
}

func (b *BaseBig2) RemovePlayValidator(validator PlayValidator) {
	b.playValidator = b.playValidator.RemoveValidator(validator)
}

var _ Big2 = (*BaseBig2)(nil)
