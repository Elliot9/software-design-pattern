package common

import (
	"github/elliot9/big2/core"
	"github/elliot9/big2/entities/card_patterns"
	"github/elliot9/big2/entities/play_validators"
	"github/elliot9/big2/entities/player"
)

func InitializePlayers(cli core.CLI) []core.Player {
	players := make([]core.Player, core.PlayerCount)
	for i := 0; i < core.PlayerCount; i++ {
		players[i] = player.NewHuman(cli)
	}
	return players
}

func InitializeCardPatterns() core.CardPattern {
	singlePattern := card_patterns.NewSingle()
	pairPattern := card_patterns.NewPair()
	straightPattern := card_patterns.NewStraight()
	fullHousePattern := card_patterns.NewFullHouse()

	singlePattern.SetNext(pairPattern)
	pairPattern.SetNext(straightPattern)
	straightPattern.SetNext(fullHousePattern)

	return singlePattern
}

func InitializePlayValidator() core.PlayValidator {
	firstRoundValidator := play_validators.NewFirstRound()
	passValidator := play_validators.NewPass()

	passValidator.SetNext(firstRoundValidator)

	return passValidator
}
