package game

import (
	"fmt"
	"github/elliot9/card_game_template/player"
)

// 交換玩家雙方手牌
type Exchange struct {
	player1    player.ShowdownPlayer
	player2    player.ShowdownPlayer
	startRound int
}

func executeExchange(playerA player.ShowdownPlayer, playerB player.ShowdownPlayer, startRound int) Exchange {
	playerA.SetUsedPermission()
	tempHands := playerA.GetHands()
	playerA.ReplaceHands(playerB.GetHands())
	playerB.ReplaceHands(tempHands)

	fmt.Printf("[第 %d 回合] 玩家 %s 和 %s 交換手牌\n", startRound, playerA.GetName(), playerB.GetName())

	return Exchange{
		player1:    playerA,
		player2:    playerB,
		startRound: startRound,
	}
}

func (e *Exchange) GetRound() int {
	return e.startRound
}

func (e *Exchange) ReturnBackCards() {
	tempHands := e.player1.GetHands()
	e.player1.ReplaceHands(e.player2.GetHands())
	e.player2.ReplaceHands(tempHands)

	fmt.Printf("玩家 %s 和 %s 換回手牌\n", e.player1.GetName(), e.player2.GetName())
}
