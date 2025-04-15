package config

import (
	"context"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
)

func GetDeleteWordsSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &DeleteWordsHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_DistanceBased,
		LifeCycleDistance: 2,
		UIFactory:         &DeleteWordsUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type DeleteWordsUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *DeleteWordsHandlerFactory
}

func (f *DeleteWordsUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	return scenes.NewNavigationBreadcrumb(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager))
}

func (f *DeleteWordsUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	return scenes.NewCurrentWordsMessenger(ctx.Value(scenes.ContextWordService).(*word.WordService))
}

func (f *DeleteWordsUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *DeleteWordsUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("Please input word's names (separated by commas) that you want to delete:")
}

type DeleteWordsHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *DeleteWordsHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	commands = append(commands, scenes.NewBackCommand())
	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *DeleteWordsHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return func(ctx context.Context, input string) {
		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)

		wordsDeleted, wordsNotFound := []word.Word{}, []string{}

		for _, key := range parseCommaWords(input) {
			word, found := wordService.Search(key)
			if !found {
				wordsNotFound = append(wordsNotFound, key)
			} else {
				wordsDeleted = append(wordsDeleted, word)
			}
		}

		for i := len(wordsDeleted) - 1; i >= 0; i-- {
			wordService.DeleteWord(wordsDeleted[i])
		}
		currentWords := wordService.GetAllWords()
		scene := sceneManager.Current().(*scenes.VocabularyLearningScene)
		scene.SetMessenger(scenes.NewDeleteWordsMessenger(currentWords, wordsDeleted, wordsNotFound))
	}
}
