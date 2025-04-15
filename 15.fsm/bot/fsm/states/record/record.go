package record

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/states"
	"github/elliot9/class15/pkg/fsm/plugins/transitions"
	"github/elliot9/class15/pkg/fsm/plugins/triggers"
)

var Recorder *community.User

type Record struct {
	fsm.State
	bot           *bot.Bot
	quotaRequired int
}

func NewRecord(bot *bot.Bot) *Record {
	baseFsm := fsm.NewFiniteStateMachine(func() cores.State {
		community := bot.GetCommunity()
		if community.GetRecorder() != nil {
			return NewRecording(bot)
		}
		return NewWaiting(bot)
	}, []cores.Transition{
		transitions.NewBaseTransition(triggers.NewBaseTrigger(cores.NewEvent(string(community.GoBroadcasting))),
			nil,
			NewWaiting(bot),
			NewRecording(bot),
			[]cores.Action{}),
		transitions.NewBaseTransition(triggers.NewBaseTrigger(cores.NewEvent(string(community.StopBroadcasting))),
			nil,
			NewRecording(bot),
			NewWaiting(bot),
			[]cores.Action{}),
	})

	state := states.NewSubStateMachine(baseFsm, "Record")

	record := &Record{
		bot:           bot,
		quotaRequired: 3,
	}

	record.State = state
	return record
}

func (r *Record) OnEntry(context context.Context) context.Context {
	r.bot.SetQuota(r.bot.GetQuota() - r.quotaRequired)
	Recorder = context.Value(bot.ContextUser).(*community.User)
	return r.State.OnEntry(context)
}

var _ fsm.State = &Record{}
