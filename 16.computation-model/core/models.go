package core

type Models interface {
	CreateModel(modelName string) Model
}

type BaseModels struct{}

func NewModels() Models {
	return &BaseModels{}
}

func (m *BaseModels) CreateModel(modelName string) Model {
	return NewModel(modelName)
}

var _ Models = &BaseModels{}
