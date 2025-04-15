package record

import (
	"github/elliot9/class15/bot"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/states"
)

type Waiting struct {
	fsm.State
	bot *bot.Bot
}

func NewWaiting(bot *bot.Bot) *Waiting {
	state := states.NewBaseState("Waiting")

	waiting := &Waiting{
		bot: bot,
	}

	waiting.State = state
	return waiting
}

var _ fsm.State = &Waiting{}
