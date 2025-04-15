package core

type LayoutStrategy interface {
	Format(message, loggerName string, levelThreshold LevelThreshold) string
}
