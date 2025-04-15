package config

import (
	"context"
	"fmt"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
)

func GetSearchWordsSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &SearchWordsHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_RemoveOnExit,
		LifeCycleDistance: 0,
		UIFactory:         &SearchWordsUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type SearchWordsUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *SearchWordsHandlerFactory
}

func (f *SearchWordsUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	return scenes.NewNavigationBreadcrumb(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager))
}

func (f *SearchWordsUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	return scenes.NewSimpleMessenger("Genius is one percent inspiration and ninety-nine percent perspiration.")
}

func (f *SearchWordsUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *SearchWordsUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("Please input a word's name: ")
}

type SearchWordsHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *SearchWordsHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	commands = append(commands, scenes.NewBackCommand())
	commands = append(commands, scenes.NewEscCommand())
	return commands
}

func (f *SearchWordsHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return func(ctx context.Context, input string) {
		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)

		word, found := wordService.Search(input)
		scene := sceneManager.Current().(*scenes.VocabularyLearningScene)

		if !found {
			scene.SetMessenger(scenes.NewSimpleMessenger(fmt.Sprintf("Cannot find the word '%s'.", input)))
			return
		}

		scene.SetMessenger(scenes.NewSearchWordsMessenger(&word))

		addWordCommand := scenes.NewCommand("/1", fmt.Sprintf("Add %s into word repository", word.Name), func(ctx context.Context) {
			wordService.AddWord(word)
			ctx = context.WithValue(ctx, scenes.ContextWord, word)
			sceneManager.Next("Add word", ctx)
		})

		commands := []scenes.Command{addWordCommand, scenes.NewBackCommand(), scenes.NewEscCommand()}
		scene.RegisterCommands(commands)
		scene.SetMenu(scenes.NewSimpleMenu(commands))
	}
}
