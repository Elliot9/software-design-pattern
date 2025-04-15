package normal

import (
	"github/elliot9/class15/bot"
	"github/elliot9/class15/bot/fsm/guards"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/states"
	"github/elliot9/class15/pkg/fsm/plugins/transitions"
	"github/elliot9/class15/pkg/fsm/plugins/triggers"
)

type Normal struct {
	fsm.State
	bot *bot.Bot
}

const (
	// 10 - 1(電腦) = 9
	OnlineUserGuardNumber = 9
)

func NewNormal(bot *bot.Bot) *Normal {
	baseFsm := fsm.NewFiniteStateMachine(func() cores.State {
		community := bot.GetCommunity()
		if len(community.GetOnlineUsers()) < OnlineUserGuardNumber {
			return NewDefaultConversation(bot)
		}
		return NewInteracting(bot)
	}, []cores.Transition{
		transitions.NewBaseTransition(triggers.NewBaseTrigger(cores.NewEvent(string(community.OnlineUserChanged))),
			guards.NewOnlineUserNumberGuard(guards.GreaterAndEqual(OnlineUserGuardNumber)),
			NewDefaultConversation(bot),
			NewInteracting(bot),
			[]cores.Action{}),
		transitions.NewBaseTransition(triggers.NewBaseTrigger(cores.NewEvent(string(community.OnlineUserChanged))),
			guards.NewOnlineUserNumberGuard(guards.Less(OnlineUserGuardNumber)),
			NewInteracting(bot),
			NewDefaultConversation(bot),
			[]cores.Action{}),
	})

	state := states.NewSubStateMachine(baseFsm, "Normal")

	normal := &Normal{
		bot: bot,
	}

	normal.State = state
	return normal
}

var _ fsm.State = &Normal{}
