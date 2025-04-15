package main

import (
	"fmt"
	"github/elliot9/class13/internal/core"
	"github/elliot9/class13/internal/core/exporters"
	"github/elliot9/class13/internal/core/layouts"
	"github/elliot9/class13/pkg"
)

var Logs *pkg.Logs

func init() {
	Logs = pkg.NewLogs()
}

func main() {
	Logs.LoadFromJson("./pkg/config.json")
	mockGaming()
}

func _() {
	// 需求 AB
	rootLogger := core.NewLogger("Root",
		nil,
		core.WithLevelThreshold(core.Trace),
		core.WithExporter(exporters.NewConsoleExporter()),
		core.WithLayout(layouts.NewStandardLayout()),
	)

	gameLogger := core.NewLogger("app.game",
		rootLogger,
		core.WithLevelThreshold(core.Info),
		core.WithExporter(exporters.NewCompositeExporter([]core.Exporter{
			exporters.NewConsoleExporter(),
			exporters.NewCompositeExporter([]core.Exporter{
				exporters.NewFileExporter("game.log"),
				exporters.NewFileExporter("game.backup.log"),
			}),
		})),
	)

	aiLogger := core.NewLogger("app.game.ai",
		gameLogger,
		core.WithLevelThreshold(core.Trace),
		core.WithLayout(layouts.NewStandardLayout()),
	)

	Logs.DeclareLoggers(
		rootLogger,
		gameLogger,
		aiLogger,
	)
}

func mockGaming() {
	gameLogger := Logs.GetLogger("app.game")
	aiLogger := Logs.GetLogger("app.game.ai")

	players := []string{"AI 1", "AI 2", "AI 3", "AI 4"}

	gameLogger.Info("The game begins.")
	for _, player := range players {
		gameLogger.Trace(fmt.Sprintf("The player %s begins his turn.", player))

		aiLogger.Trace(fmt.Sprintf("%s starts making decisions...", player))
		aiLogger.Warn(fmt.Sprintf("%s decides to give up.", player))
		aiLogger.Error(fmt.Sprintf("Something goes wrong when %s gives up.", player))
		aiLogger.Trace(fmt.Sprintf("%s completes its decision.", player))

		gameLogger.Trace(fmt.Sprintf("The player %s finishes his turn.", player))
	}

	gameLogger.Debug("Game ends.")
}
