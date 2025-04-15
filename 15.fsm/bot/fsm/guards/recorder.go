package guards

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/bot/fsm/states/record"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	guard "github/elliot9/class15/pkg/fsm/plugins/guards"
)

type Recorder struct {
	cores.Guard
}

func NewRecorder() *Recorder {
	q := &Recorder{}
	q.Guard = guard.NewBaseGuard(func(ctx context.Context) bool {
		user := ctx.Value(bot.ContextUser).(*community.User)
		return record.Recorder.GetId() == user.GetId()
	})
	return q
}
