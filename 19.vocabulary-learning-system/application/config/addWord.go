package config

import (
	"context"
	"fmt"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
)

func GetAddWordSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &AddWordHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_RemoveOnExit,
		LifeCycleDistance: 0,
		UIFactory:         &AddWordUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type AddWordUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *AddWordHandlerFactory
}

func (f *AddWordUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
	fullPath := sceneManager.GetFullPath(sceneManager.Current().GetName())[1:len(sceneManager.GetFullPath(sceneManager.Current().GetName()))]

	word := ctx.Value(scenes.ContextWord).(word.Word)
	fullPath = append(fullPath, fmt.Sprintf("Add '%s' into word repository", word.Name))
	return scenes.NewSimpleBreadcrumb(fullPath)
}

func (f *AddWordUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	word := ctx.Value(scenes.ContextWord).(word.Word)
	return scenes.NewSimpleMessenger(fmt.Sprintf("The word '%s' has been added.", word.Name))
}

func (f *AddWordUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *AddWordUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("Command: ")
}

type AddWordHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *AddWordHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	commands = append(commands, scenes.NewAnyCommand("Okay, I got it", func(ctx context.Context) {
		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		sceneManager.Back(ctx)
	}))

	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *AddWordHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return nil
}
