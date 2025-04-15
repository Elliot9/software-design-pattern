package normal

import (
	"github/elliot9/class15/bot"
	"github/elliot9/class15/bot/fsm/listeners"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/states"
)

type DefaultConversation struct {
	fsm.State
	bot *bot.Bot
}

func NewDefaultConversation(bot *bot.Bot) *DefaultConversation {
	state := states.NewBaseState("DefaultConversation")

	defaultConversation := &DefaultConversation{
		bot: bot,
	}

	state.AddListeners(map[fsm.Event][]fsm.StateListener{
		cores.NewEvent(string(community.NewMessage)): {
			listeners.NewNewMessageReplyListener([]string{
				"good to hear",
				"thank you",
				"How are you",
			}),
		},
		cores.NewEvent(string(community.NewPost)): {
			listeners.NewNewPostReplyListener("Nice post", false),
		},
	})

	defaultConversation.State = state
	return defaultConversation
}

var _ fsm.State = &DefaultConversation{}
