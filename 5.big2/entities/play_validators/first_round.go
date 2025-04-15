package play_validators

import (
	"errors"
	"github/elliot9/big2/core"
)

type FirstRound struct {
	core.BasePlayValidator
}

// 第一回合出牌必須要有梅花三
func (f *FirstRound) IsValid(game core.Big2, cards []core.Card) error {
	if game.GetRound() != 1 {
		return nil
	}

	for _, card := range cards {
		if card.Rank == core.Three && card.Suit == core.Club {
			return nil
		}
	}

	return errors.New("此牌型不合法，請再嘗試一次。")
}

func NewFirstRound() *FirstRound {
	f := &FirstRound{
		BasePlayValidator: core.BasePlayValidator{},
	}
	f.PlayValidator = f
	return f
}

func (f *FirstRound) ShouldRemove(game core.Big2, cards []core.Card) bool {
	// 檢查是否包含梅花三
	for _, card := range cards {
		if card.Rank == core.Three && card.Suit == core.Club {
			return true
		}
	}
	return false
}

var _ core.PlayValidator = (*FirstRound)(nil)
