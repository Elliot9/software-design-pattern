package entities

import (
	"github/elliot9/yt/core"
	"github/elliot9/yt/infra"
)

type HumanSubscriber struct {
	User                *User
	CLI                 infra.CLI
	OnVideoNotification func(video *core.Video)
	core.ChannelSubscriber
}

func (h *HumanSubscriber) ReactToNotification(video *core.Video) {
	h.OnVideoNotification(video)
}

func (h *HumanSubscriber) GetName() string {
	return h.User.Name
}

func NewHumanSubscriber(user *User, cli infra.CLI, onVideoNotification func(video *core.Video)) *HumanSubscriber {
	return &HumanSubscriber{User: user, CLI: cli, OnVideoNotification: onVideoNotification, ChannelSubscriber: core.ChannelSubscriber{Name: user.Name}}
}

var _ core.Subscriber = &HumanSubscriber{}
