package scenes

import (
	"context"
	"github/elliot9/class19/pkg/sceneManagement"
	"strings"
)

type VocabularyLearningScene struct {
	sceneManagement.BaseScene
	breadcrumb   Breadcrumb
	messenger    Messenger
	menu         Menu
	prompt       Prompt
	commands     []Command
	inputHandler func(ctx context.Context, input string)
}

func (s *VocabularyLearningScene) Render() []string {
	ui := []string{}
	ui = append(ui, s.breadcrumb.Render()...)
	ui = append(ui, "---")

	if s.messenger != nil {
		messages := s.messenger.Render()
		if len(messages) > 0 {
			ui = append(ui, messages...)
			ui = append(ui, "---")
		}
	}

	ui = append(ui, s.menu.Render()...)
	ui = append(ui, "---")
	ui = append(ui, s.prompt.Render()...)
	return ui
}

func (s *VocabularyLearningScene) HandleInput(ctx context.Context, input string) {
	for _, command := range s.commands {
		if _, ok := command.(*AnyCommand); ok {
			command.Execute(ctx)
			return
		}

		if command.GetKey() == strings.ToUpper(input) {
			command.Execute(ctx)
			return
		}
	}

	s.HandleSpecialInput(ctx, input)
}

func (s *VocabularyLearningScene) HandleSpecialInput(ctx context.Context, input string) {
	if s.inputHandler != nil {
		s.inputHandler(ctx, input)
	}
}

func (s *VocabularyLearningScene) SetBreadcrumb(breadcrumb Breadcrumb) {
	s.breadcrumb = breadcrumb
}

func (s *VocabularyLearningScene) SetMessenger(messenger Messenger) {
	s.messenger = messenger
}

func (s *VocabularyLearningScene) SetMenu(menu Menu) {
	s.menu = menu
}

func (s *VocabularyLearningScene) SetPrompt(prompt Prompt) {
	s.prompt = prompt
}

func (s *VocabularyLearningScene) SetInputHandler(inputHandler func(ctx context.Context, input string)) {
	s.inputHandler = inputHandler
}

func (s *VocabularyLearningScene) RegisterCommands(commands []Command) {
	s.commands = commands
}

func NewVocabularyLearningScene(
	name string,
	lifeCyclePolicy sceneManagement.LifeCyclePolicy,
	lifeCycleDistance int,
	breadcrumb Breadcrumb,
	messenger Messenger,
	menu Menu,
	prompt Prompt,
	commands []Command,
	inputHandler func(ctx context.Context, input string),
) *VocabularyLearningScene {
	scene := &VocabularyLearningScene{
		breadcrumb:   breadcrumb,
		messenger:    messenger,
		menu:         menu,
		prompt:       prompt,
		commands:     commands,
		inputHandler: inputHandler,
	}

	scene.BaseScene = *sceneManagement.NewBaseScene(name, lifeCyclePolicy, lifeCycleDistance)
	scene.BaseScene.Scene = scene

	return scene
}
