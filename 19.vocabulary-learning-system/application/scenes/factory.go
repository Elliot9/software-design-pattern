package scenes

import (
	"context"
	"github/elliot9/class19/pkg/sceneManagement"
)

type UIComponentFactory interface {
	CreateBreadcrumb(ctx context.Context) Breadcrumb
	CreateMessenger(ctx context.Context) Messenger
	CreateMenu(ctx context.Context) Menu
	CreatePrompt(ctx context.Context) Prompt
}

type HandlerFactory interface {
	CreateCommands(ctx context.Context) []Command
	CreateInputHandler() func(ctx context.Context, input string)
}

type SceneConfig struct {
	Name              string
	LifeCyclePolicy   sceneManagement.LifeCyclePolicy
	LifeCycleDistance int
	UIFactory         UIComponentFactory
	HandlerFactory    HandlerFactory
}

type VocabularySceneFactory struct {
	sceneManagement.SceneFactory
	Config SceneConfig
}

func NewVocabularySceneFactory(config SceneConfig) *VocabularySceneFactory {
	return &VocabularySceneFactory{
		Config: config,
	}
}

func (f *VocabularySceneFactory) Create(ctx context.Context) sceneManagement.Scene {
	return NewVocabularyLearningScene(
		f.Config.Name,
		f.Config.LifeCyclePolicy,
		f.Config.LifeCycleDistance,
		f.Config.UIFactory.CreateBreadcrumb(ctx),
		f.Config.UIFactory.CreateMessenger(ctx),
		f.Config.UIFactory.CreateMenu(ctx),
		f.Config.UIFactory.CreatePrompt(ctx),
		f.Config.HandlerFactory.CreateCommands(ctx),
		f.Config.HandlerFactory.CreateInputHandler(),
	)
}
