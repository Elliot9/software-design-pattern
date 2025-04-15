package play_validators

import (
	"errors"
	"github/elliot9/big2/core"
)

type Pass struct {
	core.BasePlayValidator
}

func (p *Pass) IsValid(game core.Big2, cards []core.Card) error {
	if len(cards) != 0 {
		return nil
	}

	if game.GetTopPlayer() == nil {
		return errors.New("你不能在新的回合中喊 PASS")
	}

	return nil
}

func NewPass() *Pass {
	p := &Pass{
		BasePlayValidator: core.BasePlayValidator{},
	}
	p.PlayValidator = p
	return p
}

var _ core.PlayValidator = (*Pass)(nil)
