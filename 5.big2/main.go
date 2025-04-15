package main

import (
	"github/elliot9/big2/common"
	"github/elliot9/big2/entities/big2"
	"github/elliot9/big2/entities/deck"
	"github/elliot9/big2/infra/cli"
)

func main() {
	cli := cli.NewConsoleIO()
	players := common.InitializePlayers(cli)
	deck := deck.NewNormalDeck()
	big2 := big2.NewNormalBig2(cli, players, deck)

	// 設定牌型
	cardPattern := common.InitializeCardPatterns()
	big2.SetCardPattern(cardPattern)

	// 設定出牌規則
	playValidator := common.InitializePlayValidator()
	big2.SetPlayValidator(playValidator)

	// 開始遊戲
	big2.Start()
}
