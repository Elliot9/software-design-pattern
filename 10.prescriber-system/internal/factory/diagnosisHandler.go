package factory

import "github/elliot9/class10/internal/core"

type DiagnosisHandlerFactory struct {
	handlers map[string]func() core.DiagnosisHandler
}

func NewDiagnosisHandlerFactory() *DiagnosisHandlerFactory {
	return &DiagnosisHandlerFactory{
		handlers: make(map[string]func() core.DiagnosisHandler),
	}
}

func (f *DiagnosisHandlerFactory) Register(name string, creator func() core.DiagnosisHandler) {
	f.handlers[name] = creator
}

func (f *DiagnosisHandlerFactory) Create(name string) core.DiagnosisHandler {
	if creator, exists := f.handlers[name]; exists {
		return creator()
	}
	return nil
}
