package scenes

import (
	"context"
	"fmt"
	"github/elliot9/class19/pkg/sceneManagement"
)

type Command interface {
	Execute(ctx context.Context)
	GetKey() string
	GetGoal() string
}

type BaseCommand struct {
	Key     string
	Goal    string
	Handler func(ctx context.Context)
}

func NewCommand(key string, goal string, handler func(ctx context.Context)) Command {
	return &BaseCommand{Key: key, Goal: goal, Handler: handler}
}

func (c *BaseCommand) Execute(ctx context.Context) {
	c.Handler(ctx)
}

func (c *BaseCommand) GetKey() string {
	return c.Key
}

func (c *BaseCommand) GetGoal() string {
	return c.Goal
}

func NewEscCommand() Command {
	return &BaseCommand{Key: "/ESC", Goal: "Exit", Handler: func(ctx context.Context) {
		sceneManager := ctx.Value(ContextSceneManager).(*sceneManagement.SceneManager)
		sceneManager.Stop()
	}}
}

func NewBackCommand() Command {
	return &BaseCommand{Key: "/B", Goal: "Previous Page", Handler: func(ctx context.Context) {
		sceneManager := ctx.Value(ContextSceneManager).(*sceneManagement.SceneManager)
		sceneManager.Back(ctx)
	}}
}

func NewNavigateCommand(sceneManager *sceneManagement.SceneManager, currentScene string) []Command {
	next := sceneManager.GetChildren(currentScene)
	commands := make([]Command, len(next))

	for i, scene := range next {
		commands[i] = &BaseCommand{Key: fmt.Sprintf("/%d", i+1), Goal: scene, Handler: func(ctx context.Context) {
			sceneManager.Next(scene, ctx)
		}}
	}

	return commands
}

type AnyCommand struct {
	BaseCommand
}

func NewAnyCommand(Goal string, Handler func(ctx context.Context)) Command {
	return &AnyCommand{BaseCommand: BaseCommand{Key: "/*", Goal: Goal, Handler: Handler}}
}

var _ Command = &AnyCommand{}
var _ Command = &BaseCommand{}
