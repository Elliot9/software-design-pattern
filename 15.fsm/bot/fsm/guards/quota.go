package guards

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/pkg/fsm/cores"
	guard "github/elliot9/class15/pkg/fsm/plugins/guards"
)

type Quota struct {
	Number int
	cores.Guard
}

func NewQuota(number int) *Quota {
	q := &Quota{Number: number}
	q.Guard = guard.NewBaseGuard(func(ctx context.Context) bool {
		bot := ctx.Value(bot.ContextBot).(*bot.Bot)
		return bot.GetQuota() >= number
	})
	return q
}
