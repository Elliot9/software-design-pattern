package knowledgeking

import (
	"context"
	"fmt"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/states"
	"time"
)

const (
	ThanksForJoiningTimer = "ThanksForJoiningTimer"
)

type ThanksForJoining struct {
	cores.State
	bot *bot.Bot
}

func NewThanksForJoining(bot *bot.Bot) *ThanksForJoining {
	state := states.NewBaseState("ThanksForJoining")

	thanksForJoining := &ThanksForJoining{
		bot: bot,
	}

	thanksForJoining.State = state
	return thanksForJoining
}

func (t *ThanksForJoining) OnEntry(ctx context.Context) context.Context {
	messages := "Tie!"

	winner := ctx.Value(bot.ContextWinner).(string)
	if winner != "" {
		messages = fmt.Sprintf("The winner is %s", winner)
	}

	if t.bot.GetCommunity().GetRecorder() != nil {
		t.bot.Message(messages)
	} else {
		t.bot.GoBroadcasting()
		t.bot.Speak(messages)
		t.bot.StopBroadcasting()
	}

	t.bot.GetTimeManager().AfterFunc(ThanksForJoiningTimer, 20*time.Second, func() {
		t.bot.GetFsm().SendEvent(cores.NewEvent(FinishThanksForJoining), ctx)
	})

	return ctx
}

func (t *ThanksForJoining) OnExit(ctx context.Context) context.Context {
	t.bot.GetTimeManager().CancelTimer(ThanksForJoiningTimer)
	return ctx
}

var _ fsm.State = &ThanksForJoining{}
