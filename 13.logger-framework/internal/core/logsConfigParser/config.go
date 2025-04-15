package logsConfigParser

import (
	"encoding/json"
	"fmt"
	"github/elliot9/class13/internal/core"
	"github/elliot9/class13/internal/core/exporters"
	"github/elliot9/class13/internal/core/layouts"
)

// 無特殊要求, Json -> Logger
// 如果有要求, 可調整成 Json -> Config -> Logger
const (
	ParseRootNode         = "loggers"
	RootLoggerName        = "Root"
	ParseLevelThreshold   = "levelThreshold"
	ParseExporter         = "exporter"
	ParseLayout           = "layout"
	ParseExporterType     = "type"
	ParseExporterFileName = "fileName"
	ParseExporterChildren = "children"
)

var loggers []*core.Logger

func Parse(jsonStr string) ([]*core.Logger, error) {
	loggers = []*core.Logger{}

	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return nil, err
	}

	loggersRaw, ok := raw[ParseRootNode].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ParseRootNode 格式錯誤")
	}

	parseLoggers(RootLoggerName, loggersRaw, nil)

	return loggers, nil
}

func parseLoggers(name string, data map[string]interface{}, parent *core.Logger) {
	options := []core.LoggerOption{}
	for key, value := range data {
		if key == ParseLevelThreshold || key == ParseExporter || key == ParseLayout {
			switch key {
			case ParseLevelThreshold:
				options = append(options, parseLevelThreshold(value.(string)))
			case ParseExporter:
				options = append(options, core.WithExporter(parseExporter(value)))
			case ParseLayout:
				options = append(options, parseLayout(value.(string)))
			}

			delete(data, key)
		}
	}

	currentLogger := core.NewLogger(name, parent, options...)
	loggers = append(loggers, currentLogger)

	if len(data) > 0 {
		for key, value := range data {
			parseLoggers(key, value.(map[string]interface{}), currentLogger)
		}
	}
}

func parseLevelThreshold(levelThreshold string) core.LoggerOption {
	return core.WithLevelThreshold(core.StringToLevelThreshold(levelThreshold))
}

func parseLayout(layout string) core.LoggerOption {
	switch layout {
	case "standard":
		return core.WithLayout(layouts.NewStandardLayout())
	}
	return nil
}

func parseExporter(data interface{}) core.Exporter {
	if data == nil {
		return nil
	}

	exporterMap, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}

	switch exporterMap[ParseExporterType] {
	case "console":
		return exporters.NewConsoleExporter()
	case "file":
		return exporters.NewFileExporter(exporterMap[ParseExporterFileName].(string))
	case "composite":
		childrenArray, ok := exporterMap[ParseExporterChildren].([]interface{})
		if ok {
			childExporters := []core.Exporter{}
			for _, child := range childrenArray {
				childExporter := parseExporter(child)
				if childExporter != nil {
					childExporters = append(childExporters, childExporter)
				}
			}
			return exporters.NewCompositeExporter(childExporters)
		}
	}

	return nil
}
