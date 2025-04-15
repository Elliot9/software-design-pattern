package listeners

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/stateListeners"
)

type NewMessageReplyListener struct {
	fsm.StateListener
	messages []string
	pivot    int
}

func NewNewMessageReplyListener(replyMessages []string) *NewMessageReplyListener {
	listener := &NewMessageReplyListener{
		messages: replyMessages,
		pivot:    0,
	}

	listener.StateListener = stateListeners.NewBaseStateListener(func(context context.Context) {
		message := listener.messages[listener.pivot%len(listener.messages)]
		botInstance := context.Value(bot.ContextBot).(*bot.Bot)
		user := context.Value(bot.ContextUser).(*community.User)
		botInstance.ReplyMessage(user.GetId(), message)
		listener.pivot++
	})
	return listener
}

var _ fsm.StateListener = &NewMessageReplyListener{}
