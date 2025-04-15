package application

import (
	"context"
	"github/elliot9/class19/application/config"
	"github/elliot9/class19/application/scenes"
	"github/elliot9/class19/pkg/sceneManagement"
	"github/elliot9/class19/word"
)

type VocabularyLearningSystem struct {
	wordService  *word.WordService
	sceneControl *sceneManagement.CLISceneControl
}

func NewVocabularyLearningSystem(finder word.Finder, repository word.Repository, randomizer word.Randomizer, cli sceneManagement.CLI) *VocabularyLearningSystem {
	wordService := word.NewWordService(finder, repository, randomizer)
	sceneManager := sceneManagement.NewSceneManager()
	sceneControl := sceneManagement.NewCLISceneControl(sceneManager, cli)

	// 註冊場景
	registerScenes(sceneManager)

	return &VocabularyLearningSystem{
		wordService:  wordService,
		sceneControl: sceneControl,
	}
}

type SceneName string

const (
	SceneName_Home        SceneName = "root"
	SceneName_ManageWords SceneName = "Manage words"
	SceneName_SearchWords SceneName = "Search words"
	SceneName_AddWord     SceneName = "Add word"
	SceneName_AddNewWords SceneName = "Add new words"
	SceneName_DeleteWords SceneName = "Delete words"
	SceneName_ReviewWords SceneName = "Review words"
	SceneName_Question    SceneName = "Question"
)

// 場景註冊函數 - 集中管理所有場景
func registerScenes(sceneManager *sceneManagement.SceneManager) {
	// 根場景
	sceneManager.SetRootSceneFactory(string(SceneName_Home), scenes.NewVocabularySceneFactory(config.GetHomeSceneConfig(string(SceneName_Home), sceneManager)))

	// 其他所有場景
	childSceneFactories := []struct {
		ParentSceneName SceneName
		ChildSceneName  SceneName
		SceneFactory    sceneManagement.SceneFactory
	}{
		{SceneName_Home, SceneName_ManageWords, scenes.NewVocabularySceneFactory(config.GetManageWordsSceneConfig(string(SceneName_ManageWords), sceneManager))},
		{SceneName_ManageWords, SceneName_SearchWords, scenes.NewVocabularySceneFactory(config.GetSearchWordsSceneConfig(string(SceneName_SearchWords), sceneManager))},
		{SceneName_SearchWords, SceneName_AddWord, scenes.NewVocabularySceneFactory(config.GetAddWordSceneConfig(string(SceneName_AddWord), sceneManager))},
		{SceneName_ManageWords, SceneName_AddNewWords, scenes.NewVocabularySceneFactory(config.GetAddNewWordsSceneConfig(string(SceneName_AddNewWords), sceneManager))},
		{SceneName_ManageWords, SceneName_DeleteWords, scenes.NewVocabularySceneFactory(config.GetDeleteWordsSceneConfig(string(SceneName_DeleteWords), sceneManager))},
		{SceneName_Home, SceneName_ReviewWords, scenes.NewVocabularySceneFactory(config.GetReviewWordsSceneConfig(string(SceneName_ReviewWords), sceneManager))},
		{SceneName_ReviewWords, SceneName_Question, scenes.NewVocabularySceneFactory(config.GetQuestionSceneConfig(string(SceneName_Question), sceneManager))},
	}

	for _, factory := range childSceneFactories {
		sceneManager.AppendChildSceneFactory(string(factory.ParentSceneName), string(factory.ChildSceneName), factory.SceneFactory)
	}
}

func (v *VocabularyLearningSystem) GetWordService() *word.WordService {
	return v.wordService
}

func (v *VocabularyLearningSystem) GetSceneControl() *sceneManagement.CLISceneControl {
	return v.sceneControl
}

func (v *VocabularyLearningSystem) Run() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, scenes.ContextSceneManager, v.sceneControl.GetSceneManager())
	ctx = context.WithValue(ctx, scenes.ContextWordService, v.wordService)
	v.sceneControl.Start(ctx)
}
