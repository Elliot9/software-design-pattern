package big2

import (
	"fmt"
	"github/elliot9/big2/core"
)

type Normal struct {
	core.BaseBig2
}

func NewNormalBig2(cli core.CLI, players []core.Player, deck core.Deck) *Normal {
	n := &Normal{
		BaseBig2: core.BaseBig2{
			Cli: cli,
		},
	}
	n.Big2 = n
	n.SetPlayers(players)
	n.SetDeck(deck)
	return n
}

func (n *Normal) InitialPlayerIndex() {
	for index, player := range n.BaseBig2.GetPlayers() {
		for _, card := range player.GetHands() {
			if card.Rank == core.Three && card.Suit == core.Club {
				n.BaseBig2.SetCurrentPlayerIndex(index)
				return
			}
		}
	}
}

func (n *Normal) IsGameOver() bool {
	return n.getWinner() != nil
}

func (n *Normal) BeforeRound() {
	n.Cli.Println("新的回合開始了。")
}

func (n *Normal) IsRoundOver() bool {
	if n.BaseBig2.GetTopPlayer() == nil {
		return false
	}

	return n.BaseBig2.GetCurrentPlayer() == n.BaseBig2.GetTopPlayer()
}

func (n *Normal) BeforeTakeTurn() {
	n.Cli.Println(fmt.Sprintf("輪到%s了", n.BaseBig2.GetCurrentPlayer().GetName()))
}

func (n *Normal) AfterTakeTurn(playCards []core.Card, pattern core.CardPattern) {
	if len(playCards) == 0 {
		n.Cli.Println(fmt.Sprintf("玩家 %s PASS.", n.BaseBig2.GetCurrentPlayer().GetName()))
		return
	}

	output := fmt.Sprintf("玩家 %s 打出了 %s", n.BaseBig2.GetCurrentPlayer().GetName(), pattern.GetName())
	core.SortCards(playCards)
	for _, card := range playCards {
		output += fmt.Sprintf(" %s", card.String())
	}
	n.Cli.Println(output)
}

func (n *Normal) AfterRound() {

}

func (n *Normal) getWinner() core.Player {
	for _, player := range n.BaseBig2.GetPlayers() {
		if len(player.GetHands()) == 0 {
			return player
		}
	}
	return nil
}

func (n *Normal) PrintWinner() {
	n.Cli.Println(fmt.Sprintf("遊戲結束，遊戲的勝利者為 %s", n.getWinner().GetName()))
	n.Cli.Println("")
}

var _ core.Big2 = (*Normal)(nil)
