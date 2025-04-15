package guards

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/pkg/fsm/cores"
	guard "github/elliot9/class15/pkg/fsm/plugins/guards"
)

type OnlineUserNumberGuard struct {
	numberGuard NumberGuard
	cores.Guard
}

func NewOnlineUserNumberGuard(numberGuard NumberGuard) *OnlineUserNumberGuard {
	g := &OnlineUserNumberGuard{numberGuard: numberGuard}
	g.Guard = guard.NewBaseGuard(func(ctx context.Context) bool {
		bot := ctx.Value(bot.ContextBot).(*bot.Bot)
		return numberGuard(len(bot.GetCommunity().GetOnlineUsers()))
	})
	return g
}

var _ cores.Guard = &OnlineUserNumberGuard{}
