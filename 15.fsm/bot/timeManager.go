package bot

import (
	"time"
)

type TimeManager interface {
	Now() time.Time
	AfterFunc(name string, d time.Duration, f func()) *time.Timer
	Advance(d time.Duration)
	CancelTimer(name string)
}

type MockTimeManager struct {
	currentTime time.Time
	timers      map[string]*MockTimer
}

type MockTimer struct {
	duration time.Duration
	callback func()
}

func NewMockTimeManager(startTime time.Time) *MockTimeManager {
	return &MockTimeManager{
		currentTime: startTime,
		timers:      make(map[string]*MockTimer),
	}
}

func (m *MockTimeManager) Now() time.Time {
	return m.currentTime
}

func (m *MockTimeManager) AfterFunc(name string, d time.Duration, f func()) *time.Timer {
	timer := &MockTimer{duration: d, callback: f}
	m.timers[name] = timer
	return nil
}

func (m *MockTimeManager) Advance(d time.Duration) {
	m.currentTime = m.currentTime.Add(d)
	for name, timer := range m.timers {
		if timer.duration <= d {
			timer.callback()
			delete(m.timers, name)
		}
	}
}

func (m *MockTimeManager) CancelTimer(name string) {
	delete(m.timers, name)
}

var _ TimeManager = &MockTimeManager{}
