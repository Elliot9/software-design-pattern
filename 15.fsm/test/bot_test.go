package test

import (
	"encoding/json"
	"fmt"
	"github/elliot9/class15/bootstrap"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	"github/elliot9/class15/infra"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestChatBotDefaultInitialMessageResponse(t *testing.T) {
	test(t, "data/ChatBot_Default_InitialMessageResponse.in", "data/ChatBot_Default_InitialMessageResponse.out")
}

func TestChatBotDefaultMessageCycle(t *testing.T) {
	test(t, "data/ChatBot_Default_MessageCycle.in", "data/ChatBot_Default_MessageCycle.out")
}

func TestChatBotDefaultUserTagsBot(t *testing.T) {
	test(t, "data/ChatBot_Default_UserTagsBot.in", "data/ChatBot_Default_UserTagsBot.out")
}

func TestChatBotDefaultMultipleUsers(t *testing.T) {
	test(t, "data/ChatBot_Default_MultipleUsers.in", "data/ChatBot_Default_MultipleUsers.out")
}

func TestChatBotInteractiveInitialMessageResponse(t *testing.T) {
	test(t, "data/ChatBot_Interactive_InitialMessageResponse.in", "data/ChatBot_Interactive_InitialMessageResponse.out")
}

func TestChatBotInteractiveMessageCycle(t *testing.T) {
	test(t, "data/ChatBot_Interactive_MessageCycle.in", "data/ChatBot_Interactive_MessageCycle.out")
}

func TestChatBotInteractiveLogoutSwitchToDefault(t *testing.T) {
	test(t, "data/ChatBot_Interactive_LogoutSwitchToDefault.in", "data/ChatBot_Interactive_LogoutSwitchToDefault.out")
}

func TestForumBotCommentDefaultState(t *testing.T) {
	test(t, "data/Forum_BotComment_DefaultState.in", "data/Forum_BotComment_DefaultState.out")
}

func TestForumBotCommentInteractiveState(t *testing.T) {
	test(t, "data/Forum_BotComment_InteractiveState.in", "data/Forum_BotComment_InteractiveState.out")
}

func TestForumBotComment10Users(t *testing.T) {
	test(t, "data/Forum_BotComment_10Users.in", "data/Forum_BotComment_10Users.out")
}

func TestBroadcastSingleUser(t *testing.T) {
	test(t, "data/Broadcast_SingleUser.in", "data/Broadcast_SingleUser.out")
}

func TestBroadcastMultipleSpeakers(t *testing.T) {
	test(t, "data/Broadcast_MultipleSpeakers.in", "data/Broadcast_MultipleSpeakers.out")
}

func TestBroadcastRecordActiveSpeaker(t *testing.T) {
	test(t, "data/Broadcast_Record_ActiveSpeaker.in", "data/Broadcast_Record_ActiveSpeaker.out")
}

func TestBroadcastRecordStopBroadcasting(t *testing.T) {
	test(t, "data/Broadcast_Record_StopBroadcasting.in", "data/Broadcast_Record_StopBroadcasting.out")
}

func TestCommandRecordQuotaLimit(t *testing.T) {
	test(t, "data/Command_Record_QuotaLimit.in", "data/Command_Record_QuotaLimit.out")
}

func TestCommandRecordStopByNonRecorder(t *testing.T) {
	test(t, "data/Command_Record_StopByNonRecorder.in", "data/Command_Record_StopByNonRecorder.out")
}

func TestCommandRecordStopByRecorder(t *testing.T) {
	test(t, "data/Command_Record_StopByRecorder.in", "data/Command_Record_StopByRecorder.out")
}

func TestKnowledgeKingStartByAdmin(t *testing.T) {
	test(t, "data/KnowledgeKing_StartByAdmin.in", "data/KnowledgeKing_StartByAdmin.out")
}

func TestKnowledgeKingStartByMember(t *testing.T) {
	test(t, "data/KnowledgeKing_StartByMember.in", "data/KnowledgeKing_StartByMember.out")
}

func TestKnowledgeKing_CorrectAnswerFirstUser(t *testing.T) {
	test(t, "data/KnowledgeKing_CorrectAnswer_FirstUser.in", "data/KnowledgeKing_CorrectAnswer_FirstUser.out")
}

func TestKnowledgeKingCorrectAnswerMultipleUsers(t *testing.T) {
	test(t, "data/KnowledgeKing_CorrectAnswer_MultipleUsers.in", "data/KnowledgeKing_CorrectAnswer_MultipleUsers.out")
}

func TestKnowledgeKingTieGame(t *testing.T) {
	test(t, "data/KnowledgeKing_TieGame.in", "data/KnowledgeKing_TieGame.out")
}

func TestKnowledgeKingIncorrectAnswer(t *testing.T) {
	test(t, "data/KnowledgeKing_IncorrectAnswer.in", "data/KnowledgeKing_IncorrectAnswer.out")
}

func TestKnowledgeKingTimeOut(t *testing.T) {
	test(t, "data/KnowledgeKing_TimeOut.in", "data/KnowledgeKing_TimeOut.out")
}

func TestKnowledgeKingWinnerAnnouncement(t *testing.T) {
	test(t, "data/KnowledgeKing_WinnerAnnouncement.in", "data/KnowledgeKing_WinnerAnnouncement.out")
}

func TestKnowledgeKingPlayAgain(t *testing.T) {
	test(t, "data/KnowledgeKing_PlayAgain.in", "data/KnowledgeKing_PlayAgain.out")
}

func TestCommandKingValidAdmin(t *testing.T) {
	test(t, "data/Command_King_ValidAdmin.in", "data/Command_King_ValidAdmin.out")
}

func TestCommandKingInvalidMember(t *testing.T) {
	test(t, "data/Command_King_InvalidMember.in", "data/Command_King_InvalidMember.out")
}

func TestCommandKingQuotaLimit(t *testing.T) {
	test(t, "data/Command_King_QuotaLimit.in", "data/Command_King_QuotaLimit.out")
}

var communityApi community.CommunityApi
var botInstance *bot.Bot

func test(t *testing.T, inputFile string, outputFile string) {
	inputContent, err := os.ReadFile(inputFile)
	cli := infra.NewMockCLI()
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	inputs := strings.Split(string(inputContent), "\n")

	expectedOutput, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	outputs := strings.Split(string(expectedOutput), "\n")

	communityApi = community.NewMockCommunity()
	userMap := make(map[string]*community.User)

outerLoop:
	for _, input := range inputs {
		eventName, eventData := parseEvent(input)

		switch eventName {
		case "started":
			startTime, _ := time.Parse("2006-01-02 15:04:05", eventData["time"].(string))
			quota := int(eventData["quota"].(float64))
			timeManager := bot.NewMockTimeManager(startTime)
			botInstance = bootstrap.Bootstrap(communityApi, timeManager, quota, cli)
		case "login":
			userMap[eventData["userId"].(string)] = community.NewUser(eventData["userId"].(string), eventData["isAdmin"].(bool))
			communityApi.Login(userMap[eventData["userId"].(string)])
		case "logout":
			communityApi.Logout(userMap[eventData["userId"].(string)])
			delete(userMap, eventData["userId"].(string))
		case "new message":
			tags := []string{}
			for _, tag := range eventData["tags"].([]interface{}) {
				tags = append(tags, tag.(string))
			}

			message := fmt.Sprintf("💬 %s: %s", eventData["authorId"].(string), eventData["content"].(string))
			if len(tags) > 0 {
				message = fmt.Sprintf("%s %s", message, parseTags(tags))
			}

			cli.Println(message)

			communityApi.NewMessage(userMap[eventData["authorId"].(string)], community.Message{
				Content: eventData["content"].(string),
				Tags:    tags,
			})

		case "new post":
			tags := []string{}
			for _, tag := range eventData["tags"].([]interface{}) {
				tags = append(tags, tag.(string))
			}

			message := fmt.Sprintf("%s: 【%s】%s", eventData["authorId"].(string), eventData["title"].(string), eventData["content"].(string))
			cli.Println(message)

			communityApi.NewPost(userMap[eventData["authorId"].(string)], community.Post{
				Id:      eventData["id"].(string),
				Title:   eventData["title"].(string),
				Content: eventData["content"].(string),
				Tags:    tags,
			})
		case "go broadcasting":
			message := fmt.Sprintf("📢 %s is broadcasting...", eventData["speakerId"].(string))
			cli.Println(message)
			communityApi.GoBroadcasting(userMap[eventData["speakerId"].(string)])
		case "stop broadcasting":
			message := fmt.Sprintf("📢 %s stop broadcasting", eventData["speakerId"].(string))
			cli.Println(message)
			communityApi.StopBroadcasting(userMap[eventData["speakerId"].(string)])
		case "speak":
			message := fmt.Sprintf("📢 %s: %s", eventData["speakerId"].(string), eventData["content"].(string))
			cli.Println(message)
			communityApi.Speak(userMap[eventData["speakerId"].(string)], community.SpeakContent{
				Content: eventData["content"].(string),
			})
		case "end":
			break outerLoop
		default:
			re := regexp.MustCompile(`^(\d+) (seconds|minutes|hours) elapsed$`)
			matches := re.FindStringSubmatch(eventName)
			if len(matches) > 1 {
				message := fmt.Sprintf("🕑 %s %s elapsed...", matches[1], matches[2])
				cli.Println(message)

				num, _ := strconv.Atoi(matches[1])
				var duration time.Duration
				switch matches[2] {
				case "seconds":
					duration = time.Duration(num) * time.Second
				case "minutes":
					duration = time.Duration(num) * time.Minute
				case "hours":
					duration = time.Duration(num) * time.Hour
				}

				botInstance.GetTimeManager().Advance(duration)
			}
		}
	}

	for i, _ := range outputs {
		if i != len(outputs)-1 {
			if strings.TrimSpace(cli.GetOutputs()[i]) != strings.TrimSpace(outputs[i]) {
				t.Fatalf("expected %s, got %s", outputs[i], cli.GetOutputs()[i])
			}
		}
	}
}

func parseEvent(input string) (string, map[string]interface{}) {
	re := regexp.MustCompile(`^\[(.*)]\s*(\{(.*)})?$`)
	matches := re.FindStringSubmatch(input)
	var eventData map[string]interface{}
	json.Unmarshal([]byte(matches[2]), &eventData)
	return matches[1], eventData
}

func parseTags(tags []string) string {
	formattedTags := []string{}
	for _, tag := range tags {
		formattedTags = append(formattedTags, fmt.Sprintf("@%s", tag))
	}
	return strings.Join(formattedTags, ", ")
}
