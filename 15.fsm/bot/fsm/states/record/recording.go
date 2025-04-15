package record

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/stateListeners"
	"github/elliot9/class15/pkg/fsm/plugins/states"
)

type Recording struct {
	fsm.State
	messages []string
	bot      *bot.Bot
}

func NewRecording(botInstance *bot.Bot) *Recording {
	state := states.NewBaseState("Recording")

	recording := &Recording{
		bot:      botInstance,
		messages: []string{},
	}

	state.AddListeners(map[fsm.Event][]fsm.StateListener{
		cores.NewEvent(string(community.NewSpeak)): {
			stateListeners.NewBaseStateListener(func(context context.Context) {
				speakContent := context.Value(bot.ContextSpeakContent).(community.SpeakContent)
				recording.appendMessage(speakContent.Content)
			}),
		},
	})

	recording.State = state
	return recording
}

func (r *Recording) appendMessage(message string) {
	r.messages = append(r.messages, message)
}

func (r *Recording) OnExit(context context.Context) context.Context {
	r.bot.RecordReplay(r.messages, Recorder.GetId())
	Recorder = nil
	return context
}

var _ fsm.State = &Recording{}
