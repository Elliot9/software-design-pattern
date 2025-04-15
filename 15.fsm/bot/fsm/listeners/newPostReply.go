package listeners

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/stateListeners"
)

type NewPostReplyListener struct {
	fsm.StateListener
	message   string
	markedAll bool
}

func NewNewPostReplyListener(message string, markedAll bool) *NewPostReplyListener {
	listener := &NewPostReplyListener{
		message:   message,
		markedAll: markedAll,
	}

	listener.StateListener = stateListeners.NewBaseStateListener(func(context context.Context) {
		botInstance := context.Value(bot.ContextBot).(*bot.Bot)
		post := context.Value(bot.ContextPost).(community.Post)
		author := context.Value(bot.ContextUser).(*community.User)
		tags := []string{}
		if listener.markedAll {
			tags = append(tags, botInstance.GetUser().GetId())
			for _, user := range botInstance.GetCommunity().GetOnlineUsers() {
				tags = append(tags, user.GetId())
			}
		} else {
			tags = append(tags, author.GetId())
		}
		botInstance.ReplyComment(post.Id, listener.message, tags)
	})
	return listener
}

var _ fsm.StateListener = &NewPostReplyListener{}
