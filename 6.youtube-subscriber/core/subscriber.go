package core

type Subscriber interface {
	ReactToNotification(video *Video)
	GetName() string
}

type ChannelSubscriber struct {
	Channel *Channel
	Name    string
}

func (c *ChannelSubscriber) ReactToNotification(video *Video) {
	panic("ReactToNotification not implemented")
}

func (c *ChannelSubscriber) GetName() string {
	return c.Name
}

var _ Subscriber = &ChannelSubscriber{}
