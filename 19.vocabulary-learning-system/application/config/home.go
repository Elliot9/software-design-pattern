package config

import (
	"context"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
)

func GetHomeSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &HomeHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_KeepForver,
		LifeCycleDistance: 0,
		UIFactory:         &HomeUIComponentFactory{handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type HomeUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *HomeHandlerFactory
}

func (f *HomeUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	return scenes.NewNavigationBreadcrumb(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager))
}

func (f *HomeUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	return scenes.NewSimpleMessenger("Hello, welcome to vocabulary learning system.")
}

func (f *HomeUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *HomeUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewSimplePrompt("What are you looking for? ")
}

type HomeHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *HomeHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	navCommands := scenes.NewNavigateCommand(ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager), "root")
	commands = append(commands, navCommands...)

	commands = append(commands, scenes.NewBackCommand())
	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *HomeHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return nil
}
