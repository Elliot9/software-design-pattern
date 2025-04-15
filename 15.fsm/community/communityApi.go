package community

import (
	"sort"
	"strconv"
)

type CommunityApi interface {
	NewMessage(author *User, message Message)
	NewPost(author *User, post Post)
	NewComment(replier *User, comment Comment)
	GoBroadcasting(user *User)
	StopBroadcasting(user *User)
	Speak(speaker *User, speak SpeakContent)
	Login(user *User)
	Logout(user *User)
	GetOnlineUsers() []*User
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	GetRecorder() *User
}

type MockCommunityApi struct {
	Users     map[string]*User
	Listeners []Listener
	Recorder  *User
}

func NewMockCommunity() *MockCommunityApi {
	return &MockCommunityApi{
		Users:     make(map[string]*User),
		Listeners: make([]Listener, 0),
		Recorder:  nil,
	}
}

type EventDataKey string

const (
	EventDataUser    EventDataKey = "user"
	EventDataMessage EventDataKey = "message"
	EventDataPost    EventDataKey = "post"
	EventDataComment EventDataKey = "comment"
	EventDataSpeak   EventDataKey = "speak"
)

func (c *MockCommunityApi) AddListener(listener Listener) {
	c.Listeners = append(c.Listeners, listener)
}

func (c *MockCommunityApi) RemoveListener(listener Listener) {
	for i, l := range c.Listeners {
		if l == listener {
			c.Listeners = append(c.Listeners[:i], c.Listeners[i+1:]...)
			break
		}
	}
}

func (c *MockCommunityApi) NewMessage(author *User, message Message) {
	eventData := map[EventDataKey]interface{}{
		EventDataUser:    author,
		EventDataMessage: message,
	}
	c.sendEvent(NewMessage, eventData)
}

func (c *MockCommunityApi) NewPost(author *User, post Post) {
	eventData := map[EventDataKey]interface{}{
		EventDataUser: author,
		EventDataPost: post,
	}
	c.sendEvent(NewPost, eventData)
}

func (c *MockCommunityApi) NewComment(replier *User, comment Comment) {
	eventData := map[EventDataKey]interface{}{
		EventDataUser:    replier,
		EventDataComment: comment,
	}
	c.sendEvent(NewComment, eventData)
}

func (c *MockCommunityApi) GoBroadcasting(user *User) {
	if c.Recorder != nil {
		return
	}

	c.Recorder = user
	eventData := map[EventDataKey]interface{}{
		EventDataUser: user,
	}
	c.sendEvent(GoBroadcasting, eventData)
}

func (c *MockCommunityApi) StopBroadcasting(user *User) {
	if user.GetId() != c.Recorder.GetId() {
		return
	}

	eventData := map[EventDataKey]interface{}{
		EventDataUser: user,
	}
	c.sendEvent(StopBroadcasting, eventData)
}

func (c *MockCommunityApi) Speak(speaker *User, speak SpeakContent) {
	if c.Recorder == nil {
		return
	}

	eventData := map[EventDataKey]interface{}{
		EventDataUser:  speaker,
		EventDataSpeak: speak,
	}
	c.sendEvent(NewSpeak, eventData)
}

func (c *MockCommunityApi) Login(user *User) {
	c.Users[user.GetId()] = user
	c.sendEvent(OnlineUserChanged, nil)
}

func (c *MockCommunityApi) Logout(user *User) {
	delete(c.Users, user.GetId())
	c.sendEvent(OnlineUserChanged, nil)
}

func (c *MockCommunityApi) GetOnlineUsers() []*User {
	users := make([]*User, 0, len(c.Users))
	for _, user := range c.Users {
		users = append(users, user)
	}
	sort.Slice(users, func(i, j int) bool {
		userI, _ := strconv.Atoi(users[i].GetId())
		userJ, _ := strconv.Atoi(users[j].GetId())
		return userI < userJ
	})
	return users
}

func (c *MockCommunityApi) GetRecorder() *User {
	return c.Recorder
}

func (c *MockCommunityApi) sendEvent(eventType EventType, data any) {
	for _, l := range c.Listeners {
		l.OnEvent(NewEvent(eventType, data))
	}
}

var _ CommunityApi = &MockCommunityApi{}
