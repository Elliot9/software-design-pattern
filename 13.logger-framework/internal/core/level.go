package core

type LevelThreshold int

const (
	Trace LevelThreshold = iota
	Info
	Debug
	Warn
	Error
)

func (l LevelThreshold) String() string {
	return []string{"TRACE", "INFO", "DEBUG", "WARN", "ERROR"}[l]
}

func StringToLevelThreshold(s string) LevelThreshold {
	for i, v := range []string{"TRACE", "INFO", "DEBUG", "WARN", "ERROR"} {
		if v == s {
			return LevelThreshold(i)
		}
	}

	return Trace
}
