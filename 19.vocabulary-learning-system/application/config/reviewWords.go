package config

import (
	"context"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
)

func GetReviewWordsSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &ReviewWordsHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_RemoveOnExit,
		LifeCycleDistance: 0,
		UIFactory:         &ReviewWordsUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type ReviewWordsUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *ReviewWordsHandlerFactory
}

func (f *ReviewWordsUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	return scenes.NewNavigationBreadcrumb(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager))
}

func (f *ReviewWordsUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	return scenes.NewGreetingMessenger(ctx.Value(scenes.ContextWordService).(*word.WordService))
}

func (f *ReviewWordsUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *ReviewWordsUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("Answer: ")
}

type ReviewWordsHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *ReviewWordsHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	commands = append(commands, scenes.NewCommand("Y", "Let's do it.", func(ctx context.Context) {
		wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)
		wordService.StartNewQuiz(10)

		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		sceneManager.Next("Question", ctx)
	}))

	commands = append(commands, scenes.NewCommand("N", "No, not today.", func(ctx context.Context) {
		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		sceneManager.Back(ctx)
	}))
	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *ReviewWordsHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return nil
}
