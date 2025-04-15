package main

import (
	"github/elliot9/class19/application"
	"github/elliot9/class19/infra"
)

func main() {
	app := application.NewVocabularyLearningSystem(
		infra.NewLocalJsonWordFinder("data/words.json"),
		infra.NewMockWordRepository(),
		infra.NewRealRandomizer(),
		infra.NewConsoleIO(),
	)
	app.Run()
}
