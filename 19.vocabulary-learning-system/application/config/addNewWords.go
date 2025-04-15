package config

import (
	"context"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
	"strings"
)

func GetAddNewWordsSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &AddNewWordsHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_DistanceBased,
		LifeCycleDistance: 2,
		UIFactory:         &AddNewWordsUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type AddNewWordsUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *AddNewWordsHandlerFactory
}

func (f *AddNewWordsUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	return scenes.NewNavigationBreadcrumb(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager))
}

func (f *AddNewWordsUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	return scenes.NewCurrentWordsMessenger(ctx.Value(scenes.ContextWordService).(*word.WordService))
}

func (f *AddNewWordsUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *AddNewWordsUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("Please input word's names (separated by commas) that you want to add: ")
}

type AddNewWordsHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *AddNewWordsHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	commands = append(commands, scenes.NewBackCommand())
	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *AddNewWordsHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return func(ctx context.Context, input string) {
		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)

		wordsAdded, wordsNotFound := []word.Word{}, []string{}

		for _, key := range parseCommaWords(input) {
			word, found := wordService.Search(key)
			if !found {
				wordsNotFound = append(wordsNotFound, key)
			} else {
				wordsAdded = append(wordsAdded, word)
			}
		}

		for _, addWord := range wordsAdded {
			wordService.AddWord(addWord)
		}

		scene := sceneManager.Current().(*scenes.VocabularyLearningScene)
		scene.SetMessenger(scenes.NewAddNewWordsMessenger(wordService.GetAllWords(), wordsAdded, wordsNotFound))
	}
}

func parseCommaWords(input string) []string {
	words := strings.Split(input, ",")
	result := []string{}

	for _, word := range words {
		trimmed := strings.TrimSpace(word)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
