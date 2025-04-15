package normal

import (
	"github/elliot9/class15/bot"
	"github/elliot9/class15/bot/fsm/listeners"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/states"
)

type Interacting struct {
	fsm.State
	bot *bot.Bot
}

func NewInteracting(bot *bot.Bot) *Interacting {
	state := states.NewBaseState("Interacting")

	interacting := &Interacting{
		bot: bot,
	}

	state.AddListeners(map[fsm.Event][]fsm.StateListener{
		cores.NewEvent(string(community.NewMessage)): {
			listeners.NewNewMessageReplyListener([]string{
				"Hi hi😁",
				"I like your idea!",
			}),
		},
		cores.NewEvent(string(community.NewPost)): {
			listeners.NewNewPostReplyListener("How do you guys think about it?", true),
		},
	})

	interacting.State = state
	return interacting
}

var _ fsm.State = &Interacting{}
