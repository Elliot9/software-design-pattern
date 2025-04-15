package sceneManagement

import (
	"context"
)

type CLISceneControl struct {
	sceneManager *SceneManager
	cli          CLI
}

func NewCLISceneControl(sceneManager *SceneManager, cli CLI) *CLISceneControl {
	return &CLISceneControl{
		sceneManager: sceneManager,
		cli:          cli,
	}
}

func (c *CLISceneControl) Start(ctx context.Context) {
	c.sceneManager.Start(ctx)

	for c.sceneManager.ShouldContinue() {
		// 渲染當前場景
		for _, line := range c.sceneManager.Current().Render() {
			c.cli.Println(line)
		}

		// 獲取用戶輸入
		input := c.cli.ReadLine()

		// 處理輸入
		c.sceneManager.Current().HandleInput(ctx, input)
	}
}

func (c *CLISceneControl) GetSceneManager() *SceneManager {
	return c.sceneManager
}
