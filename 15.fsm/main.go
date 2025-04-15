package main

import (
	"fmt"
	"github/elliot9/class15/bootstrap"
	"github/elliot9/class15/bot"
	"github/elliot9/class15/community"
	"github/elliot9/class15/infra"
	"time"
)

func main() {
	// go bot_test.go file

	// arrange
	communityApi := community.NewMockCommunity()
	timeManager := bot.NewMockTimeManager(time.Now())
	cli := infra.NewMockCLI()
	bootstrap.Bootstrap(communityApi, timeManager, 100, cli)
	user := community.NewUser("elliot", true)

	// act
	communityApi.Login(user)
	communityApi.NewMessage(user, community.Message{
		Content: "今天天氣真好",
		Tags:    []string{"bot"},
	})

	// assert
	fmt.Println(cli.GetOutputs())
}
