package triggers

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	fsm "github/elliot9/class15/pkg/fsm/cores"
)

type Command struct {
	EventName string
}

func NewCommand(eventName string) *Command {
	return &Command{
		EventName: eventName,
	}
}

func (c *Command) Match(event fsm.Event, context context.Context) bool {
	message, ok := context.Value(bot.ContextMessage).(community.Message)
	botInstance := context.Value(bot.ContextBot).(*bot.Bot)

	if !ok || event.GetName() != string(community.NewMessage) {
		return false
	}

	// tags 只呼呼叫 bot
	tags := message.Tags
	if len(tags) != 1 || tags[0] != botInstance.GetUser().GetId() {
		return false
	}

	return message.Content == c.EventName
}

var _ fsm.Trigger = &Command{}
