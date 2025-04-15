package entities

import (
	"github/elliot9/yt/core"
	"github/elliot9/yt/infra"
)

type User struct {
	Name          string
	Subscriptions []*HumanSubscriber
	CLI           infra.CLI
}

func (u *User) Subscribe(channels []*core.Channel, onVideoNotification func(video *core.Video)) {
	for _, channel := range channels {
		subscriber := NewHumanSubscriber(u, u.CLI, onVideoNotification)
		subscriber.Channel = channel
		channel.Subscribed(subscriber)
		u.Subscriptions = append(u.Subscriptions, subscriber)
	}
}

func (u *User) Unsubscribe(channel *core.Channel) {
	for i, subscriber := range u.Subscriptions {
		if subscriber.Channel == channel {
			channel.Unsubscribed(subscriber)
			u.Subscriptions = append(u.Subscriptions[:i], u.Subscriptions[i+1:]...)
			break
		}
	}
}

func NewUser(name string, cli infra.CLI) *User {
	return &User{Name: name, Subscriptions: make([]*HumanSubscriber, 0), CLI: cli}
}
