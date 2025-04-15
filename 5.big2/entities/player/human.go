package player

import (
	"github/elliot9/big2/core"
)

type Human struct {
	core.BasePlayer
	cli core.CLI
}

func NewHuman(cli core.CLI) *Human {
	human := &Human{
		BasePlayer: core.BasePlayer{},
		cli:        cli,
	}
	human.Player = human
	return human
}

func (h *Human) PlayCards() []core.Card {
	inputs := h.cli.ReadNumber()

	cards := []core.Card{}
	// -1 表示不出牌
	if len(inputs) == 1 && inputs[0] == -1 {
		return cards
	}

	for _, input := range inputs {
		cards = append(cards, h.BasePlayer.GetHands()[input])
	}

	return cards
}

func (h *Human) NameSelf() {
	name := h.cli.ReadLine()
	h.BasePlayer.SetName(name)
}

var _ core.Player = &Human{}
