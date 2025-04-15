package config

import (
	"context"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
)

func GetManageWordsSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &ManageWordsHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_KeepForver,
		LifeCycleDistance: 0,
		UIFactory:         &ManageWordsUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type ManageWordsUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *ManageWordsHandlerFactory
}

func (f *ManageWordsUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	return scenes.NewNavigationBreadcrumb(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager))
}

func (f *ManageWordsUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	return scenes.NewSimpleMessenger("Wanna learn a new word?")
}

func (f *ManageWordsUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *ManageWordsUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("Hello, please select a place to go: ")
}

type ManageWordsHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *ManageWordsHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	navCommands := scenes.NewNavigateCommand(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager), "Manage words")
	commands = append(commands, navCommands...)

	commands = append(commands, scenes.NewBackCommand())
	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *ManageWordsHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return nil
}
