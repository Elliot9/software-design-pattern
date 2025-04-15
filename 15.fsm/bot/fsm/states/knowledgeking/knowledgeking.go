package knowledgeking

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/bot/fsm/guards"
	"github/elliot9/class15/bot/fsm/states/normal"
	trigger "github/elliot9/class15/bot/fsm/triggers"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/actions"
	"github/elliot9/class15/pkg/fsm/plugins/states"
	"github/elliot9/class15/pkg/fsm/plugins/transitions"
	"github/elliot9/class15/pkg/fsm/plugins/triggers"
)

const (
	FinishQuestioning      string = "FinishQuestioning"
	FinishThanksForJoining string = "FinishThanksForJoining"
)

type KnowledgeKing struct {
	fsm.State
	bot           *bot.Bot
	quotaRequired int
}

func NewKnowledgeKing(bot *bot.Bot) *KnowledgeKing {
	baseFsm := fsm.NewFiniteStateMachine(func() cores.State {
		return NewQuestioning(bot)
	}, []cores.Transition{
		transitions.NewBaseTransition(triggers.NewBaseTrigger(cores.NewEvent(FinishQuestioning)),
			nil,
			NewQuestioning(bot),
			NewThanksForJoining(bot),
			[]cores.Action{}),
		transitions.NewBaseTransition(triggers.NewBaseTrigger(cores.NewEvent(FinishThanksForJoining)),
			nil,
			NewThanksForJoining(bot),
			normal.NewNormal(bot),
			[]cores.Action{}),
		transitions.NewBaseTransition(trigger.NewCommand("play again"),
			guards.NewQuota(5),
			NewThanksForJoining(bot),
			NewQuestioning(bot),
			[]cores.Action{actions.NewBaseAction(func(ctx context.Context) context.Context {
				bot.Message("KnowledgeKing is gonna start again!")
				return ctx
			})},
		),
	})

	state := states.NewSubStateMachine(baseFsm, "KnowledgeKing")

	knowledgeKing := &KnowledgeKing{
		bot:           bot,
		quotaRequired: 5,
	}

	knowledgeKing.State = state
	return knowledgeKing
}

func (r *KnowledgeKing) OnEntry(context context.Context) context.Context {
	r.bot.SetQuota(r.bot.GetQuota() - r.quotaRequired)
	r.bot.Message("KnowledgeKing is started!")
	return r.State.OnEntry(context)
}

var _ fsm.State = &KnowledgeKing{}
