package bot

import (
	"fmt"
	"github/elliot9/class15/community"
	"github/elliot9/class15/infra"
	"github/elliot9/class15/pkg/fsm/cores"
	"strings"
)

type Bot struct {
	quota        int
	fsm          *cores.FiniteStateMachine
	communityApi community.CommunityApi
	user         *community.User
	cli          *infra.MockCLI
	timeManager  TimeManager
}

func NewBot(communityApi community.CommunityApi, quota int, cli *infra.MockCLI, timeManager TimeManager) *Bot {
	bot := &Bot{
		communityApi: communityApi,
		quota:        quota,
		fsm:          cores.NewFiniteStateMachine(nil, []cores.Transition{}),
		user:         community.NewUser("bot", true),
		cli:          cli,
		timeManager:  timeManager,
	}
	communityApi.AddListener(bot)
	return bot
}

func (b *Bot) GetCommunity() community.CommunityApi {
	return b.communityApi
}

func (b *Bot) GetQuota() int {
	return b.quota
}

func (b *Bot) SetQuota(quota int) {
	b.quota = quota
}

func (b *Bot) GetFsm() *cores.FiniteStateMachine {
	return b.fsm
}

func (b *Bot) GetUser() *community.User {
	return b.user
}

func (b *Bot) GetTimeManager() TimeManager {
	return b.timeManager
}

func (b *Bot) OnEvent(event community.Event[any]) {
	ctx := b.createContext(event)
	b.fsm.SendEvent(cores.NewEvent(string(event.GetType())), ctx)
}

// output
func (b *Bot) ReplyMessage(userID, message string) {
	b.cli.Println(fmt.Sprintf("🤖: %s @%s", message, userID))
}

func (b *Bot) ReplyComment(postID, message string, tags []string) {
	for i, t := range tags {
		tags[i] = fmt.Sprintf("@%s", t)
	}

	tagString := strings.Join(tags, ", ")
	b.cli.Println(fmt.Sprintf("🤖 comment in post %s: %s %s", postID, message, tagString))
}

func (b *Bot) RecordReplay(messages []string, userID string) {
	message := strings.Join(messages, "\n")
	message = fmt.Sprintf("🤖: [Record Replay] %s @%s", message, userID)

	lines := strings.Split(message, "\n")
	for _, line := range lines {
		b.cli.Println(line)
	}
}

func (b *Bot) Message(message string) {
	b.cli.Println(fmt.Sprintf("🤖: %s", message))
}

func (b *Bot) NewQuestion(index int, question string, options [4]string) {
	message := fmt.Sprintf("🤖: %d. %s\n", index, question)
	optionsStrings := []string{}

	for i, option := range options {
		optionsStrings = append(optionsStrings, fmt.Sprintf("%s) %s", string(rune(i+65)), option))
	}

	message += strings.Join(optionsStrings, "\n")
	lines := strings.Split(message, "\n")
	for _, line := range lines {
		b.cli.Println(line)
	}
}

func (b *Bot) GoBroadcasting() {
	b.cli.Println("🤖 go broadcasting...")
}

func (b *Bot) Speak(content string) {
	b.cli.Println(fmt.Sprintf("🤖 speaking: %s", content))
}

func (b *Bot) StopBroadcasting() {
	b.cli.Println("🤖 stop broadcasting...")
}

var _ community.Listener = &Bot{}
