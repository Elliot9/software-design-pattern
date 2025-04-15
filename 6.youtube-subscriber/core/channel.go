package core

import (
	"fmt"
	"github/elliot9/yt/infra"
)

type Channel struct {
	Cli         infra.CLI
	Name        string
	Videos      []*Video
	VideoChan   chan *Video  // mock the stream of video upload
	Subscribers []Subscriber // can be more effective if we use a map and use the id as the key
}

func NewChannel(cli infra.CLI, name string) *Channel {
	channel := &Channel{
		Cli:         cli,
		Name:        name,
		Subscribers: make([]Subscriber, 0),
		VideoChan:   make(chan *Video),
		Videos:      make([]*Video, 0),
	}

	go func() {
		for video := range channel.VideoChan {
			channel.Videos = append(channel.Videos, video)
			channel.notify()
		}
	}()

	return channel
}

func (c *Channel) notify() {
	latestVideo := c.Videos[len(c.Videos)-1]
	c.Cli.Println(fmt.Sprintf("頻道 %s 上架了一則新影片 \"%s\"。", c.Name, latestVideo.Title))

	for _, subscriber := range c.Subscribers {
		subscriber.ReactToNotification(latestVideo)
	}
}

func (c *Channel) Upload(video *Video) {
	c.VideoChan <- video
}

func (c *Channel) Unsubscribed(subscriber Subscriber) {
	subscribers := make([]Subscriber, 0, len(c.Subscribers))

	for _, s := range c.Subscribers {
		if s != subscriber {
			subscribers = append(subscribers, s)
		}
	}

	c.Subscribers = subscribers
	c.Cli.Println(fmt.Sprintf("%s 解除訂閱了 %s。", subscriber.GetName(), c.Name))
}

func (c *Channel) Subscribed(subscriber Subscriber) {
	c.Subscribers = append(c.Subscribers, subscriber)
	c.Cli.Println(fmt.Sprintf("%s 訂閱了 %s。", subscriber.GetName(), c.Name))
}
