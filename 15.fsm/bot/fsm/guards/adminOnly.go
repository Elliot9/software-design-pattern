package guards

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	guard "github/elliot9/class15/pkg/fsm/plugins/guards"
)

type AdminOnly struct {
	cores.Guard
}

func NewAdminOnly() *AdminOnly {
	q := &AdminOnly{}
	q.Guard = guard.NewBaseGuard(func(ctx context.Context) bool {
		user := ctx.Value(bot.ContextUser).(*community.User)
		return user.IsAdmin()
	})
	return q
}
