package bootstrap

import (
	"github/elliot9/class15/bot"
	"github/elliot9/class15/bot/fsm/guards"
	"github/elliot9/class15/bot/fsm/states/knowledgeking"
	"github/elliot9/class15/bot/fsm/states/normal"
	"github/elliot9/class15/bot/fsm/states/record"
	"github/elliot9/class15/bot/fsm/triggers"
	"github/elliot9/class15/community"
	"github/elliot9/class15/infra"
	"github/elliot9/class15/pkg/fsm/cores"
	baseGuard "github/elliot9/class15/pkg/fsm/plugins/guards"
	"github/elliot9/class15/pkg/fsm/plugins/transitions"
)

func Bootstrap(communityApi community.CommunityApi, timer bot.TimeManager, quota int, cli *infra.MockCLI) *bot.Bot {
	botInstance := bot.NewBot(communityApi, quota, cli, timer)
	setupTransitions(botInstance)
	setupInitialState(botInstance)
	return botInstance
}

func setupTransitions(botInstance *bot.Bot) {
	botInstance.GetFsm().AddTransitions([]cores.Transition{
		transitions.NewBaseTransition(triggers.NewCommand("record"),
			guards.NewQuota(3),
			normal.NewNormal(botInstance),
			record.NewRecord(botInstance),
			[]cores.Action{}),
		transitions.NewBaseTransition(triggers.NewCommand("stop-recording"),
			guards.NewRecorder(),
			record.NewRecord(botInstance),
			normal.NewNormal(botInstance),
			[]cores.Action{}),
		transitions.NewBaseTransition(triggers.NewCommand("king"),
			baseGuard.NewAndGuard([]cores.Guard{
				guards.NewQuota(5),
				guards.NewAdminOnly(),
			}),
			normal.NewNormal(botInstance),
			knowledgeking.NewKnowledgeKing(botInstance),
			[]cores.Action{}),
		transitions.NewBaseTransition(triggers.NewCommand("king-stop"),
			guards.NewAdminOnly(),
			knowledgeking.NewKnowledgeKing(botInstance),
			normal.NewNormal(botInstance),
			[]cores.Action{}),
	})
}

func setupInitialState(botInstance *bot.Bot) {
	botInstance.GetFsm().SetInitialStateResolver(func() cores.State {
		return normal.NewNormal(botInstance)
	})
}
