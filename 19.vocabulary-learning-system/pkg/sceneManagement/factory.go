package sceneManagement

import "context"

type SceneFactory interface {
	Create(context context.Context) Scene
}
