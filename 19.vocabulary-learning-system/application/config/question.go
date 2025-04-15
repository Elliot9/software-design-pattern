package config

import (
	"context"
	"fmt"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
)

func GetQuestionSceneConfig(name string, sceneManager *sceneManagement.SceneManager) scenes.SceneConfig {
	handlerFactory := &QuestionHandlerFactory{}

	return scenes.SceneConfig{
		Name:              name,
		LifeCyclePolicy:   sceneManagement.LifeCyclePolicy_DistanceBased,
		LifeCycleDistance: 3,
		UIFactory:         &QuestionUIComponentFactory{name: name, handlerFactory: handlerFactory},
		HandlerFactory:    handlerFactory,
	}
}

type QuestionUIComponentFactory struct {
	scenes.UIComponentFactory
	handlerFactory *QuestionHandlerFactory
	name           string
}

func (f *QuestionUIComponentFactory) CreateBreadcrumb(ctx context.Context) scenes.Breadcrumb {
	sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
	fullPath := sceneManager.GetFullPath(f.name)[1 : len(sceneManager.GetFullPath(f.name))-1]
	wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)

	questionIndex := wordService.GetCurrentQuiz().GetCurrentIndex()

	fullPath = append(fullPath, fmt.Sprintf("Question %d", questionIndex+1))
	return scenes.NewSimpleBreadcrumb(fullPath)
}

func (f *QuestionUIComponentFactory) CreateMessenger(ctx context.Context) scenes.Messenger {
	wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)
	return scenes.NewQuestionMessenger(wordService)
}

func (f *QuestionUIComponentFactory) CreateMenu(ctx context.Context) scenes.Menu {
	return scenes.NewSimpleMenu(f.handlerFactory.CreateCommands(ctx))
}

func (f *QuestionUIComponentFactory) CreatePrompt(ctx context.Context) scenes.Prompt {
	return scenes.NewQuestionPrompt(ctx.Value(scenes.ContextWordService).(*word.WordService))
}

type QuestionHandlerFactory struct {
	scenes.HandlerFactory
}

func (f *QuestionHandlerFactory) CreateCommands(ctx context.Context) (commands []scenes.Command) {
	commands = append(commands, scenes.NewBackCommand())
	commands = append(commands, scenes.NewEscCommand())

	return commands
}

func (f *QuestionHandlerFactory) CreateInputHandler() func(ctx context.Context, input string) {
	return func(ctx context.Context, input string) {
		sceneManager := ctx.Value(scenes.ContextSceneManager).(*sceneManagement.SceneManager)
		scene := sceneManager.Current().(*scenes.VocabularyLearningScene)

		wordService := ctx.Value(scenes.ContextWordService).(*word.WordService)
		if wordService.GetCurrentQuiz() == nil {
			sceneManager.Back(ctx)
			return
		}

		isCorrect, answer := wordService.CheckAnswer(input)

		if isCorrect {
			scene.SetMessenger(scenes.NewSimpleMessenger(fmt.Sprintf("You got the answer. The answer is %s.", answer)))
		} else {
			scene.SetMessenger(scenes.NewSimpleMessenger(fmt.Sprintf("You missed it! The answer is %s.", answer)))
		}

		commands := []scenes.Command{
			scenes.NewAnyCommand("Continue", func(ctx context.Context) {
				currentQuiz := wordService.GetCurrentQuiz()

				if currentQuiz == nil {
					fullPath := sceneManager.GetFullPath(sceneManager.Current().GetName())[1 : len(sceneManager.GetFullPath(sceneManager.Current().GetName()))-1]
					fullPath = append(fullPath, "Exam's Over")
					scene.SetBreadcrumb(scenes.NewSimpleBreadcrumb(fullPath))
					commands := []scenes.Command{scenes.NewAnyCommand("End", func(ctx context.Context) {
						sceneManager.Back(ctx)
					}), scenes.NewEscCommand()}
					scene.RegisterCommands(commands)
					scene.SetMenu(scenes.NewSimpleMenu(commands))
					scene.SetMessenger(scenes.NewResultMessenger(wordService))
				} else {
					sceneManager.Refresh(ctx)
				}
			}),

			scenes.NewEscCommand(),
		}

		scene.RegisterCommands(commands)
		scene.SetMenu(scenes.NewSimpleMenu(commands))

		fullPath := sceneManager.GetFullPath(sceneManager.Current().GetName())[1 : len(sceneManager.GetFullPath(sceneManager.Current().GetName()))-1]
		fullPath = append(fullPath, "Result")
		scene.SetBreadcrumb(scenes.NewSimpleBreadcrumb(fullPath))
		scene.SetPrompt(scenes.NewSimplePrompt("Command: "))
	}
}
