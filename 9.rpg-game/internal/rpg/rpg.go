package rpg

import (
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/battle"
	"github/elliot9/class9/internal/interfaces"
)

type RPG struct {
	Battle *battle.Battle
}

func NewRPG(troop1 []interfaces.Role, troop2 []interfaces.Role, cli cli.CLI) *RPG {
	playerTroop := battle.NewTroop(1)
	for _, role := range troop1 {
		playerTroop.AddRole(role)
	}

	enemyTroop := battle.NewTroop(2)
	for _, role := range troop2 {
		enemyTroop.AddRole(role)
	}

	battle := battle.NewBattle(playerTroop, enemyTroop, cli)

	return &RPG{
		Battle: battle,
	}
}
