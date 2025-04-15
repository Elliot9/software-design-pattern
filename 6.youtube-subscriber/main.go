package main

import (
	"fmt"
	"github/elliot9/yt/core"
	"github/elliot9/yt/core/entities"
	"github/elliot9/yt/infra"
	"time"
)

func main() {
	cli := infra.NewMockCLI()

	// given
	waterSchool := core.NewChannel(cli, "水球軟體學院")
	pewDiePie := core.NewChannel(cli, "PewDiePie")
	water := entities.NewUser("水球", cli)
	fire := entities.NewUser("火球", cli)

	// when
	//水球訂閱 PewDiePie 和 水球軟體學院
	//火球訂閱 PewDiePie 和 水球軟體學院
	water.Subscribe([]*core.Channel{waterSchool, pewDiePie}, mockWaterOnVideoNotification(water, cli))
	fire.Subscribe([]*core.Channel{waterSchool, pewDiePie}, mockFireOnVideoNotification(fire, cli))

	//水球軟體學院上傳一部影片：標題：”C1M1S2”、敘述：”這個世界正是物件導向的呢！”、影片長度：4 分鐘。
	waterSchool.Upload(core.NewVideo("C1M1S2", "這個世界正是物件導向的呢！", 4*time.Minute))
	time.Sleep(1 * time.Second)

	//PewDiePie 上傳一部影片：標題：”Hello guys”、敘述：”Clickbait”、影片長度：30 秒。
	pewDiePie.Upload(core.NewVideo("Hello guys", "Clickbait", 30*time.Second))
	time.Sleep(1 * time.Second)

	// 水球軟體學院上傳一部影片：標題：”C1M1S3”、敘述：”物件 vs. 類別”、影片長度：1 分鐘。
	waterSchool.Upload(core.NewVideo("C1M1S3", "物件 vs. 類別", 1*time.Minute))
	time.Sleep(1 * time.Second)

	// PewDiePie 上傳一部影片：標題：”Minecraft”、敘述：”Let’s play Minecraft”、影片長度：30 分鐘。
	pewDiePie.Upload(core.NewVideo("Minecraft", "Let’s play Minecraft", 30*time.Minute))
	time.Sleep(1 * time.Second)

	// then
	if !isPass(cli) {
		fmt.Println("測試失敗")
	} else {
		fmt.Println("測試成功")
	}
}

func isPass(cli infra.CLI) bool {
	expected := []string{
		"水球 訂閱了 水球軟體學院。",
		"水球 訂閱了 PewDiePie。",
		"火球 訂閱了 水球軟體學院。",
		"火球 訂閱了 PewDiePie。",
		"頻道 水球軟體學院 上架了一則新影片 \"C1M1S2\"。",
		"水球 對影片 \"C1M1S2\" 按讚。",
		"頻道 PewDiePie 上架了一則新影片 \"Hello guys\"。",
		"火球 解除訂閱了 PewDiePie。",
		"頻道 水球軟體學院 上架了一則新影片 \"C1M1S3\"。",
		"火球 解除訂閱了 水球軟體學院。",
		"頻道 PewDiePie 上架了一則新影片 \"Minecraft\"。",
		"水球 對影片 \"Minecraft\" 按讚。",
	}

	for i, line := range cli.GetOutput() {
		if line != expected[i] {
			return false
		}
	}

	return true
}

func mockWaterOnVideoNotification(user *entities.User, cli infra.CLI) func(video *core.Video) {
	return func(video *core.Video) {
		if video.Length >= 3*time.Minute {
			cli.Println(fmt.Sprintf("%s 對影片 \"%s\" 按讚。", user.Name, video.Title))
		}
	}
}

func mockFireOnVideoNotification(user *entities.User, _ infra.CLI) func(video *core.Video) {
	return func(video *core.Video) {
		if video.Length <= 1*time.Minute {
			for _, subscriber := range user.Subscriptions {
				if len(subscriber.Channel.Videos) > 0 && subscriber.Channel.Videos[len(subscriber.Channel.Videos)-1] == video {
					user.Unsubscribe(subscriber.Channel)
					return
				}
			}
		}
	}
}
