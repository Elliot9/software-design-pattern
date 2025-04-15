package battle

import (
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/role"
)

type Battle struct {
	PlayerTroop interfaces.Troop
	EnemyTroop  interfaces.Troop
	CLI         cli.CLI
}

func NewBattle(playerTroop interfaces.Troop, enemyTroop interfaces.Troop, cli cli.CLI) *Battle {
	b := &Battle{
		PlayerTroop: playerTroop,
		EnemyTroop:  enemyTroop,
		CLI:         cli,
	}

	playerTroop.SetBattle(b)
	enemyTroop.SetBattle(b)

	return b
}

func (b *Battle) Start() {
	for {
		for !b.IsGameOver() {
			currentRole := b.PlayerTroop.GetNextRole()
			if currentRole == nil {
				break
			}
			currentRole.TakeTurn()
		}
		for !b.IsGameOver() {
			currentRole := b.EnemyTroop.GetNextRole()
			if currentRole == nil {
				break
			}
			currentRole.TakeTurn()
		}

		if b.IsGameOver() {
			break
		}
	}
}

func (b *Battle) IsGameOver() bool {
	if len(b.PlayerTroop.GetRoles()) == 0 || len(b.EnemyTroop.GetRoles()) == 0 {
		return true
	}
	return b.isHeroDead()
}

func (b *Battle) GetEnemyTroop(troop interfaces.Troop) interfaces.Troop {
	if troop == b.PlayerTroop {
		return b.EnemyTroop
	}
	return b.PlayerTroop
}

func (b *Battle) PrintResult() {
	if b.isHeroDead() {
		b.CLI.Println("你失敗了！")
	} else {
		b.CLI.Println("你獲勝了！")
	}
	b.CLI.Println("")
}

func (b *Battle) isHeroDead() bool {
	if len(b.PlayerTroop.GetRoles()) == 0 {
		return true
	}
	_, ok := b.PlayerTroop.GetRoles()[0].(*role.Hero)
	return !ok
}

var _ interfaces.Battle = &Battle{}
