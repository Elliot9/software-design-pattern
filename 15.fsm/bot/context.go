package bot

import (
	"context"
	"github/elliot9/class15/community"
)

type ContextKey string

const (
	ContextBot          ContextKey = "bot"
	ContextComment      ContextKey = "comment"
	ContextMessage      ContextKey = "message"
	ContextPost         ContextKey = "post"
	ContextSpeakContent ContextKey = "speakContent"
	ContextUser         ContextKey = "user"
	ContextWinner       ContextKey = "winner"
)

func (b *Bot) createContext(event community.Event[any]) context.Context {
	ctx := context.WithValue(context.Background(), ContextBot, b)

	switch event.GetType() {
	case community.NewMessage:
		ctx = b.withMessage(ctx, event.GetData())
	case community.NewPost:
		ctx = b.withPost(ctx, event.GetData())
	case community.NewComment:
		ctx = b.withComment(ctx, event.GetData())
	case community.NewSpeak:
		ctx = b.withSpeak(ctx, event.GetData())
	case community.GoBroadcasting, community.StopBroadcasting:
		ctx = b.withUser(ctx, event.GetData())
	}

	return ctx
}

func (b *Bot) withMessage(ctx context.Context, data any) context.Context {
	ctx = b.withUser(ctx, data)
	return b.withContextValue(ctx, data, community.EventDataMessage, ContextMessage)
}

func (b *Bot) withPost(ctx context.Context, data any) context.Context {
	ctx = b.withUser(ctx, data)
	return b.withContextValue(ctx, data, community.EventDataPost, ContextPost)
}

func (b *Bot) withComment(ctx context.Context, data any) context.Context {
	ctx = b.withUser(ctx, data)
	return b.withContextValue(ctx, data, community.EventDataComment, ContextComment)
}

func (b *Bot) withSpeak(ctx context.Context, data any) context.Context {
	ctx = b.withUser(ctx, data)
	return b.withContextValue(ctx, data, community.EventDataSpeak, ContextSpeakContent)
}

func (b *Bot) withUser(ctx context.Context, data any) context.Context {
	return b.withContextValue(ctx, data, community.EventDataUser, ContextUser)
}

func (b *Bot) withContextValue(ctx context.Context, data any, key community.EventDataKey, contextKey ContextKey) context.Context {
	dataMap, ok := data.(map[community.EventDataKey]interface{})

	if !ok {
		return ctx
	}

	value, ok := dataMap[key]
	if !ok {
		return ctx
	}
	ctx = context.WithValue(ctx, contextKey, value)
	return ctx
}
