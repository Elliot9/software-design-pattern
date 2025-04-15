package knowledgeking

import (
	"context"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	"github/elliot9/class15/pkg/fsm/cores"
	fsm "github/elliot9/class15/pkg/fsm/cores"
	"github/elliot9/class15/pkg/fsm/plugins/stateListeners"
	"github/elliot9/class15/pkg/fsm/plugins/states"
	"time"
)

const (
	QuestioningTimer = "QuestioningTimer"
)

type Question struct {
	Title   string
	Options [4]string
	Answer  string
}

type Questioning struct {
	cores.State
	bot       *bot.Bot
	quetions  []Question
	pivot     int
	userScore map[string]int
}

func NewQuestioning(bot *bot.Bot) *Questioning {
	state := states.NewBaseState("Questioning")

	questioning := &Questioning{
		bot: bot,
		quetions: []Question{
			{
				Title:   "請問哪個 SQL 語句用於選擇所有的行？",
				Options: [4]string{"SELECT *", "SELECT ALL", "SELECT ROWS", "SELECT DATA"},
				Answer:  "A",
			},
			{
				Title:   "請問哪個 CSS 屬性可用於設置文字的顏色？",
				Options: [4]string{"text-align", "font-size", "color", "padding"},
				Answer:  "C",
			},
			{
				Title:   "請問在計算機科學中，「XML」代表什麼？",
				Options: [4]string{"Extensible Markup Language", "Extensible Modeling Language", "Extended Markup Language", "Extended Modeling Language"},
				Answer:  "A",
			},
		},
		pivot:     0,
		userScore: map[string]int{},
	}

	state.AddListeners(map[fsm.Event][]fsm.StateListener{
		cores.NewEvent(string(community.NewMessage)): {NewQuestionReplyListener(questioning)},
	})

	questioning.State = state
	return questioning
}

func (q *Questioning) OnEntry(ctx context.Context) context.Context {
	q.CurrentQuestion()

	// 若在 1 小時之後，這 3 題尚未全部答完，那麼也會立即中斷且進入感謝參與
	q.bot.GetTimeManager().AfterFunc(QuestioningTimer, 1*time.Hour, func() {
		q.bot.GetFsm().SendEvent(cores.NewEvent(FinishQuestioning), ctx)
	})

	return ctx
}

func (q *Questioning) OnExit(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, bot.ContextWinner, q.GetCurrentWinner())
	q.bot.GetTimeManager().CancelTimer(QuestioningTimer)
	return ctx
}

func (q *Questioning) CurrentQuestion() {
	q.bot.NewQuestion(q.pivot, q.quetions[q.pivot].Title, q.quetions[q.pivot].Options)
}

func (q *Questioning) GetCurrentWinner() string {
	highestScore := 0
	winner := ""
	for user, score := range q.userScore {
		if score > highestScore {
			highestScore = score
			winner = user
		} else if score == highestScore {
			winner = ""
		}
	}

	return winner
}

func (q *Questioning) AnswerQuestion(user *community.User, answer string, ctx context.Context) {
	if answer != q.quetions[q.pivot].Answer {
		return
	}

	q.pivot++
	q.userScore[user.GetId()]++
	q.bot.ReplyMessage(user.GetId(), "Congrats! you got the answer!")
	if q.pivot < len(q.quetions) {
		q.CurrentQuestion()
	} else {
		// 進入感謝環節
		q.bot.GetFsm().SendEvent(cores.NewEvent(FinishQuestioning), ctx)
	}
}

var _ fsm.State = &Questioning{}

type QuestionReplyListener struct {
	fsm.StateListener
	questioning *Questioning
}

func NewQuestionReplyListener(questioning *Questioning) *QuestionReplyListener {
	listener := &QuestionReplyListener{
		questioning: questioning,
	}

	listener.StateListener = stateListeners.NewBaseStateListener(func(context context.Context) {
		botInstance := context.Value(bot.ContextBot).(*bot.Bot)
		user := context.Value(bot.ContextUser).(*community.User)
		message := context.Value(bot.ContextMessage).(community.Message)

		if len(message.Tags) != 1 || message.Tags[0] != botInstance.GetUser().GetId() {
			return
		}

		answer := message.Content
		questioning.AnswerQuestion(user, answer, context)
	})
	return listener
}

var _ fsm.StateListener = &QuestionReplyListener{}
