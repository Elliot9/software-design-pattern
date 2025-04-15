package pkg

import (
	"github/elliot9/class13/internal/core"
	"github/elliot9/class13/internal/core/logsConfigParser"
	"os"
)

type Logs struct {
	loggers map[string]*core.Logger
}

func NewLogs() *Logs {
	return &Logs{
		loggers: make(map[string]*core.Logger),
	}
}

func (l *Logs) GetLogger(name string) *core.Logger {
	return l.loggers[name]
}

func (l *Logs) DeclareLoggers(loggers ...*core.Logger) {
	for _, logger := range loggers {
		l.loggers[logger.Name] = logger
	}
}

func (l *Logs) LoadFromJson(fileName string) {
	jsonStr, _ := os.ReadFile(fileName)
	loggers, err := logsConfigParser.Parse(string(jsonStr))
	if err != nil {
		panic(err)
	}

	l.DeclareLoggers(loggers...)
}
