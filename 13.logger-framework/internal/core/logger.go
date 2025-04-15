package core

type Logger struct {
	Name           string
	levelThreshold *LevelThreshold
	exporter       Exporter
	layout         LayoutStrategy
	parent         *Logger
}

type LoggerOption func(*Logger)

func NewLogger(name string, parent *Logger, options ...LoggerOption) *Logger {
	logger := &Logger{
		Name:           name,
		levelThreshold: nil,
		exporter:       nil,
		layout:         nil,
		parent:         parent,
	}

	for _, option := range options {
		option(logger)
	}

	// root logger
	if parent == nil {
		logger.Name = "Root"
	} else {
		if logger.levelThreshold == nil {
			logger.levelThreshold = parent.levelThreshold
		}

		if logger.exporter == nil {
			logger.exporter = parent.exporter
		}

		if logger.layout == nil {
			logger.layout = parent.layout
		}
	}

	return logger
}

func WithLevelThreshold(level LevelThreshold) LoggerOption {
	return func(l *Logger) {
		l.levelThreshold = &level
	}
}

func WithExporter(exporter Exporter) LoggerOption {
	return func(l *Logger) {
		l.exporter = exporter
	}
}

func WithLayout(layout LayoutStrategy) LoggerOption {
	return func(l *Logger) {
		l.layout = layout
	}
}

func (l *Logger) Trace(message string) {
	l.log(Trace, message)
}

func (l *Logger) Info(message string) {
	l.log(Info, message)
}

func (l *Logger) Debug(message string) {
	l.log(Debug, message)
}

func (l *Logger) Warn(message string) {
	l.log(Warn, message)
}

func (l *Logger) Error(message string) {
	l.log(Error, message)
}

func (l *Logger) log(level LevelThreshold, message string) {
	if level < *l.levelThreshold {
		return
	}

	formattedMessage := l.layout.Format(message, l.Name, level)
	l.exporter.Export(formattedMessage)
}
