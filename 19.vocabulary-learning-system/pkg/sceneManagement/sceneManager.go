package sceneManagement

import (
	"context"
)

type SceneManager struct {
	sceneStack   map[string]Scene
	factories    map[string]SceneFactory
	currentScene Scene
	relations    map[string][]string // parent -> children
	parents      map[string]string   // child -> parent
	done         chan struct{}
	rootName     string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		sceneStack:   make(map[string]Scene),
		factories:    make(map[string]SceneFactory),
		currentScene: nil,
		relations:    make(map[string][]string),
		parents:      make(map[string]string),
		done:         make(chan struct{}),
		rootName:     "",
	}
}

func (s *SceneManager) Start(context context.Context) {
	s.moveTo(s.rootName, context)
}

func (s *SceneManager) Stop() {
	close(s.done)
}

func (s *SceneManager) SetRootSceneFactory(rootName string, factory SceneFactory) {
	s.rootName = rootName
	s.factories[s.rootName] = factory
}

func (s *SceneManager) AppendChildSceneFactory(parentName string, childName string, childFactory SceneFactory) {
	if s.relations[parentName] == nil {
		s.relations[parentName] = []string{}
	}

	s.factories[childName] = childFactory
	s.relations[parentName] = append(s.relations[parentName], childName)
	s.parents[childName] = parentName
}

func (s *SceneManager) Next(name string, context context.Context) {
	if s.relations[s.currentScene.GetName()] == nil {
		return
	}

	for _, child := range s.GetChildren(s.currentScene.GetName()) {
		if child == name {
			s.moveTo(name, context)
			return
		}
	}
}

func (s *SceneManager) Back(context context.Context) {
	if _, ok := s.parents[s.currentScene.GetName()]; !ok {
		return
	}

	s.moveTo(s.parents[s.currentScene.GetName()], context)
}

func (s *SceneManager) Refresh(context context.Context) {
	delete(s.sceneStack, s.currentScene.GetName())
	s.setCurrentScene(s.currentScene.GetName(), context)
}

func (s *SceneManager) Current() Scene {
	return s.currentScene
}

func (s *SceneManager) moveTo(name string, context context.Context) {
	s.onExit()
	s.setCurrentScene(name, context)
	s.onEnter()
}

func (s *SceneManager) calculateDistance(fromSceneName, toSceneName string) int {
	if fromSceneName == toSceneName {
		return 0
	}

	// 使用BFS計算最短距離
	visited := make(map[string]bool)
	queue := []struct {
		name     string
		distance int
	}{
		{fromSceneName, 0},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.name == toSceneName {
			return current.distance
		}

		if visited[current.name] {
			continue
		}
		visited[current.name] = true

		for _, childName := range s.relations[current.name] {
			if !visited[childName] {
				queue = append(queue, struct {
					name     string
					distance int
				}{childName, current.distance + 1})
			}
		}

		if parentName, exists := s.parents[current.name]; exists && !visited[parentName] {
			queue = append(queue, struct {
				name     string
				distance int
			}{parentName, current.distance + 1})
		}
	}

	return -1
}

func (s *SceneManager) setCurrentScene(name string, context context.Context) {
	current := s.sceneStack[name]

	if current == nil {
		current = s.factories[name].Create(context)
	}

	s.currentScene = current
}

func (s *SceneManager) onExit() {
	if s.currentScene == nil {
		return
	}

	// 離開即清空
	if s.currentScene.GetLifeCyclePolicy() == LifeCyclePolicy_RemoveOnExit {
		delete(s.sceneStack, s.currentScene.GetName())
	}

	// 其餘皆會保留
	s.sceneStack[s.currentScene.GetName()] = s.currentScene
}

func (s *SceneManager) onEnter() {
	if s.currentScene == nil {
		return
	}

	validSceneStack := map[string]Scene{}

	// 進入即計算哪些需要清空
	for _, scene := range s.sceneStack {
		if scene.GetLifeCyclePolicy() == LifeCyclePolicy_KeepForver {
			validSceneStack[scene.GetName()] = scene
			continue
		}

		distance := s.calculateDistance(s.currentScene.GetName(), scene.GetName())
		if distance <= scene.GetLifeCycleDistance() {
			validSceneStack[scene.GetName()] = scene
		}
	}

	s.sceneStack = validSceneStack
}

func (s *SceneManager) GetParent(sceneName string) (string, bool) {
	parent, exists := s.parents[sceneName]
	return parent, exists
}

func (s *SceneManager) GetChildren(sceneName string) []string {
	return s.relations[sceneName]
}

func (s *SceneManager) GetFullPath(sceneName string) []string {
	path := []string{}
	current := sceneName

	for current != "" {
		path = append(path, current)

		if parent, ok := s.parents[current]; ok {
			current = parent
		} else {
			current = ""
		}
	}

	// reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (s *SceneManager) ShouldContinue() bool {
	select {
	case <-s.done:
		return false
	default:
		return true
	}
}
