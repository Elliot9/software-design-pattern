package main

import (
	"github/elliot9/class9/infra/cli"
	"github/elliot9/class9/internal/action"
	"github/elliot9/class9/internal/interfaces"
	"github/elliot9/class9/internal/one_punch_handlers"
	"github/elliot9/class9/internal/role"
	"github/elliot9/class9/internal/rpg"
)

var HPOver500Handler interfaces.OnePunchHandler = one_punch_handlers.NewHPOver500()
var PoisonedOrPetrochemicalHandler interfaces.OnePunchHandler = one_punch_handlers.NewPoisonedOrPetrochemical()
var CheerUpHandler interfaces.OnePunchHandler = one_punch_handlers.NewCheerUp()
var NormalHandler interfaces.OnePunchHandler = one_punch_handlers.NewNormal()

var actions map[string]interfaces.Action = map[string]interfaces.Action{
	"普通攻擊": action.NewBasicAttack(),
	"水球":   action.NewWaterball(),
	"火球":   action.NewFireball(),
	"自我治療": action.NewSelfHealing(),
	"石化":   action.NewPetrochemical(),
	"下毒":   action.NewPoison(),
	"召喚":   action.NewSummon(),
	"自爆":   action.NewSelfExplosion(),
	"鼓舞":   action.NewCheerUp(),
	"詛咒":   action.NewCurse(),
	"一拳攻擊": action.NewOnePunch(HPOver500Handler),
}

func main() {
	cli := cli.NewConsoleIO()

	player := role.NewHero("英雄", 1000, 2000, 200, cli, []interfaces.Action{actions["普通攻擊"], actions["水球"], actions["自我治療"], actions["下毒"], actions["石化"], actions["召喚"]})
	enemy := role.NewSeedAI("巫師", 1000, 2000, 100, cli, []interfaces.Action{actions["普通攻擊"], actions["下毒"]})
	enemy2 := role.NewSeedAI("術士", 1000, 2000, 100, cli, []interfaces.Action{actions["普通攻擊"], actions["鼓舞"]})

	rpg := rpg.NewRPG([]interfaces.Role{player}, []interfaces.Role{enemy, enemy2}, cli)
	rpg.Battle.Start()
	rpg.Battle.PrintResult()
}

func init() {
	HPOver500Handler.SetNext(PoisonedOrPetrochemicalHandler)
	PoisonedOrPetrochemicalHandler.SetNext(CheerUpHandler)
	CheerUpHandler.SetNext(NormalHandler)
}
