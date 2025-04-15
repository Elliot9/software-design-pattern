package sceneManagement

import "context"

type LifeCyclePolicy int

const (
	LifeCyclePolicy_KeepForver    LifeCyclePolicy = -1
	LifeCyclePolicy_RemoveOnExit  LifeCyclePolicy = 0
	LifeCyclePolicy_DistanceBased LifeCyclePolicy = 1
)

type Scene interface {
	Render() []string
	HandleInput(ctx context.Context, input string)
	GetName() string
	GetLifeCyclePolicy() LifeCyclePolicy
	GetLifeCycleDistance() int
	SetSceneManager(sceneManager *SceneManager)
	GetSceneManager() *SceneManager
}

type BaseScene struct {
	name              string
	lifeCyclePolicy   LifeCyclePolicy
	lifeCycleDistance int
	sceneManager      *SceneManager
	Scene
}

func NewBaseScene(name string, lifeCyclePolicy LifeCyclePolicy, lifeCycleDistance int) *BaseScene {
	return &BaseScene{
		name:              name,
		lifeCyclePolicy:   lifeCyclePolicy,
		lifeCycleDistance: lifeCycleDistance,
	}
}

// need to implement Render()
func (s *BaseScene) Render() []string {
	return s.Scene.Render()
}

// need to implement HandleInput
func (s *BaseScene) HandleInput(ctx context.Context, input string) {
	s.Scene.HandleInput(ctx, input)
}

func (s *BaseScene) GetName() string {
	return s.name
}

func (s *BaseScene) GetLifeCyclePolicy() LifeCyclePolicy {
	return s.lifeCyclePolicy
}

func (s *BaseScene) GetLifeCycleDistance() int {
	return s.lifeCycleDistance
}

func (s *BaseScene) SetSceneManager(sceneManager *SceneManager) {
	s.sceneManager = sceneManager
}

func (s *BaseScene) GetSceneManager() *SceneManager {
	return s.sceneManager
}
