package layouts

import (
	"fmt"
	"time"

	"github/elliot9/class13/internal/core"
)

type StandardLayout struct{}

var _ core.LayoutStrategy = &StandardLayout{}

func (l *StandardLayout) Format(message, loggerName string, levelThreshold core.LevelThreshold) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	return fmt.Sprintf("%s |-%s %s - %s", timestamp, levelThreshold.String(), loggerName, message)
}

func NewStandardLayout() *StandardLayout {
	return &StandardLayout{}
}
